package main

//-------------------------------------------------------
// data structures
//-------------------------------------------------------
type Hub struct {
    // Registered connections.
    connections map[*Connection]bool

    // Inbound messages from the connections.
    broadcast chan string

    // Register requests from the connections.
    register chan *Connection

    // Unregister requests from connections.
    unregister chan *Connection
}

var h = Hub{
    broadcast:   make(chan string),
    register:    make(chan *Connection),
    unregister:  make(chan *Connection),
    connections: make(map[*Connection]bool),
}

//-------------------------------------------------------
// methods
//-------------------------------------------------------
func (h *Hub) run() {
    for {
        select {
        case c := <-h.register:
            h.connections[c] = true
        case c := <-h.unregister:
            c.Close()
        case m := <-h.broadcast:
            for c := range h.connections {
                select {
                case c.send <- m:
                default:
                    c.Close()
                }
            }
        }
    }
}

