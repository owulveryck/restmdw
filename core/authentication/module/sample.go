package main

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/owulveryck/restmdw/services/models"
	zmq "github.com/pebbe/zmq4"
	"golang.org/x/crypto/bcrypt"
	"strconv"

	"encoding/json"
	"log"
)

// Sample authentication module
func main() {
	//  Socket to talk to clients
	responder, _ := zmq.NewSocket(zmq.REP)
	defer responder.Close()
	responder.Bind("tcp://localhost:5555")

	for {
		//  Wait for next request from client
		msg, _ := responder.Recv(0)
		log.Println("Received ", msg)
		var user models.User
		err := json.Unmarshal([]byte(msg), &user)
		if err != nil {
			log.Printf("Error: msg %v is not a json encoded username\n", msg)
		}
		//  Do some 'work'
		//time.Sleep(time.Second)

		//  Send reply back to client
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), 10)

		testUser := models.User{
			UUID:     uuid.New(),
			Username: "test",
			Password: string(hashedPassword),
		}

		reply := user.Username == testUser.Username && bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)) == nil
		responder.Send(strconv.FormatBool(reply), 0)
	}
}
