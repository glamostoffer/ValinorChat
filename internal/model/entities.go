package model

type (
	Room struct {
		ID        int64   `json:"id" db:"id"`
		Name      string  `json:"name" db:"name"`
		OwnerID   int64   `json:"ownerID" db:"owner"`
		ClientIDs []int64 `json:"clientIDs" db:"client_ids"`
	}

	Message struct {
		RoomID   int64  `json:"roomID" db:"room_id"`
		ClientID int64  `json:"clientID" db:"clientID"`
		Message  string `json:"message" db:"message"`
		SentAt   int64  `json:"sentAt" db:"sent_at"` // unix
	}
)
