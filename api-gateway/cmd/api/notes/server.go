package main

import (
	_go "api-gateway/api/v1/gen/go"
	repository "api-gateway/internal/adapters/repository/postgres"
	"api-gateway/internal/core/services"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type NoteServiceServer struct {
	svc services.NoteService
	_go.UnimplementedNoteServer
}

func (nss *NoteServiceServer) CreateNote(ctx context.Context, req *_go.CreateNoteRequest) (*emptypb.Empty, error) {
	//log.Println("Print 1 :", req.GetTitle(), req.GetContent())
	err := nss.svc.Create(req.GetTitle(), req.GetContent())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (nss *NoteServiceServer) GetAllNotes(ctx context.Context, req *emptypb.Empty) (*_go.GetAllNotesResponse, error) {
	notes, err := nss.svc.GetAll()
	if err != nil {
		return nil, err
	}

	var notesResponse []*_go.NoteMessage

	for _, note := range notes {
		notesResponse = append(notesResponse, &_go.NoteMessage{
			Title:   note.Title,
			Content: note.Content,
		})
	}

	return &_go.GetAllNotesResponse{Notes: notesResponse}, nil
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler, gatewayRepository *repository.APIGatewayRepository) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Missing context metadata")
	}

	token := md["authorization"]
	if len(token) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "Missing authorization token")
	}

	if !gatewayRepository.AuthTokenExists(token[0]) {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
	}

	return handler(ctx, req)
}

func main() {
	listener, err := net.Listen("tcp", ":50053")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	store := repository.NewAPIGatewayRepository()
	svc := services.NewNoteService(store)

	server := NoteServiceServer{svc: *svc}

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return AuthInterceptor(ctx, req, info, handler, store)
	}))
	grpcServer := grpc.NewServer(opts...)

	_go.RegisterNoteServer(grpcServer, &server)

	log.Println("Serving Notes-service in gRPC Server on :50053")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
