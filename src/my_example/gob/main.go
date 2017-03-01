package main

import (
    "encoding/gob"
    "os"
)

type User struct {
    Username string
    Password string 
}

func main() {

    user := User{
        "zola",
        "supersecretpassword",
    }

    file, _ := os.Create("user.gob")

    defer file.Close()

    encoder := gob.NewEncoder(file)

    encoder.Encode(user)



}