package main

import (
	"fmt"
	"time"

	webrtc "github.com/pion/webrtc/v3"
	signal "github.com/yurikoex/crobloc/src/signal"
)

type Block struct {
	Index     int64  `json:i`
	Previous  string `json:p`
	Timestamp int64  `json:ts`
	Data      string `json:d`
}

type Peer struct {
	Address        string
	PeerConnection webrtc.PeerConnection
	Connected      bool
}

func main() {
	//https://github.com/pion/webrtc/blob/master/examples/data-channels/main.go
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	// Set the handler for ICE connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", connectionState.String())
	})

	// Register data channel creation handling
	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
		fmt.Printf("New DataChannel %s %d\n", d.Label(), d.ID())

		// Register channel opening handling
		d.OnOpen(func() {
			fmt.Printf("Data channel '%s'-'%d' open. Random messages will now be sent to any connected DataChannels every 5 seconds\n", d.Label(), d.ID())

			for range time.NewTicker(5 * time.Second).C {
				message := signal.RandSeq(15)
				fmt.Printf("Sending '%s'\n", message)

				// Send the message as text
				sendErr := d.SendText(message)
				if sendErr != nil {
					panic(sendErr)
				}
			}
		})

		// Register text message handling
		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("Message from DataChannel '%s': '%s'\n", d.Label(), string(msg.Data))
		})
	})

	//connect to known peers
	connectToPeers()

}

func connectToPeers() {

}

func broadcastGetBlocks() {

}

func hash(str string) string {

	return ""
}

func calcHash(hash string) string {

	return ""
}
