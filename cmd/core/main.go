package main

import (
	// "context"

	// "log"
	"context"
	"fmt"
	"os"

	// "github.com/lkfox993/orkio-core/internal/api"
	// "github.com/lkfox993/orkio-core/internal/app"
	// "github.com/lkfox993/orkio-core/internal/transports"
	"github.com/lkfox993/orkio-core/internal/app"
	"github.com/lkfox993/orkio-core/internal/workflow"
)

func main() {

	file, err := os.ReadFile("example/workflow.hcl")

	definition, err := workflow.ParseDefinition(file)

	if err != nil {
		panic(err)
	}

	fmt.Println(definition.Key, definition.Nodes[0].GetChildren()[0].NodeType())

	ctx := context.Background()
	app.New(ctx)

	// httpTransport := transports.NewHttpServer()
	// httpTransport.Configure()

	// api.RegisterModules(httpTransport)

	// if err := httpTransport.Start(":8080"); err != nil {
	// 	log.Fatal(err)
	// }

	// defer store.Close()
}
