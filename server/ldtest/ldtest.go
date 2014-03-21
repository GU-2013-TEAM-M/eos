package main

import (
    "os"
    "fmt"
    "log"
    "time"
    "bytes"
    "strconv"
    "math/rand"
    "eos/server/db"
    "labix.org/v2/mgo/bson"
    "code.google.com/p/go.net/websocket"
)

// users
var users = make([]*db.User, 0, 100)
// daemons
var daemons = make([]*db.Daemon, 0, 100)
// org2daemon
var o2d = make(map[bson.ObjectId][]*db.Daemon)

func main() {
    db.Connect()
    args := os.Args[1:]

    if len(args) != 3 {
        fmt.Println("Wrong arguments")
        return
    }

    fmt.Println("Starting the simulation.")
    fmt.Println("Organisations:", args[0])
    fmt.Println("Users per org:", args[1])
    fmt.Println("Daemons per org:", args[2])
    fmt.Println()

    // setting up stub db
    norg, err := strconv.Atoi(args[0])
    if err != nil {
        log.Fatal(err)
    }
    nu, err := strconv.Atoi(args[1])
    if err != nil {
        log.Fatal(err)
    }
    nd, err := strconv.Atoi(args[2])
    if err != nil {
        log.Fatal(err)
    }
    setupDBStub(norg, nu, nd)

    // running the simulation
    runSimulation()

    // cleaning up
    cleanTheDB()
}

// populates the database with stub data to match the counts
func setupDBStub(norg, nu, nd int) {
    for i := 0; i < norg; i++ {
        org := &db.Organisation{ Name: strconv.FormatInt(rand.Int63(), 10) }
        db.AddTemp("organisations", org)
        for j := 0; j < nu; j++ {
            generateUser(org)
        }
        ds := make([]*db.Daemon, 0, 100)
        for j := 0; j < nd; j++ {
            ds = append(ds, generateDaemon(org))
        }
        // saving for the future reference:
        o2d[org.Id] = ds
    }
}

// create a user for that organisation
func generateUser(org *db.Organisation) {
    u := &db.User{
        OrgId: org.Id,
        Email: strconv.FormatInt(rand.Int63(), 10),
        Password: strconv.FormatInt(rand.Int63(), 10),
    }
    db.AddTemp("users", u)
    users = append(users, u)
}

// create a daemon for that organisation
func generateDaemon(org *db.Organisation) *db.Daemon {
    d := &db.Daemon{
        OrgId: org.Id,
        Password: strconv.FormatInt(rand.Int63(), 10),
        Name: strconv.FormatInt(rand.Int63(), 10),
        Status: "Running",
        Platform: "Linux",
        Parameters: []string{ "cpu" },
        Monitored: []string{ "cpu" },
    }
    db.AddTemp("daemons", d)
    // generating monitoring data for the daemon
    dataN := rand.Intn(100)
    for i := 0; i < dataN; i++ {
        dp := &db.Data {
            Parameter: "cpu",
            Time: time.Now().Unix() - int64(rand.Intn(1000)),
            Value: float64(rand.Intn(100)),
        }
        db.AddTemp("monitoring_of_" + d.Name, dp)
    }
    daemons = append(daemons, d)
    return d
}

// performs the actual load test
func runSimulation() {
    nu := len(users)
    c := make(chan int)
    var done int

    nd := len(daemons)
    startD := make(chan int)
    stopD := make(chan int, nd)

    fmt.Println("Starting daemons")
    // create all the daemons
    for _, d := range(daemons) {
        go simulateDaemon(d, startD, stopD)
    }
    done = 0
    for done < nd {
        r := <-startD
        if r == 0 {
            fmt.Printf("!")
        } else {
            fmt.Printf(".")
        }
        done += 1
    }
    fmt.Println("\nDaemons have started!")

    // run user simulation
    fmt.Println("Running users")
    for _, u := range(users) {
        go simulateUser(u, c)
    }
    done = 0
    for done < nu {
        r := <-c
        if r == 0 {
            fmt.Printf("!")
        } else {
            fmt.Printf(".")
        }
        done += 1
    }
    fmt.Println("\nUsers done!")

    // stop the daemons
    fmt.Println("Stopping daemons")
    for i := 0; i < nd; i++ {
        stopD <- 1
    }
    fmt.Println("Daemons have been stopped")
}

// removes all the stub data from the database
func cleanTheDB() {
    // delete all the users
    db.DelTemps("users")
    // all the monitoring information
    for _, d := range daemons {
        db.C("monitoring_of_" + d.Name).DropCollection()
    }
    // all the daemons
    db.DelTemps("daemons")
    // all the orgs
    db.DelTemps("organisations")
}

// simulating the user behaviour
func simulateUser(u *db.User, c chan int) {
    // create a websocket connection
    origin := "http://localhost/"
    url := "ws://localhost:8080/wsclient"
    ws, err := websocket.Dial(url, "", origin)
    if err != nil {
        c <- 0
        return
    }

    // login
    loginMsg := `{
        "type": "login",
        "data": {
            "email": "` + u.Email + `",
            "password": "` + u.Password + `"
        }
    }`
    if _, err := ws.Write([]byte(loginMsg)); err != nil {
        c <- 0
        return
    }

    // login response
    var msg = make([]byte, 5096)
    if _, err = ws.Read(msg); err != nil {
        c <- 0
        return
    }
    if !bytes.Contains(msg, []byte("session_id")) {
        c <- 0
        return
    }

    // daemons
    daemonsMsg := `{"type": "daemons", "data": {}}`
    if _, err := ws.Write([]byte(daemonsMsg)); err != nil {
        c <- 0
        return
    }

    // daemons response

    // for all daemons in the org:
        // daemon
        // monitoring

    // close a websocket connection

    c <- 1
}

func simulateDaemon(d *db.Daemon, startD, stopD chan int) {
    // create a websocket connection
    origin := "http://localhost/"
    url := "ws://localhost:8080/wsdaemon"
    ws, err := websocket.Dial(url, "", origin)
    if err != nil {
        startD <- 0
        return
    }

    // daemon just has to log in, in order to be present
    loginMsg := `{
        "type": "login",
        "data": {
            "name": "` + d.Name + `",
            "password": "` + d.Password + `",
            "org_id": "` + d.OrgId.Hex() + `"
        }
    }`
    if _, err := ws.Write([]byte(loginMsg)); err != nil {
        startD <- 0
        return
    }

    // login response
    var msg = make([]byte, 5096)
    if _, err = ws.Read(msg); err != nil {
        startD <- 0
        return
    }
    if !bytes.Contains(msg, []byte(`"id"`)) {
        startD <- 0
        return
    }

    // if we are here, then everything is fine
    startD <- 1

    // wait till the end, and then
    // some unelegant way of stopping the daemon
    schr := <-stopD
    if schr != 0 {
        return
    }
}
