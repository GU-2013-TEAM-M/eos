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
}

//-------------------------------------------------------
// methods
//-------------------------------------------------------
// obtain Id, OrgId and add yourself to a organisation
func (u *User) Authorise() error {
    return nil
}

// remove itself from the organisation
func (u *User) Deauthorise() error {
    org, err := GetOrg(u.OrgId)
    if err != nil {
        return err
    }
    org.delUser(u.Id)
    return nil
}

