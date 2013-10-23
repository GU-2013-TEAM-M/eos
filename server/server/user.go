package main

import (
//    "eos/server/db"
)

// a user entity, tied to a connection
// is anonymous if only session id is supplied
type User struct {
    Id string
    OrgId string
    SessionId string
}

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
    return org.delUser(u.Id)
}

