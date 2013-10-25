package main

import (
    "errors"
)

//-------------------------------------------------------
// data structures
//-------------------------------------------------------
// a data structure, used to establish connections withing an organisation
// it maps ids to connections, which have User/Daemon objects internally
type Organisation struct {
    Users map[string]*Connection
    Daemons map[string]*Connection
}

// an interface for Users and Daemons
type Organisable interface {
    Authorise() error
    Deauthorise() error
}

// actual map of organisations
var orgs = make(map[string]*Organisation)

// constant for no organisation
const NO_ORG = ""

//-------------------------------------------------------
// handling the organisations
//-------------------------------------------------------
// create the organisation
func NewOrg(orgId string) *Organisation {
    org := &Organisation{
        Users: make(map[string]*Connection),
        Daemons: make(map[string]*Connection),
    }
    orgs[orgId] = org
    return org
}

// get the organisation
func GetOrg(orgId string) (*Organisation, error) {
    if org, ok := orgs[orgId]; ok == true {
        return org, nil
    }
    return nil, errors.New("Organisation not found")
}

// delete the organisation
func DelOrg(orgId string) error {
    org, err := GetOrg(orgId)
    if err != nil {
        return errors.New("Organisation not found")
    }

    if len(org.Users) > 0 || len(org.Daemons) > 0 {
        return errors.New("Deleting non empty organisation")
    }

    delete(orgs, orgId)

    return nil
}

//-------------------------------------------------------
// user/daemon management
//-------------------------------------------------------
// adds an authorised user
func (o *Organisation) addUser(u string, c *Connection) error {
    if _, ok := o.Users[u]; ok {
        return errors.New("Such user already exists")
    }

    o.Users[u] = c
    return nil
}

// removes a user from the list of authorised for this org
func (o *Organisation) delUser(u string) {
    delete(o.Users, u)
}

// adds an authorised daemon
func (o *Organisation) addDaemon(d string, c *Connection) error {
    if _, ok := o.Daemons[d]; ok {
        return errors.New("Such daemon already exists")
    }

    o.Daemons[d] = c
    return nil
}

// removes a daemon from the list of authorised for this org
func (o *Organisation) delDaemon(d string) {
    delete(o.Daemons, d)
}

//-------------------------------------------------------
// now the following are methods to enhance communication
//-------------------------------------------------------
// send a message to one user
func (o* Organisation) sendToUser(uId, msg string) error {
    user, ok := o.Users[uId]
    if !ok {
        return errors.New("User not found")
    }

    user.send <- msg
    return nil
}

// send a message to all users
func (o* Organisation) sendToUsers(msg string) {
    for _, user := range o.Users {
        go func(user *Connection) {
            user.send <- msg
        } (user)
    }
}

// send a message to one daemon
func (o* Organisation) sendToDaemon(dId, msg string) error {
    daemon, ok := o.Daemons[dId]
    if !ok {
        return errors.New("User not found")
    }

    daemon.send <- msg
    return nil
}

// send a message to all daemons
func (o* Organisation) sendToDaemons(msg string) {
    for _, daemon := range o.Daemons {
        go func(daemon *Connection) {
            daemon.send <- msg
        } (daemon)
    }
}
