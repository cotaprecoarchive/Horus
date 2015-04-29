package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

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

const (
	VERSION = "0.1.0-beta"
)

func main() {
	flag.Usage = func() {
		flag.CommandLine.SetOutput(os.Stdout)

		fmt.Fprintf(os.Stdout, "Horus v.%s\nUsage: horus [...OPTIONS] :-)\n\n", VERSION)
		flag.PrintDefaults()

		os.Exit(0)
	}

	// --
	// TODO: ...remove setup from main, encapsulate it
	udpHost := flag.String(
		"receiver-udp-host",
		util.EnvOrDefault("UDP_RECEIVER_HOST", "0.0.0.0"),
		"Defines the host IP for `UdpReceiver`",
	)

	udpReceiverPort, _ := strconv.Atoi(util.EnvOrDefault("UDP_RECEIVER_PORT", "7600"))

	udpPort := flag.Int(
		"receiver-udp-port",
		udpReceiverPort,
		"Defines which port `UdpReceiver` will be listening",
	)

	wsHost := flag.String(
		"ws-host",
		util.EnvOrDefault("WS_HOST", "0.0.0.0"),
		"Where websocket will be available?",
	)

	wsDefaultPort, _ := strconv.Atoi(util.EnvOrDefault("WS_PORT", "8000"))

	wsPort := flag.Int(
		"ws-port",
		wsDefaultPort,
		"And in which port people will connect?",
	)

	flag.Parse()
	// --

	bus := command.NewGenericCommandBus()
	hub := ws.NewTaggedConnectionHub()

	bus.PushHandler(hub)
	bus.PushHandler(wsc.NewARTagCommandRedispatcher(bus))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			if _, ok := err.(websocket.HandshakeError); !ok {
				util.Invariant(
					err == nil,
					"...`%s` on attempt to upgrade/handshake connection",
					err,
				)
			}
		}

		defer conn.Close()

		hub.Subscribe(conn)

		for {
			messageType, message, err := conn.ReadMessage()

			if err != nil {
				hub.Unsubscribe(conn)
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
