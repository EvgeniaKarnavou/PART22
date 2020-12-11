package stubs

var NextState = "EngineOperations.CalculateNextState"

var Distributor = "ControllerOperations.Controller"



type Response struct {
	W [][]byte
}



type Request struct {
	W [][]byte
	Param Parameters
}

type Parameters struct{
	Turns int
	Threads     int
	ImageWidth  int
	ImageHeight int
}



