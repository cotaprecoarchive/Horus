package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

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

var (
	VERSION   = "0.1.0-beta"
	GITCOMMIT = "N/A"
)

var (
	defaultWsHost          = util.EnvOrDefault("WS_HOST", "0.0.0.0")
	defaultWsPort          = util.EnvOrDefault("WS_PORT", "8000")
	defaultUdpReceiverHost = util.EnvOrDefault("UDP_RECEIVER_HOST", "0.0.0.0")
	defaultUdpReceiverPort = util.EnvOrDefault("UDP_RECEIVER_PORT", "7600")
)

var (
	flgVersion = flag.Bool("v", false, "")
	udpHost    = flag.String("receiver-udp-host", defaultUdpReceiverHost, "")
	udpPort    = flag.Int("receiver-udp-port", util.Str2int(defaultUdpReceiverPort), "")
	wsHost     = flag.String("ws-host", defaultWsHost, "")
	wsPort     = flag.Int("ws-port", util.Str2int(defaultWsPort), "")
)

func main() {
	flag.Usage = func() {
		flag.CommandLine.SetOutput(os.Stdout)

		var help = strings.Trim(`
Horus — An event-hub for pipelining events from any direction to the client :-)

USAGE:
	horus [...OPTIONS]

OPTIONS:
%s
`, "\n")

		var opts string

		for _, opt := range [][]string{
			{
				"-v",
				"Prints the current version of `Horus`",
			}, {
				"-ws-host",
				"Defines in which IP WebSocket will bind to",
			}, {
				"-ws-port",
				"Defines the port for the WebSocket server listen for connections",
			}, {
				"-receiver-udp-host",
				"Defines in which IP the UDP receiver will bind to",
			}, {
				"-receiver-udp-port",
				"Defines the port for receiver listen on",
			},
		} {
			opts += fmt.Sprintf("\t%-18.20s /* %s */\n", opt[0], opt[1])
		}

		fmt.Printf(help, opts)

		os.Exit(0)
	}

	flag.Parse()

	if *flgVersion {
		fmt.Printf("Horus v%s, build %s\n", VERSION, GITCOMMIT)
		return
	}

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
