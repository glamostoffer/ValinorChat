package grpc

import (
	"context"
	"github.com/glamostoffer/ValinorChat/internal/client/usecase"
	"github.com/glamostoffer/ValinorChat/pkg/convert"
	clientProto "github.com/glamostoffer/ValinorProtos/chat/client_chat"
	"github.com/golang/protobuf/ptypes/empty"
)

type ClientService struct {
	uc usecase.UseCase
	clientProto.UnimplementedClientChatServiceServer
}

func New(uc usecase.UseCase) clientProto.ClientChatServiceServer {
	return &ClientService{
		uc: uc,
	}
}

func (s *ClientService) CreateRoom(
	ctx context.Context,
	req *clientProto.CreateRoomRequest,
) (res *clientProto.CreateRoomResponse, err error) {
	roomID, err := s.uc.CreateRoom(ctx, req.GetName(), req.GetClientID())
	if err != nil {
		return nil, err
	}

	return &clientProto.CreateRoomResponse{RoomID: roomID}, nil
}

func (s *ClientService) GetListOfRooms(
	ctx context.Context,
	req *clientProto.GetListOfRoomsRequest,
) (res *clientProto.GetListOfRoomsResponse, err error) {
	rooms, err := s.uc.GetListOfRooms(ctx, req.GetClientID())
	if err != nil {
		return nil, err
	}

	return &clientProto.GetListOfRoomsResponse{
		Rooms: convert.ListOfRoomsToProto(rooms),
	}, nil
}

func (s *ClientService) AddClientToRoom(
	ctx context.Context,
	req *clientProto.AddClientToRoomRequest,
) (_ *empty.Empty, err error) {
	err = s.uc.AddClientToRoom(ctx, req.GetRoomID(), req.GetClientID())

	return nil, err
}

func (s *ClientService) RemoveClientFromRoom(
	ctx context.Context,
	req *clientProto.RemoveClientFromRoomRequest,
) (_ *empty.Empty, err error) {
	err = s.uc.RemoveClientFromRoom(ctx, req.GetRoomID(), req.GetClientID())

	return nil, err
}

func (s *ClientService) GetMessagesFromRoom(
	ctx context.Context,
	req *clientProto.GetMessagesFromRoomRequest,
) (resp *clientProto.GetMessagesFromRoomResponse, err error) {
	messages, err := s.uc.GetMessages(ctx, req.GetRoomID())

	return &clientProto.GetMessagesFromRoomResponse{
		Messages: convert.MessagesToProto(messages),
	}, err
}
