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
