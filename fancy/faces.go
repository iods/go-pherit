package main

import (
	"github.com/thedarksociety/go-pherit/gadget"
)

type Player interface {
	Play(string)
	Stop()
} // this means both TapePlayer and TapeRecorder


func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	player.Record()
}
func playlist(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	TryOut(gadget.TapeRecorder{})
}
