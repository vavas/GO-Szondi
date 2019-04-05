package main

import (

	"log"
	"net/http"
	"os"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"github.com/rs/cors"

)

var (
	gStatus chan string
)

//setSetupChannel set setup channel to use for server status
func setSetupChannel(status chan string) chan string {
	gStatus = status
	return gStatus
}

// Exception custom json exception
type Exception struct {
	Message string `json:"message"`
}

//getSetupChannel get setup channel for server status
func getSetupChannel() chan string {
	return gStatus
}

func main() {

	log.Println("Project ", os.Getenv("PROJECT_NAME"))
	log.Println("Loading services...")

	es := new(EngineService)

	log.Println("Attaching services...")

	// rpc server for project endpoint
	s := rpc.NewServer()
	s.RegisterService(es, "engine")
	s.RegisterCodec(json.NewCodec(), "application/json")

	allowedHeaders := []string{"Authorization", "Content-Type"}
	c := cors.New(cors.Options{
		AllowedOrigins:   getConfig().CorsOrigins,
		AllowCredentials: true,
		AllowedHeaders:   allowedHeaders,
	})

	http.Handle("/rpc", c.Handler(s))

	ch := getSetupChannel()
	if ch != nil {
		ch <- "done"
	}

	log.Println("Starting listener...")
	log.Println("Listening on port 80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}



