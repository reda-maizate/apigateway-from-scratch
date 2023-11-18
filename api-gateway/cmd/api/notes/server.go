package main

import (
	_go "api-gateway/api/v1/gen/go"
	repository "api-gateway/internal/adapters/repository/postgres"
	"api-gateway/internal/core/services"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type NoteServiceServer struct {
	svc services.NoteService
	_go.UnimplementedNoteServer
}

func (nss *NoteServiceServer) CreateNote(ctx context.Context, req *_go.CreateNoteRequest) (*emptypb.Empty, error) {
	err := nss.svc.Create(req.Note.Title, req.Note.Content)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (nss *NoteServiceServer) GetAllNotes(ctx context.Context, req *_go.GetAllNotesRequest) (*_go.GetAllNotesResponse, error) {
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

func main() {
	listener, err := net.Listen("tcp", ":50053")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	store := repository.NewAPIGatewayRepository()
	svc := services.NewNoteService(store)

	server := NoteServiceServer{svc: *svc}

	grpcServer := grpc.NewServer()

	_go.RegisterNoteServer(grpcServer, &server)

	log.Println("Serving Notes-service in gRPC Server on :50053")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
