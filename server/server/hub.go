package main

//-------------------------------------------------------
// data structures
//-------------------------------------------------------
type Hub struct {
    // Registered connections.
    connections map[*Connection]bool

    // Inbound messages from the connections.
    broadcast chan *Message

    // Register requests from the connections.
    register chan *Connection

    // Unregister requests from connections.
    unregister chan *Connection
}

// a structure used for broadcasting
type Message struct {
    msg string
    c *Connection
}

var h = Hub{
    broadcast:   make(chan *Message),
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
            c.owner.Authorise()
            h.connections[c] = true
        case c := <-h.unregister:
            c.Close()
        case m := <-h.broadcast:
            org := m.c.owner.GetOrg()
            if m.c.owner.IsUser() {
                org.sendToDaemons(m.msg)
            } else {
                org.sendToUsers(m.msg)
            }
            /*for c := range h.connections {
                select {
                case c.send <- m:
                default:
                    c.Close()
                }
            }*/
        }
    }
}

