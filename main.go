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

var (
	VERSION   = "N/A"
	GITCOMMIT = "N/A"
)

var (
	defaultWsHost           = util.EnvOrDefault("WS_HOST", "0.0.0.0")
	defaultWsPort           = util.EnvOrDefault("WS_PORT", "8000")
	defaultUdpReceiverHost  = util.EnvOrDefault("UDP_RECEIVER_HOST", "0.0.0.0")
	defaultUdpReceiverPort  = util.EnvOrDefault("UDP_RECEIVER_PORT", "7600")
	defaultUdpMaxPacketSize = util.EnvOrDefault("UDP_PACKET_SIZE", "8192")
)

var (
	flgVersion    = flag.Bool("v", false, "")
	udpHost       = flag.String("receiver-udp-host", defaultUdpReceiverHost, "")
	udpPort       = flag.Int("receiver-udp-port", util.Str2int(defaultUdpReceiverPort), "")
	wsHost        = flag.String("ws-host", defaultWsHost, "")
	wsPort        = flag.Int("ws-port", util.Str2int(defaultWsPort), "")
	udpPacketSize = flag.Int("udp-max-packet-size", util.Str2int(defaultUdpMaxPacketSize), "")
)

func main() {
	flag.Usage = func() {
		flag.CommandLine.SetOutput(os.Stdout)

		var help = strings.Trim(`
Horus — A simple and minimalist event-hub for pipelining events :-)

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
			}, {
				"-udp-max-packet-size",
				"Defines the maximum buffer size for packet",
			},
		} {
			opts += fmt.Sprintf("\t%-25.20s /* %s */\n", opt[0], opt[1])
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

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		connection, err := ws.WebSocketUpgrader.Upgrade(writer, request, nil)

		if err != nil {
			return
		}

		defer connection.Close()

		hub.Subscribe(connection)

		defer hub.Unsubscribe(connection)

		for {
			_type, message, err := connection.ReadMessage()

			if err != nil {
				return
			}

			if _type == websocket.TextMessage {
				bus.Dispatch(wsc.NewSimpleTextCommand(string(message), connection))
			}
		}
	})

	receiver := udp.NewUdpReceiver(*udpHost, *udpPort, *udpPacketSize, new(udp.NullByteReceiveStrategy))
	receiver.Attach(hub)

	go receiver.Receive()

	fmt.Printf(
		"Udp Receiver — %s:%d\nWebSocket — %s:%d\n",
		*udpHost,
		*udpPort,
		*wsHost,
		*wsPort,
	)

	err := http.ListenAndServe(
		fmt.Sprintf("%s:%d", *wsHost, *wsPort),
		nil,
	)

	if err != nil {
		fmt.Print(err)
	}
}
