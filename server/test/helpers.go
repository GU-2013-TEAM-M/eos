package test

import (
    "testing"
    "reflect"
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
// stubbing methods
// taken online from someone called imosquera
//-------------------------------------------------------
// Restorer holds a function that can be used
// to restore some previous state.
type Restorer func()

// Restore restores some previous state.
func (r Restorer) Restore() {
        r()
}

// Patch sets the value pointed to by the given destination to the given
// value, and returns a function to restore it to its original value.  The
// value must be assignable to the element type of the destination.
func Patch(dest, value interface{}) Restorer {
        destv := reflect.ValueOf(dest).Elem()
        oldv := reflect.New(destv.Type()).Elem()
        oldv.Set(destv)
        valuev := reflect.ValueOf(value)
        if !valuev.IsValid() {
                // This isn't quite right when the destination type is not
                // nilable, but it's better than the complex alternative.
                valuev = reflect.Zero(destv.Type())
        }
        destv.Set(valuev)
        return func() {
                destv.Set(oldv)
        }
}
