package main

import (
	"GriffinBackend/rest"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}

var griffin = rest.GriffinWS{}

func main() {
	griffin = griffin.
		StartService().
		PingTest().
		Version()

	griffin = griffin.
		GetEmployee()

	griffin.Conn.Run()
}
