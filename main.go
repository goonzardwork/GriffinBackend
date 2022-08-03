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
	// landing page
	griffin = griffin.
		StartService().
		PingTest().
		Version()

	// employee CRUD operations
	griffin = griffin.
		GetEmployee().
		AddEmployee()

	griffin.Conn.Run()
}
