package main

import (
    "testing"
    "eos/server/test"
)

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

func Test_u_IsAuthorised(t *testing.T) {
    u := &User{OrgId:NO_ORG}
    test.Assert(!u.IsAuthorised(), "it spots unauthorised person", t)
    u.OrgId = "random"
    test.Assert(u.IsAuthorised(), "it tells when a person is authorised", t)
}

func Test_u_TypeChecks(t *testing.T) {
    u := &User{OrgId:NO_ORG}
    test.Assert(u.IsUser(), "it thinks that it is a user", t)
    test.Assert(!u.IsDaemon(), "it does not think that is is a daemon", t)
}

func Test_u_GetOrg(t *testing.T) {
    org := setupOrg()

    u := &User{OrgId:NO_ORG}
    test.Assert(u.GetOrg() == nil, "it does not return an org if the daemon is not authorised", t)

    u = &User{OrgId: "123"}
    test.Assert(u.GetOrg() == org, "it does return an org if the daemon is authorised", t)
}
