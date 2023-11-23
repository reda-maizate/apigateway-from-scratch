package main

import (
	pb "api-gateway/api/v1/gen/go"
	repository "api-gateway/internal/adapters/repository/postgres"
	"api-gateway/internal/core/services"
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

const (
	SERVICE  = "notes"
	RESOURCE = "note"
)

type NoteServiceServer struct {
	notesSvc services.NoteService
	pb.UnimplementedNoteServer
}

func (nss *NoteServiceServer) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*emptypb.Empty, error) {
	hasPermission, err := CheckPermission(ctx, "create")
	if err != nil || !hasPermission {
		return nil, err
	}

	userUuid := ctx.Value("userUuid").(*pb.MeUserResponse).GetId()

	err = nss.notesSvc.Create(req.GetTitle(), req.GetContent(), userUuid)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (nss *NoteServiceServer) GetAllNotes(ctx context.Context, req *emptypb.Empty) (*pb.GetAllNotesResponse, error) {
	notes, err := nss.notesSvc.GetAll()
	if err != nil {
		return nil, err
	}

	var notesResponse []*pb.NoteMessage

	for _, note := range notes {
		notesResponse = append(notesResponse, &pb.NoteMessage{
			Title:   note.Title,
			Content: note.Content,
		})
	}

	return &pb.GetAllNotesResponse{Notes: notesResponse}, nil
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Missing context metadata")
	}

	token := md["authorization"]
	if len(token) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "Missing authorization token")
	}

	usersClient, err := grpc.Dial("users_service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial Users-service gRPC server: %v", err)
	}
	defer usersClient.Close()

	usersConn := pb.NewUserClient(usersClient)
	userUuid, err := usersConn.UserFromToken(ctx, &pb.MeUserRequest{Token: token[0]})

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
	}

	ctx = context.WithValue(ctx, "userUuid", userUuid)
	//log.Println("Inserted userUuid in context: ", userUuid)

	return handler(ctx, req)
}

func CheckPermission(ctx context.Context, Action string) (bool, error) {
	userUuid, ok := ctx.Value("userUuid").(*pb.MeUserResponse)

	if !ok {
		return false, status.Errorf(codes.Unauthenticated, "Missing userUuid")
	}

	permissionsClient, err := grpc.Dial("permissions_service:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial Permissions-service gRPC server: %v", err)
	}
	defer permissionsClient.Close()

	permissionsService := pb.NewPermissionClient(permissionsClient)
	hasPermission, err := permissionsService.CheckPermission(ctx, &pb.CheckPermissionRequest{
		UserUuid: userUuid.GetId(),
		Service:  SERVICE,
		Resource: RESOURCE,
	})

	//log.Println("PermissionsInterceptor hasPermission:", hasPermission)
	if err != nil || !hasPermission.GetHasPermission() {
		return false, status.Errorf(codes.PermissionDenied, "You don't have permission to create note")
	}

	return true, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50053")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	store := repository.NewAPIGatewayRepository()
	notesSvc := services.NewNoteService(store)

	server := NoteServiceServer{notesSvc: *notesSvc}

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc.UnaryServerInterceptor(
			AuthInterceptor,
		),
	)))
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterNoteServer(grpcServer, &server)

	log.Println("Serving Notes-service in gRPC Server on :50053")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
