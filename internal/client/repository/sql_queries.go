package repository

// room
const (
	queryCreateRoom = `
insert into chat."room"
("name", "owner")
values 
($1, $2)
returning id;
`
	// TODO вот это надо бы проверить
	queryGetRooms = `
select r.id, r.name, r.owner, array_agg(ur.user_id) as client_ids
from chat."room" r
join user_room ur on r.id = ur.room_id
where ur.user_id = $1
group by 
r.id, r.name, r.owner;
`
	queryAddClientToRoom = `
insert into user_room
(user_id, room_id)
values 
($1, $2);
`
	queryRemoveClientFromRoom = `
delete from user_room
where user_id = $1
and room_id = $2;
`
)

// message
const (
	queryCreateMessage = `
insert into chat."message"
(room_id, client_id, "message", sent_at)
values 
($1, $2, $3, to_timestamp($4));
`
	queryGetMessagesFromRoom = `
select (id, room_id, client_id, "message", sent_at)
from chat."message"
where room_id = $1
`
)
