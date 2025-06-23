package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

var messageTypes = []string{"ping", "update", "status", "event", "log"}
var contents = []string{
	"Temperature: 23.4C",
	"User joined",
	"Ping request",
	"Battery low",
	"Motion detected",
	"File uploaded",
	"Disconnected",
}

func randomMessage() Message {
	return Message{
		Type:    messageTypes[rand.Intn(len(messageTypes))],
		Content: contents[rand.Intn(len(contents))],
	}
}

func clientWorker(id int, serverURL string, wg *sync.WaitGroup) {
	defer wg.Done()

	u := url.URL{Scheme: "ws", Host: serverURL, Path: "/ws"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Printf("❌ [%d] Connection error: %v", id, err)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	log.Printf("✅ [%d] Connected", id)

	for {
		// Генерируем случайный JSON
		msg := randomMessage()
		payload, _ := json.Marshal(msg)

		// Отправка
		if err := conn.WriteMessage(websocket.TextMessage, payload); err != nil {
			log.Printf("❌ [%d] Write error: %v", id, err)
			return
		}
		log.Printf("📤 [%d] Sent: %s", id, payload)

		// Чтение ответа
		_, resp, err := conn.ReadMessage()
		if err != nil {
			log.Printf("❌ [%d] Read error: %v", id, err)
			return
		}
		log.Printf("📩 [%d] Response: %s", id, resp)

		// Ждём случайное время
		sleep := time.Duration(500+rand.Intn(1500)) * time.Millisecond
		time.Sleep(sleep)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const (
		totalClients = 3
		serverAddr   = "localhost:3000"
	)

	var wg sync.WaitGroup
	wg.Add(totalClients)

	log.Printf("🚀 Starting %d clients...\n", totalClients)

	for i := 0; i < totalClients; i++ {
		go clientWorker(i, serverAddr, &wg)
	}

	wg.Wait()
}
