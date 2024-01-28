package main

import packetv1 "github.com/wcygan/kafka-on-kubernetes/generated/go/packet/v1"

func main() {
	packet := packetv1.Packet{Number: 1}
	println("Consuming packet: ", packet.Number)
}
