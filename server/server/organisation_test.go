package main

import (
    "testing"
    "eos/server/test"
)

//-------------------------------------------------------
// test handling the organisations
//-------------------------------------------------------
func Test_NewOrg(t *testing.T) {
    orgs = make(map[string]*Organisation)

    orgId := "123"
    org := NewOrg(orgId)

    test.Assert(org != nil, "it creates a new organisation", t)
    test.Assert(len(orgs) == 1, "it is stored in the map of orgs", t)
}

func Test_GetOrg(t *testing.T) {
    orgs = make(map[string]*Organisation)
    orgs["123"] = &Organisation{}

    org, err := GetOrg("123")
    test.Assert(err == nil && org != nil, "it gets an existing organisation", t)

    org, err = GetOrg("nonexistent")
    test.Assert(err != nil && org == nil, "it returns an error if org not found", t)
}

func Test_DelOrg(t *testing.T) {
    orgs = make(map[string]*Organisation)
    orgs["123"] = &Organisation{
        Users: make(map[string]*Connection),
        Daemons: make(map[string]*Connection),
    }

    orgs["123"].Users["test"] = &Connection{}
    err := DelOrg("123")

    test.Assert(err != nil, "it does not delete the empty organisation", t)

    delete(orgs["123"].Users, "test")
    err = DelOrg("123")

    org, err2 := GetOrg("123")
    test.Assert(err == nil && err2 != nil && org == nil, "it deletes the organisation", t)
}

//-------------------------------------------------------
// test user/daemon management
//-------------------------------------------------------
func Test_addUser(t *testing.T) {
    orgs = make(map[string]*Organisation)
    orgs["123"] = &Organisation{
        Users: make(map[string]*Connection),
        Daemons: make(map[string]*Connection),
    }
    org := orgs["123"]

    test.Assert(len(org.Users) == 0, "there are no users initially", t)
    org.addUser("test", &Connection{})
    test.Assert(len(org.Users) == 1, "it adds a user to an organisation", t)

    err := org.addUser("test", &Connection{})
    test.Assert(err != nil, "it doesn't add a duplicate user", t)

}

func Test_delUser(t *testing.T) {
    orgs = make(map[string]*Organisation)
    orgs["123"] = &Organisation{
        Users: make(map[string]*Connection),
        Daemons: make(map[string]*Connection),
    }
    org := orgs["123"]
    org.Users["test"] = &Connection{}

    org.delUser("test")
    test.Assert(len(org.Users) == 0, "it deletes the user", t)
}

func Test_addDaemon(t *testing.T) {
    orgs = make(map[string]*Organisation)
    orgs["123"] = &Organisation{
        Users: make(map[string]*Connection),
        Daemons: make(map[string]*Connection),
    }
    org := orgs["123"]

    test.Assert(len(org.Daemons) == 0, "there are no daemons initially", t)
    org.addDaemon("test", &Connection{})
    test.Assert(len(org.Daemons) == 1, "it adds a daemon to an organisation", t)

    err := org.addDaemon("test", &Connection{})
    test.Assert(err != nil, "it doesn't add a duplicate daemon", t)

}

func Test_delDaemon(t *testing.T) {
    orgs = make(map[string]*Organisation)
    orgs["123"] = &Organisation{
        Users: make(map[string]*Connection),
        Daemons: make(map[string]*Connection),
    }
    org := orgs["123"]
    org.Daemons["test"] = &Connection{}

    org.delDaemon("test")
    test.Assert(len(org.Daemons) == 0, "it deletes the daemon", t)
}
