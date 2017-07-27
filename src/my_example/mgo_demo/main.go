package main

import (
	"bytes"
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func checkerr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	checkerr(err)
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	log.Println(bson.Now())
	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	checkerr(err)

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	checkerr(err)

	fmt.Println("Phone:", result.Phone)

	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)

	logger.Print("Hello, log file!")

	// log.SetPrefix("diogo")

	fmt.Print(&buf)
}
