package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket 升级器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许跨域
		return true
	},
}

// WebSocket 处理函数
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()
	log.Println("Client connected")

	// 简单的回声服务
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received: %s\n", message)

		// 回送消息
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func main() {
	// 创建 HTTPS 服务器
	server := &http.Server{
		Addr: ":8081",
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12, // 强制使用 TLS 1.2 或更高版本
		},
	}

	// 路由设置
	http.HandleFunc("/ws", handleWebSocket)

	log.Println("Starting WSS server on wss://localhost:8081/ws")
	err := server.ListenAndServeTLS("cert.pem", "key.pem") // 使用证书启动 HTTPS
	if err != nil {
		log.Fatal("ListenAndServeTLS error:", err)
	}
}
