package main

import (
    "eos/server/db"
    "labix.org/v2/mgo/bson"
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
    Entry db.Daemon
    c *Connection
}

//-------------------------------------------------------
// methods
//-------------------------------------------------------
// obtain all the data and proceed to authorisation
func (d *Daemon) Authenticate(id bson.ObjectId) error {
    c := db.C("daemons")
    err := c.FindId(id).One(&d.Entry)
    if err == nil {
        d.Id = id.Hex()
        d.OrgId = d.Entry.OrgId.Hex()
        return d.Authorise()
    }
    return err
}

// obtain OrgId and add yourself to a organisation
func (d *Daemon) Authorise() error {
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

// type checks
func (d *Daemon) IsUser() bool {
    return false
}
func (d *Daemon) IsDaemon() bool {
    return true
}

// getting the organisation
func (d *Daemon) GetOrg() *Organisation {
    if d.IsAuthorised() {
        org, _ := GetOrg(d.OrgId)
        return org
    }
    return nil
}

