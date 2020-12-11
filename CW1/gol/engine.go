
package gol

import (
	"uk.ac.bris.cs/gameoflife/stubs"
)



type EngineOperations struct {}



//the function  accepts requests from the clients and is responsible for evolving the Game Of Life logic.
func (s *EngineOperations) CalculateNextState(req stubs.Request, res *stubs.Response) (err error) {


	res.W = make([][]byte, req.Param.ImageHeight)
	for i := range res.W {
		res.W[i] = make([]byte, req.Param.ImageWidth)
	}
	//fmt.Println(res.W)
	res.W = distributor(req.Param, req.W)



	return
}
