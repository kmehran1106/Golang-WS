package main

type Message struct {
	roomId int
	jsonStr []byte
}

type Hub struct {
	rooms      map[int]map[*Client]bool
	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		rooms:      make(map[int]map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			room := h.rooms[client.roomId]
			if room == nil {
				room = make(map[*Client]bool)
				h.rooms[client.roomId] = room
			}
			room[client] = true
		case client := <-h.unregister:
			room := h.rooms[client.roomId]
			if room != nil {
				if _, ok := room[client]; ok {
					delete(room, client)
					close(client.send)
					if len(room) == 0 {
						delete(h.rooms, client.roomId)
					}
				}
			}
		case message := <-h.broadcast:
			room := h.rooms[message.roomId]
			if room != nil {
				for client := range room {
					select {
					case client.send <- message.jsonStr:
					default:
						close(client.send)
						delete(room, client)
					}
				}
				if len(room) == 0 {
					delete(h.rooms, message.roomId)
				}
			}
		}
	}
}
