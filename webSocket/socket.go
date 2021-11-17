package socket

import (
	"fmt"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

func Socket() *socketio.Server {

	// server, err := socketio.NewServer(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// server.On("connection", func(so socketio.Socket) {
	// 	log.Println("New connection")

	// 	so.Join("chat")

	// 	so.On("message", func(msg string) {
	// 		// return that message
	// 		so.Emit("message", msg)

	// 		so.BroadcastTo("chat", "message", msg)
	// 	})

	// 	so.On("disconnection", func() {
	// 		log.Println("Connection closed")
	// 	})
	// })

	// server.On("error", func(so socketio.Socket, err error) {
	// 	log.Println("error:", err)
	// })

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		s.Join("bcast")
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		server.BroadcastToRoom("", "bcast", "reply", msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			fmt.Println(err)
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()


	return server
}