package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// Message - структура входящего и исходящего JSON-сообщения
type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true, // Запускает несколько процессов
	})

	// Middleware: проверка на WebSocket upgrade
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// WebSocket обработчик
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		log.Println("🔌 WebSocket Connected")
		defer func() {
			log.Println("❌ Disconnected")
			err := c.Close()
			if err != nil {
				return
			}
		}()

		for {
			// Читаем raw-сообщение от клиента
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				break
			}

			// Парсим JSON
			var received Message
			if err := json.Unmarshal(msg, &received); err != nil {
				log.Println("❌ Invalid JSON:", err)
				continue
			}

			log.Printf("📨 Received: %+v\n", received)

			// Формируем ответ
			response := Message{
				Type:    "response",
				Content: "Hello, you said: " + received.Content,
			}

			// Сериализуем JSON и отправляем
			respBytes, _ := json.Marshal(response)
			if err := c.WriteMessage(websocket.TextMessage, respBytes); err != nil {
				log.Println("Write error:", err)
				break
			}
		}
	}))

	log.Fatal(app.Listen(":3000"))
}

// {"type": "greeting", "content": "Привет!"}
