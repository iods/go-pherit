package main

import (
	"github.com/thedarksociety/go-pherit/gadget"
)

type Player interface {
	Play(string)
	Stop()
}


/*
TryOut Function to test the various methods of the TapePlayer
and TapeRecorder types. It has a single param and the interface Player
as it's type (pass TapePlayer or TapeRecorder).
*/
func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

func playlist(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
}
