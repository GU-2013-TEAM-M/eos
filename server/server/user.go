package main

import (
//    "eos/server/db"
)

//-------------------------------------------------------
// data structures
//-------------------------------------------------------
// a user entity, tied to a connection
// is anonymous if only session id is supplied
type User struct {
    Id string
    OrgId string
    SessionId string
    c *Connection
}

//-------------------------------------------------------
// methods
//-------------------------------------------------------
// obtain Id, OrgId and add yourself to a organisation
func (u *User) Authorise() error {
    // TODO: getting organisation id from MongoDB
    orgId := "Anonymous"
    u.OrgId = orgId

    org, err := GetOrg(u.OrgId)
    // if the organisation does not exist -- create one
    if err != nil {
        org = NewOrg(u.OrgId)
    }
    org.addUser(u.Id, u.c)
    return nil
}

// remove itself from the organisation
func (u *User) Deauthorise() error {
    org, err := GetOrg(u.OrgId)
    if err != nil {
        return err
    }

    org.delUser(u.Id)
    u.OrgId = NO_ORG
    return nil
}

// check if a user is authorised
func (u *User) IsAuthorised() bool {
    return u.OrgId != NO_ORG
}

// type checks
func (u *User) IsUser() bool {
    return true
}
func (u *User) IsDaemon() bool {
    return false
}

// getting the organisation
func (u *User) GetOrg() *Organisation {
    if u.IsAuthorised() {
        org, _ := GetOrg(u.OrgId)
        return org
    }
    return nil
}
