package main

import (
    "eos/server/db"
    "labix.org/v2/mgo/bson"
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
    Entry *db.User
    c *Connection
}

//-------------------------------------------------------
// helper functions
//-------------------------------------------------------
// get the id from the session and proceed to authentication
func AuthFromSession(sessId string) (bson.ObjectId, error) {
    c := db.C("sessions")
    result := &db.Session{}
    err := c.FindId(bson.ObjectIdHex(sessId)).One(result)
    return result.UId, err
}

//-------------------------------------------------------
// methods
//-------------------------------------------------------
// obtain all the data and proceed to authorisation
func (u *User) Authenticate(id bson.ObjectId) error {
    c := db.C("users")
    err := c.FindId(id).One(&u.Entry)
    if err == nil {
        u.Id = id.Hex()
        u.OrgId = u.Entry.OrgId.Hex()
        return u.Authorise()
    }
    return err
}

// add yourself to a organisation
func (u *User) Authorise() error {
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
