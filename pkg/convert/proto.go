package convert

import (
	"github.com/glamostoffer/ValinorChat/internal/model"
	clientProto "github.com/glamostoffer/ValinorProtos/chat/client_chat"
)

func ListOfRoomsToProto(in []model.Room) (out []*clientProto.Room) {
	out = make([]*clientProto.Room, 0, len(in))

	for _, room := range in {
		protoRoom := clientProto.Room{
			RoomID:    room.ID,
			Name:      room.Name,
			OwnerID:   room.OwnerID,
			ClientIDs: room.ClientIDs,
		}

		out = append(out, &protoRoom)
	}

	return out
}

func MessagesToProto(in []model.Message) (out []*clientProto.Message) {
	out = make([]*clientProto.Message, 0, len(in))

	for _, message := range in {
		protoMessage := clientProto.Message{
			RoomID:   message.RoomID,
			ClientID: message.ClientID,
			Message:  message.Content,
			SentAt:   message.SentAt,
			Username: message.Username,
		}

		out = append(out, &protoMessage)
	}

	return out
}
