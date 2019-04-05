package main

import (

	"net/http"
	"github.com/vavas/GO-Szondi/types"
	"log"
)

// EngineService service to manage basic user information for global filter
type EngineService int

// Get basic engine information
func (t *EngineService) Get(r *http.Request, args *types.Args, reply *types.BasicInfo) error {

	log.Println(args)
	result := 1
	reply.ID = result

	return nil
}