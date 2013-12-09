package main

// a handler that deauthorises the user or the daemon
func LogoutHandler(cmd *CmdMessage) error {
    data := make(map[string]interface{})
    cmd.Conn.owner.Deauthorise()
    data["status"] = "OK"
    DispatchMessage("logout", data, cmd.Conn)
    return nil
}
