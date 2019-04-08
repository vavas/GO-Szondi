package main

import (
	"github.com/vavas/Szondi-Aggregator/types"
	"net/http"
)

// EngineService service to manage basic user information for global filter
type EngineService int

// Calculate calculate interpretation based on vectors
func (t *EngineService) Calculate(r *http.Request, args *types.Vectors, reply *types.Interpretation) error {

	result := new(types.Interpretation)
	result.ID = 1
	result.Text = "text"

	*reply = *result

	return nil
}