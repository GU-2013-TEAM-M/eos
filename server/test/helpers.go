package test

import (
    "testing"
    "fmt"
)

//-------------------------------------------------------
// color constants
//-------------------------------------------------------
const CLR_0 = "\x1b[30;1m"
const CLR_R = "\x1b[31;1m"
const CLR_G = "\x1b[32;1m"
const CLR_Y = "\x1b[33;1m"
const CLR_B = "\x1b[34;1m"
const CLR_M = "\x1b[35;1m"
const CLR_C = "\x1b[36;1m"
const CLR_W = "\x1b[37;1m"
const CLR_N = "\x1b[0m"

//-------------------------------------------------------
// jUnit
//-------------------------------------------------------
func Assert(assertion bool, msg string, t *testing.T) {
    if assertion {
        t.Log(fmt.Sprintf("%s%s%s", CLR_G, msg, CLR_N))
    } else {
        t.Error(fmt.Sprintf("%sFailed: %s%s", CLR_R, msg, CLR_N))
    }
}

//-------------------------------------------------------
// rspec
//-------------------------------------------------------
// TODO: think about how to port it to go.
var curT *testing.T
var ind int = 0
var indW int = 4

// returns a string of spaces, for indentation purposes
// FIXME: write a proper implementation, if you have time
func indent() string {
    s := ""
    for i := 0; i < ind * indW; i++ {
        s += " "
    }
    return s
}

// describe block, sets the indentation
func Describe(msg string, t *testing.T) {
    curT = t
    t.Log(indent() + msg)
}

// it block
