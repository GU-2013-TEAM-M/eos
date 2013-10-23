package main

import (
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
var Orgs = make(map[string]*Organisation)

//-------------------------------------------------------
// user/daemon management
//-------------------------------------------------------
// adds an authorised user
func (o *Organisation) addUser(u string, c *Connection) error {
    return nil
}

// removes a user from the list of authorised for this org
func (o *Organisation) delUser(u string) error {
    return nil
}

// adds an authorised daemon
func (o *Organisation) addDaemon(d string, c *Connection) error {
    return nil
}

// removes a daemon from the list of authorised for this org
func (o *Organisation) delDaemon(d string) error {
    return nil
}

//-------------------------------------------------------
// now the following are methods to enhance communication
//-------------------------------------------------------
// send a message to one user
func (o* Organisation) sendToUser(uid, msg string) error {
    return nil
}

// send a message to all users
func (o* Organisation) sendToUsers(msg string) error {
    return nil
}

// send a message to one daemon
func (o* Organisation) sendToDaemon(did, msg string) error {
    return nil
}

// send a message to all daemons
func (o* Organisation) sendToDaemons(msg string) error {
    return nil
}

//-------------------------------------------------------
// handling the organisations
//-------------------------------------------------------
// create the organisation
func NewOrg(orgId string) *Organisation {
    return nil
}

// get the organisation
func GetOrg(orgId string) (*Organisation, error) {
    return nil, nil
}

// delete the organisation
func DelOrg(orgId string) error {
    return nil
}
