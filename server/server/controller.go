package main

import (
    "encoding/json"
    "errors"
)

// a hash containing all the handlers corresponding to commands
var handlers = make(map[string] func(*CmdMessage) error)

// All of the json messages should follow this pattern:
type CmdMessage struct {
    Type string `json:"type"`
    Data map[string]interface{} `json:"data"`
    Conn *Connection
}

// perform an action described in the message
func HandleMsg(m *Message) error {
    // decode json of the message
    c, err := ParseMsg(m)
    if err != nil {
        return err
    }
    // run the corresponding function
    return RunCmd(c)
}

// registering a handler, to be executed when the specific command is sent
func RegisterHandler(cmd string, handler func(*CmdMessage) error) {
    handlers[cmd] = handler
}

// no clue why would I want to use this in this particular application
// but if I was to implement a standalone library, then it is a must.
func DeregisterHandler(cmd string) error {
    _, ok := handlers[cmd]
    if !ok {
        return errors.New("Command not found: " + cmd)
    }

    delete(handlers, cmd)
    return nil
}

// runs a handler attached to the command 
func RunCmd(cmd *CmdMessage) error {
    handler, ok := handlers[cmd.Type]
    if ok == true {
        return handler(cmd)
    }
    return errors.New("No such command was registered: " + cmd.Type)
}

// parses the message
func ParseMsg(m *Message) (*CmdMessage, error) {
    cmd := &CmdMessage{}
    err := json.Unmarshal([]byte(m.msg), cmd)
    if err != nil {
        return nil, err
    }
    cmd.Conn = m.c
    return cmd, nil
}

// creates the json message, that we can send to the channel
func GetMessage(cmd *CmdMessage) (*Message, error) {
    m := &Message{ c: cmd.Conn }
    // we don't want to send a connection
    cmd.Conn = nil

    msg, err := json.Marshal(cmd)

    // but we don't want to make side effects either
    cmd.Conn = m.c
    if err != nil {
        return nil, err
    }

    // Marshal returns bytes, we want strings
    m.msg = string(msg[:])
    return m, nil
}
