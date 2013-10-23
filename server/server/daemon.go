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
}

//-------------------------------------------------------
// methods
//-------------------------------------------------------
// obtain OrgId and add yourself to a organisation
func (d *Daemon) Authorise() error {
    return nil
}

// remove itself from the organisation
func (d *Daemon) Deauthorise() error {
    org, err := GetOrg(d.OrgId)
    if err != nil {
        return err
    }
    return org.delDaemon(d.Id)
}

