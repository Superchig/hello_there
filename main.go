package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func main() {
	embeddedSound, err := Asset("hello.wav")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(embeddedSound)

	// soundFile, err := os.Open("hello.wav")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	streamer, format, err := wav.Decode(bytes.NewReader(embeddedSound))
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Hello there.")

	waitTime, _ := time.ParseDuration("3s500ms")
	time.Sleep(waitTime)
	fmt.Println("General Kenobi.")
	reader.ReadString('\n')
}
