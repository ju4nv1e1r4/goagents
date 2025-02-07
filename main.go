package main

import (
	"goagents/app"
	"log"
	"os"
)



func main()  {
	application := app.RunAgents()
	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}