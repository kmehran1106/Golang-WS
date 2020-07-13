package main

type Message struct {
	room int
	data []byte
}

type Hub struct {
	rooms map[int]map[*Client]bool

	broadcast chan *Message

	register chan *Client

	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		rooms:    make(map[int]map[*Client]bool),
	}
}

//func (h *Hub) run() {
//	for {
//		select {
//		case client := <-h.register:
//			h.clients[client] = true
//		case client := <-h.unregister:
//			if _, ok := h.clients[client]; ok {
//				delete(h.clients, client)
//				close(client.send)
//			}
//		case message := <-h.broadcast:
//			for client := range h.clients {
//				select {
//				case client.send <- message:
//				default:
//					close(client.send)
//					delete(h.clients, client)
//				}
//			}
//		}
//	}
//}
