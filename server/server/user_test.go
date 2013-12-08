package main

import (
    "testing"
    "eos/server/test"
    "eos/server/db"
    "labix.org/v2/mgo/bson"
)

//-------------------------------------------------------
// get the id from the session
//-------------------------------------------------------
func Test_u_AuthFromSession(t *testing.T) {
    tmp := &db.Session{UId: bson.ObjectIdHex("52a4ed348350a921bd000002")}
    db.AddTemp("sessions", tmp)

    uid, err := AuthFromSession(tmp.Id.Hex())

    test.Assert(uid.Hex() == "52a4ed348350a921bd000002", "It finds the user when he is in the session", t)
    test.Assert(err == nil, "It does not throw an error then", t)

    db.DelTemps("sessions")
    uid, err = AuthFromSession(tmp.Id.Hex())

    test.Assert(uid.Hex() != "52a4ed348350a921bd000002", "It does not find the user that is not in the session", t)
    test.Assert(err != nil, "It does throw an error then", t)
}

//-------------------------------------------------------
// obtain all the data and proceed to authorisation
//-------------------------------------------------------
func Test_u_Authenticate(t *testing.T) {
    tmpU := &db.User{OrgId: bson.ObjectIdHex("52a4ed348350a921bd000002"), Email: "a", Password: "b"}
    db.AddTemp("users", tmpU)

    u := &User{}
    err := u.Authenticate(tmpU.Id)

    test.Assert(u.OrgId == "52a4ed348350a921bd000002", "It authenticates user if he exists in the database", t)
    test.Assert(err == nil, "And it does not throw an error in that case", t)

    db.DelTemps("users")

    u = &User{}
    err = u.Authenticate(tmpU.Id)

    test.Assert(u.OrgId != "52a4ed348350a921bd000002", "It fails to recognise  user if he does not exist in the database", t)
    test.Assert(err != nil, "It throws an error in that case", t)
}

//-------------------------------------------------------
// add yourself to a organisation
//-------------------------------------------------------
func Test_u_Authorise_newOrg(t *testing.T) {
    setupOrg()
    org, err := GetOrg("Anonymous")
    test.Assert(err != nil, "organisation does not exist", t)

    u := &User{Id: "test", OrgId: "Anonymous"}

    u.Authorise()
    org, err = GetOrg("Anonymous")

    test.Assert(org != nil, "It creates a new org", t)
    test.Assert(len(org.Users) == 1, "It adds user to organisation", t)
    test.Assert(u.OrgId == "Anonymous", "It stores org ID in user", t)
}

func Test_u_Authorise_exOrg(t *testing.T) {
    setupOrg()
    org := NewOrg("Anonymous")
    test.Assert(len(org.Users) == 0, "there are no daemons in the org initially", t)

    u := &User{Id: "test", OrgId: "Anonymous"}

    u.Authorise()

    test.Assert(len(org.Users) == 1, "It adds user to organisation", t)
    test.Assert(u.OrgId == "Anonymous", "It stores org ID in user", t)
}

//-------------------------------------------------------
// remove itself from the organisation
//-------------------------------------------------------
func Test_u_Deauthorise(t *testing.T) {
    org := setupOrg()
    org.Users["test"] = &Connection{}

    u := &User{Id: "test", OrgId: NO_ORG}

    err := u.Deauthorise()
    test.Assert(err != nil, "it does not deauthorise already deauthorised user", t)

    u.OrgId = "123"
    err = u.Deauthorise()
    test.Assert(err == nil && len(org.Users) == 0, "it removes a user from organisation", t)
    test.Assert(u.OrgId == NO_ORG, "it sets the user organisation to none", t)
}

//-------------------------------------------------------
// check if a user is authorised
//-------------------------------------------------------
func Test_u_IsAuthorised(t *testing.T) {
    u := &User{OrgId:NO_ORG}
    test.Assert(!u.IsAuthorised(), "it spots unauthorised person", t)
    u.OrgId = "random"
    test.Assert(u.IsAuthorised(), "it tells when a person is authorised", t)
}

//-------------------------------------------------------
// type checks
//-------------------------------------------------------
func Test_u_TypeChecks(t *testing.T) {
    u := &User{OrgId:NO_ORG}
    test.Assert(u.IsUser(), "it thinks that it is a user", t)
    test.Assert(!u.IsDaemon(), "it does not think that is is a daemon", t)
}

//-------------------------------------------------------
// getting the organisation
//-------------------------------------------------------
func Test_u_GetOrg(t *testing.T) {
    org := setupOrg()

    u := &User{OrgId:NO_ORG}
    test.Assert(u.GetOrg() == nil, "it does not return an org if the daemon is not authorised", t)

    u = &User{OrgId: "123"}
    test.Assert(u.GetOrg() == org, "it does return an org if the daemon is authorised", t)
}
