package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yyoshiki41/go-siwa"
)

func main() {
	client := siwa.NewClient()
	keys, err := client.Keys(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, k := range keys {
		fmt.Printf("%#v\n", k)
	}
}
