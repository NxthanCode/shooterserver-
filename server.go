package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},

}

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "negatve punkte")
    })

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "logged in")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		log.Println("new client connected")

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				break
			}
			log.Println("client said:", string(msg))

			err = conn.WriteMessage(websocket.TextMessage, []byte("server said:  "+string(msg)))
			if err != nil {
				log.Println(err)
				break
			}
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	
	log.Println("server ruunning on port", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))

}
