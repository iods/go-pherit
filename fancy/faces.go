package main

import (
	"github.com/thedarksociety/go-pherit/gadget"
)

type Player interface {
	Play(string)
	Stop()
} // this means both TapePlayer and TapeRecorder

func playlist(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	mixtape := []string{"Aenima", "Everlong", "Hallelujah", "Wish You Were Here"}
	var player Player = gadget.TapePlayer{}
	playlist(player, mixtape)
	player = gadget.TapeRecorder{}
	playlist(player, mixtape)
}
