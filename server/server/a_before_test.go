package main

import (
    "testing"
    "eos/server/db"
)

// things to do before running all the tests
func Test_before(t *testing.T) {
    db.Connect()
}

