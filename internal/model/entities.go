package model

type (
	Room struct {
		ID        int64   `json:"id" db:"id"`
		Name      string  `json:"name" db:"name"`
		OwnerID   int64   `json:"ownerID" db:"owner"`
		ClientIDs []int64 `json:"clients"`
	}

	Message struct {
		ID       int64  `json:"-" db:"id"`
		RoomID   int64  `json:"roomID" db:"room_id"`
		ClientID int64  `json:"clientID" db:"client_id"`
		Content  string `json:"content" db:"content"`
		SentAt   int64  `json:"sentAt" db:"sent_at"` // unix
		Username string `json:"username"`
	}

	MessageDTO struct {
		RoomID   int64  `json:"roomID"`
		Content  string `json:"content"`
		Username string `json:"username"`
	}
)
