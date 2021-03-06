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
func HandleMsg(m *Message) (error, string) {
    // decode json of the message
    c, err := ParseMsg(m)
    if err != nil {
        return err, "unknown"
    }
    // run the corresponding function
    return RunCmd(c), c.Type
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

    c := struct {
        Type string `json:"type"`
        Data map[string](interface{}) `json:"data"`
    }{ cmd.Type, cmd.Data }
    msg, err := json.Marshal(&c)

    if err != nil {
        return nil, err
    }

    // Marshal returns bytes, we want strings
    m.msg = string(msg[:])
    return m, nil
}

// dispatches the message
func DispatchMessage(t string, data map[string]interface{}, c *Connection) error {
    cmd := &CmdMessage{Type: t, Data: data, Conn: c}
    if TEST {
        StoreLastCmd(cmd)
        return nil
    }

    m, err := GetMessage(cmd)
    if err != nil {
        return err
    }
    c.send <- m.msg
    return nil
}

//----------------------------------------------------
// json helpers
//----------------------------------------------------
func ToStringSlice(arr []interface{}) ([]string, error) {
    strs := make([]string, len(arr))
    for i, v := range arr {
        s, ok := v.(string)
        if !ok {
            return nil, errors.New("Not a string given for a string slice")
        }
        strs[i] = s
    }
    return strs, nil
}

//----------------------------------------------------
// test helpers
//----------------------------------------------------
var _test_last_cmd *CmdMessage
// saving the last command
func StoreLastCmd(cmd *CmdMessage) { _test_last_cmd = cmd }
// retrieving the last command
func GetLastCmd() *CmdMessage { return _test_last_cmd }
