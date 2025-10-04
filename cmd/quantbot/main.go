package main

import (
	"log"
	"os"

	"github.com/nathanprogsun/quantbot/quantbot/client"
)

func main() {
    root := client.NewRootCmd()
    if err := root.Execute(); err != nil {
        log.Println(err)
        os.Exit(1)
    }
}
