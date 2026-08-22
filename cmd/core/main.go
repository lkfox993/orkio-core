package main

import (
	"log"
	// "os"

	// "github.com/lkfox993/orkio-core/internal/storage/postgres"
	"github.com/lkfox993/orkio-core/internal/storage/tikv"

	httpserver "github.com/lkfox993/orkio-core/internal/transport/http"
)

func main() {

	// _, err := postgres.New(os.Getenv("DATABASE_URL"))

	// if err != nil {
	// 	panic(err)
	// }

	store, err := tikv.New(tikv.Config{
		Endpoints: []string{
			"http://localhost:2379",
		},
	})

	if err != nil {
		panic(err)
	}

	server := httpserver.NewServer()

	if err := server.Start(":8080"); err != nil {
		log.Fatal(err)
	}

	defer store.Close()
}
