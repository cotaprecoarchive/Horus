package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/CotaPreco/Horus/command"
	"github.com/CotaPreco/Horus/receiver/udp"
	"github.com/CotaPreco/Horus/util"
	"github.com/CotaPreco/Horus/ws"
	wsc "github.com/CotaPreco/Horus/ws/command"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// @link https://godoc.org/github.com/gorilla/websocket#hdr-Origin_Considerations
		return true
	},
}

func main() {
	flag.Usage = func() {
		flag.CommandLine.SetOutput(os.Stdout)

		fmt.Fprint(os.Stdout, "Usage: horus [OPTIONS] :-)\n\n")

		flag.PrintDefaults()

		os.Exit(0)
	}

	// --
	udpHost := flag.String("receiver-udp-host", "0.0.0.0", "Defines the host IP for `UdpReceiver`")
	udpPort := flag.Int("receiver-udp-port", 7600, "Defines which port `UdpReceiver` will be listening")

	wsHost := flag.String("ws-host", "0.0.0.0", "Where websocket will be available?")
	wsPort := flag.Int("ws-port", 8000, "And in which port people will connect?")

	flag.Parse()
	// --

	bus := command.NewGenericCommandBus()
	hub := ws.NewTaggedConnectionHub()

	bus.PushHandler(hub)
	bus.PushHandler(wsc.NewARTagCommandRedispatcher(bus))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		defer conn.Close()

		if err != nil {
			if _, ok := err.(websocket.HandshakeError); !ok {
				util.Invariant(
					err == nil,
					"...`%s` on attempt to upgrade/handshake connection",
					err,
				)
			}
		}

		hub.Subscribe(conn)

		for {
			messageType, message, err := conn.ReadMessage()

			if err != nil {
				hub.Unsubscribe(conn)
				conn.Close()
				return
			}

			// util.Invariant(
			// 	err == nil,
			// 	"... `%s` on `ReadMessage`",
			// 	err,
			// )

			if messageType == websocket.TextMessage {
				bus.Dispatch(wsc.NewSimpleTextCommand(string(message), conn))
			}
		}
	})

	// ---
	receiver := udp.NewUdpReceiver(*udpHost, *udpPort, new(udp.NullByteReceiveStrategy))
	receiver.Attach(hub)

	go receiver.Receive()
	// ---

	err := http.ListenAndServe(
		fmt.Sprintf("%s:%d", *wsHost, *wsPort),
		nil,
	)

	util.Invariant(
		err == nil,
		"...unexpected `%s` (ListenAndServe)",
		err,
	)
}
