package gol

type Params struct {
	Turns       int
	Threads     int
	ImageWidth  int
	ImageHeight int
}
// Run starts the processing of Game of Life. It should initialise channels and goroutines.
func Run(p Params, events chan<- Event, keyPresses <-chan rune) {

	ioCommand := make(chan ioCommand)
	ioIdle := make(chan bool)
	//create a filename channel
	ioFilename := make(chan string)
	ioInput := make(chan uint8)
	ioOutput := make(chan uint8)

	c := distributorChannels{
		events,
		ioCommand,
		ioIdle,
		ioFilename,
		ioInput,
		ioOutput,
	}

	go client(p,c)

	ioChannels := ioChannels{
		command:  ioCommand,
		idle:     ioIdle,
		filename: ioFilename,
		output:   ioOutput,
		input:    ioInput,
	}



	go startIo(p, ioChannels)

}
