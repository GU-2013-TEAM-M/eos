package main

import (
    "errors"
    "testing"
    "eos/server/test"
)

//-------------------------------------------------------
// test controller function
//-------------------------------------------------------
// helper, to check function calls
func getHandlerSpy() (func(*CmdMessage) error, *bool) {
    called := false
    return func(cmd *CmdMessage) error {
        called = true
        if cmd.Type != "_test" {
            return errors.New("test error")
        }
        return nil
    }, &called
}

// this one is not-so-unit test, but there is no easy way for me
// to check, whether ParseMsg or RunCmd were called
// so I will simply check, that it handles happy and bad cases fully
func Test_HandleMsg(t *testing.T) {
    spy, called := getHandlerSpy()
    goodMsg := &Message{ msg:`{"type":"_test","data":{}}` }
    badMsg := &Message{ msg:`{"type":"_error","data":{}}` }

    handlers["_test"] = spy
    handlers["_error"] = spy

    err, hName := HandleMsg(goodMsg)
    test.Assert(*called == true, "calls the function", t)
    test.Assert(err == nil, "does not throw without an error", t)
    test.Assert(hName == "_test", "returns a handler type", t)

    *called = false
    err, hName = HandleMsg(badMsg)
    test.Assert(*called == true, "calls the another function", t)
    test.Assert(err != nil, "pipes through the error from it", t)
    test.Assert(hName == "_error", "returns a handler type", t)

    // cleaning up
    delete(handlers, "_test"); delete(handlers, "_error")
}

func Test_RegisterHandler(t *testing.T) {
    spy, called := getHandlerSpy()

    oldLen := len(handlers)
    RegisterHandler("_test", spy)

    test.Assert(len(handlers) - oldLen == 1, "created a new handler", t)
    test.Assert(*called == false, "our spy works", t)
    handlers["_test"](&CmdMessage{})
    test.Assert(*called == true, "associated with specified function", t)

    // cleaning up
    delete(handlers, "_test")
}

func Test_DeregisterHandler(t *testing.T) {
    spy, _ := getHandlerSpy()

    handlers["_test"] = spy

    oldLen := len(handlers)
    DeregisterHandler("_test")
    test.Assert(oldLen - len(handlers) == 1, "deletes the handler", t)
    err := DeregisterHandler("_again")
    test.Assert(err != nil, "fails when no such handler exists", t)

    // cleaning up
    delete(handlers, "_test")
}

func Test_RunCmd(t *testing.T) {
    spy, called := getHandlerSpy()
    goodCmd := &CmdMessage{ Type: "_test" }
    badCmd := &CmdMessage{ Type: "_error" }

    handlers["_test"] = spy
    handlers["_error"] = spy

    err := RunCmd(goodCmd)
    test.Assert(*called == true, "calls the function", t)
    test.Assert(err == nil, "does not throw without an error", t)

    *called = false
    err = RunCmd(badCmd)
    test.Assert(*called == true, "calls the another function", t)
    test.Assert(err != nil, "pipes through the error from it", t)

    // cleaning up
    delete(handlers, "_test"); delete(handlers, "_error")
}

//-------------------------------------------------------
// test parsing the json
//-------------------------------------------------------
func Test_ParseMsg(t *testing.T) {
    m := &Message{ msg: `{"type":"test","data":{"foo":"bar"}}`, c: &Connection{} }
    cmd, err := ParseMsg(m)

    test.Assert(err == nil, "successfully parses the correct message", t)
    test.Assert(cmd.Type == "test", "decodes the type", t)
    test.Assert(cmd.Conn == m.c, "still keeps the connection", t)
    foo := cmd.Data["foo"].(string)
    test.Assert(foo == "bar", "gets the internals", t)

    m = &Message{ msg: `{malformed: "json"}` }
    cmd, err = ParseMsg(m)
    test.Assert(err != nil, "fails on malformed json", t)

    m = &Message{ msg: `{"type":"test","data":{}}` }
    cmd, err = ParseMsg(m)
    test.Assert(err == nil, "does not fail on empty data", t)
}

//-------------------------------------------------------
// test generating the json
//-------------------------------------------------------
func Test_GetMessage(t *testing.T) {
    cmd := &CmdMessage{ Type: "test", Data: make(map[string]interface{}) }
    cmd.Data["foo"] = "bar"
    cmd.Conn = &Connection{}
    msg, err := GetMessage(cmd)

    test.Assert(err == nil, "successfully stores the correct message", t)
    test.Assert(msg.msg == `{"type":"test","data":{"foo":"bar"}}`, "produces the expected output", t)
    test.Assert(msg.c == cmd.Conn, "preserves the connection", t)
    test.Assert(cmd.Conn != nil, "does not interfere with the structure", t)
}
