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
	// TODO исправить
	queryGetRooms = `
select r.id, r.name, r.owner
from chat."room" r
join user_room ur on r.id = ur.room_id
where ur.user_id = $1;
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
(room_id, client_id, content, sent_at)
values 
($1, $2, $3, to_timestamp($4));
`
	queryGetMessagesFromRoom = `
select id, room_id, client_id, content, round(extract(epoch from sent_at)) as sent_at
from chat."message"
where room_id = $1;
`
	queryGetMessages = `
select id, room_id, client_id, content, round(extract(epoch from sent_at)) as sent_at
from chat."message"
order by room_id;`
)
