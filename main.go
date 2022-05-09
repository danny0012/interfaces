package main

import (
	_ "fmt"
	"interfaces/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
	
}