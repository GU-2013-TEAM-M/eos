package db

import (
    "labix.org/v2/mgo/bson"
)

// these are used for testing
// create a temporary entry in the database
func AddTemp(collection string, entry Entry) {
    c := C(collection)
    if entry.GetId().Hex() == "" {
        entry.GenId()
    }
    c.Insert(entry)
    c.UpdateId(entry.GetId(), bson.M{"$set": bson.M{"_tmp": "true"}})
}

// remove all temporary entries from the collection
func DelTemps(collection string) {
    C(collection).RemoveAll(bson.M{"_tmp": "true"})
}
