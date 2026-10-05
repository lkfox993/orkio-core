package main

import (
	"fmt"
	// "log"
	"os"

	// "github.com/lkfox993/orkio-core/internal/api"
	// "github.com/lkfox993/orkio-core/internal/storage/postgres"
	// "github.com/lkfox993/orkio-core/internal/storage/tikv"
	// "github.com/lkfox993/orkio-core/internal/transports"
	"github.com/lkfox993/orkio-core/internal/workflow"
)

func main() {

	file, err := os.ReadFile("example/workflow.hcl")

	definition, err := workflow.ParseDefinition(file)

	if err != nil {
		panic(err)
	}

	fmt.Println(333888, definition.Key)
	fmt.Println(definition.Nodes)

	// db, err := postgres.New( /*os.Getenv("DATABASE_URL")*/ "postgres://orkio:password@localhost:5439/orkio")

	// if err != nil {
	// 	panic(err)
	// }

	// store, err := tikv.New(tikv.Config{
	// 	Endpoints: []string{
	// 		"http://localhost:2379",
	// 	},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// httpTransport := transports.NewHttpServer()
	// httpTransport.Configure()

	// api.RegisterModules(httpTransport, db, store)

	// if err := httpTransport.Start(":8080"); err != nil {
	// 	log.Fatal(err)
	// }

	// defer store.Close()
}
