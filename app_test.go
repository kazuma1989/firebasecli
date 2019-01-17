package firebasecli_test

import (
	"context"
	"fmt"
	"log"

	"github.com/kazuma1989/firebasecli"
)

func ExampleApp_Login() {
	app := &firebasecli.App{}
	app.Login(context.Background(), "/path/to/key.json")
}

func ExampleApp_DbList() {
	app := &firebasecli.App{}
	app.Login(context.Background(), "/path/to/key.json")

	collections, err := app.DbList(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(collections)
}
