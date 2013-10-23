package main

import (
    "code.google.com/p/go.net/websocket"
)

type Connection struct {
    // The websocket Connection.
    ws *websocket.Conn

    // Buffered channel of outbound messages.
    send chan string

    // The User or Daemon attached to this connection
    owner Organisable
}

func (c *Connection) reader() {
    for {
        var message string
        err := websocket.Message.Receive(c.ws, &message)
        if err != nil {
            break
        }
        h.broadcast <- message
    }
    c.ws.Close()
}

func (c *Connection) writer() {
    for message := range c.send {
        err := websocket.Message.Send(c.ws, message)
        if err != nil {
            break
        }
    }
    c.ws.Close()
}

// a function that closes the connection
// deletes itself from the hub
// and removes User/Daemon from the Organisation
func (c *Connection) Close() {
    // delete User/Daemon from the organisation
    c.owner.Deauthorise()

    // remove the connection from the hub, closing its channel
    delete(h.connections, c)
    close(c.send)

    // closing websockets
    go c.ws.Close()
}
