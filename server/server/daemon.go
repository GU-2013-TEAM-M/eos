package main

import (
//    "eos/server/db"
)

//-------------------------------------------------------
// data structures
//-------------------------------------------------------
// a daemon entity, tied to a connection
// is anonymous if only OrgId and IP are supplied
type Daemon struct {
    Id string
    IP string
    OrgId string
    status int
    c *Connection
}

//-------------------------------------------------------
// methods
//-------------------------------------------------------
// obtain OrgId and add yourself to a organisation
func (d *Daemon) Authorise() error {
    // TODO: getting organisation id from MongoDB
    orgId := "Anonymous"
    d.OrgId = orgId

    org, err := GetOrg(d.OrgId)
    // if the organisation does not exist -- create one
    if err != nil {
        org = NewOrg(d.OrgId)
    }
    org.addDaemon(d.Id, d.c)
    return nil
}

// remove itself from the organisation
func (d *Daemon) Deauthorise() error {
    org, err := GetOrg(d.OrgId)
    if err != nil {
        return err
    }
    org.delDaemon(d.Id)
    d.OrgId = NO_ORG
    return nil
}

// check if a daemon is authorised
func (d *Daemon) IsAuthorised() bool {
    return d.OrgId != NO_ORG
}

