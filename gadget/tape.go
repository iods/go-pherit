package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

type TapeRecorder struct {
	Microphone int
}

func (t TapePlayer) Play(song string) {
	fmt.Printf("Playing %s.\n", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped.")
}

func (t TapeRecorder) Play(song string) {
	fmt.Printf("Playing %s.\n", song)
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording.")
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped.")
}