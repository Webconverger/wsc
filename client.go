package main

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	id := "https://webconverger.com/" + os.Getenv("webc_id")
	pond := "ws://ws.webconverger.com/fish"
	var msg = make([]byte, 512)

	var err error
	var ws *websocket.Conn
	failcount := 1

	args := os.Args

	cmd := exec.Command(args[1], args[2:]...)
	var b bytes.Buffer
	cmd.Stdout = &b

	go func() {
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	for {
		ws, err = websocket.Dial(pond, "", id)
		if err != nil {
			log.Println("Connection failed, re-trying ", failcount)
			failcount++
			time.Sleep(5 * time.Second)
			continue
		}
		log.Printf("Connected to %s", pond)

		n, err := ws.Read(msg)

		if err != nil {
			log.Println("Error reading", err)
		}

		log.Printf("Received: %s\n", msg)
		rurl := string(msg[:n])

		u, err := url.ParseRequestURI(rurl)
		if err == nil {
			switch u.Scheme {
			case "http", "https":
				ws.Close()
				fmt.Println(u)
				if err := cmd.Process.Kill(); err != nil {
					log.Fatal("failed to kill: ", err)
				}
				os.Exit(0)
			default:
				log.Println("Non-URL returned:", rurl)
			}
		}

		time.Sleep(1 * time.Second)
	}
}
