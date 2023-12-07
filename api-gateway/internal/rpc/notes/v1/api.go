package notesv1

import (
	business "api-gateway/internal/business/notes"
	"api-gateway/internal/middlewares"
	notestubs "api-gateway/stubs/go/apigateway-from-scratch/notes/v1"
	userstubs "api-gateway/stubs/go/apigateway-from-scratch/users/v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NoteServiceServer struct {
	notesBusiness business.NotesBusiness
	notestubs.UnimplementedNoteServer
}

func NewNoteServiceServer(notesBusiness *business.NotesBusiness) NoteServiceServer {
	return NoteServiceServer{
		notesBusiness: *notesBusiness,
	}
}

func (s *NoteServiceServer) CreateNote(ctx context.Context, req *notestubs.CreateNoteRequest) (*emptypb.Empty, error) {
	hasPermission, err := middlewares.CheckPermission(ctx, "create")
	if err != nil || !hasPermission {
		return nil, err
	}

	userUuid := ctx.Value("userUuid").(*userstubs.MeUserResponse).GetId()

	createNoteParams := &business.CreateNoteParams{
		Title:    req.GetTitle(),
		Content:  req.GetContent(),
		UserUuid: userUuid,
	}
	err = s.notesBusiness.CreateNote(ctx, createNoteParams)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *NoteServiceServer) GetAllNotes(ctx context.Context, req *emptypb.Empty) (*notestubs.GetAllNotesResponse, error) {
	notes, err := s.notesBusiness.GetAllNotes(ctx)
	if err != nil {
		return nil, err
	}

	var notesResponse []*notestubs.NoteMessage

	for _, note := range notes {
		notesResponse = append(notesResponse, &notestubs.NoteMessage{
			Title:   note.Title,
			Content: note.Content,
		})
	}

	return &notestubs.GetAllNotesResponse{Notes: notesResponse}, nil
}
