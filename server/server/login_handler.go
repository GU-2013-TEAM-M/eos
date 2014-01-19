package main

import (
    "errors"
    "eos/server/db"
    "labix.org/v2/mgo/bson"
    "time"
    "strconv"
)

// a handler that logs in the user
func LoginHandler(cmd *CmdMessage) error {
    // authenticate the user
    if cmd.Conn.owner.IsUser() {
        return userLoginHandler(cmd)
    }
    return daemonLoginHandler(cmd)

    data := make(map[string]interface{})
    data["session_id"] = "52a4ed348350a921bd000001"
    if cmd.Conn.owner.IsUser() {
        cmd.Conn.owner.(*User).Id = strconv.Itoa(time.Now().Nanosecond())
        cmd.Conn.owner.(*User).OrgId = "Anonymous"
        data["id"] = cmd.Conn.owner.(*User).Id
    } else {
        cmd.Conn.owner.(*Daemon).Id = strconv.Itoa(time.Now().Nanosecond())
        cmd.Conn.owner.(*Daemon).OrgId = "Anonymous"
        data["id"] = cmd.Conn.owner.(*Daemon).Id
    }
    cmd.Conn.owner.Authorise()
    DispatchMessage("login", data, cmd.Conn)
    return errors.New("Login: Not implemented")
}

func userLoginHandler(cmd *CmdMessage) error {
    data := make(map[string]interface{})
    email, ok1 := cmd.Data["email"].(string)
    pass, ok2 := cmd.Data["password"].(string)

    if !(ok1 && ok2) {
        data["msg"] = "Missing email or password"
        DispatchMessage("error", data, cmd.Conn)
        return errors.New("Login failed: missing email or password")
    }

    user := &db.User{}

    err := db.C("users").Find(bson.M{"email": email, "password": pass}).One(user)

    if err != nil {
        data["msg"] = "Wrong email or password"
        DispatchMessage("error", data, cmd.Conn)
        return errors.New("Login failed: bad email or password")
    }

    // creating the session
    sess := &db.Session{UId: user.Id}
    sess.GenId()
    db.C("sessions").Insert(sess)

    data["session_id"] = sess.Id.Hex()
    data["id"] = user.Id.Hex()

    cmd.Conn.owner.Authorise()
    DispatchMessage("login", data, cmd.Conn)
    return nil
}

func daemonLoginHandler(cmd *CmdMessage) error {
    data := make(map[string]interface{})
    name, ok1 := cmd.Data["name"].(string)
    pass, ok2 := cmd.Data["password"].(string)
    orgId, ok3 := cmd.Data["org_id"].(string)

    if !(ok1 && ok2 && ok3) {
        data["msg"] = "Missing name or password or organisation id"
        DispatchMessage("error", data, cmd.Conn)
        return errors.New("Login failed: missing name or password or org_id")
    }

    if !bson.IsObjectIdHex(orgId) {
        data["msg"] = "Organisation ID is invalid"
        DispatchMessage("error", data, cmd.Conn)
        return errors.New("Organisation ID is invalid")
    }

    daemon := &db.Daemon{}

    err := db.C("daemons").Find(bson.M{
        "name": name,
        "password": pass,
        "org_id": bson.ObjectIdHex(orgId),
    }).One(daemon)

    if err != nil {
        err = db.C("daemons").Find(bson.M{
            "name": name,
            "org_id": bson.ObjectIdHex(orgId),
        }).One(daemon)
        // if there is a mistake in password then it is an error
        if err == nil {
            data["msg"] = "Wrong password"
            DispatchMessage("error", data, cmd.Conn)
            return errors.New("Login failed: bad password")
        }
        // otherwise we have to create a new entry
        daemon = &db.Daemon{
            Name: name,
            Password: pass,
            OrgId: bson.ObjectIdHex(orgId),
            Status: "NOT_KNOWN",
        }
        daemon.GenId()
        db.C("daemons").Insert(daemon)
    }

    data["id"] = daemon.Id.Hex()

    cmd.Conn.owner.Authorise()
    DispatchMessage("login", data, cmd.Conn)
    return nil
}
