package main

import (
	"fmt"
	"ip-finder/app"
	"log"
	"os"
)
func main() {
	fmt.Println("IP Finder")
	app := app.Generate()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}