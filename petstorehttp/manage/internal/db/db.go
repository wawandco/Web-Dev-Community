package db

import (
	"log"

	"github.com/gobuffalo/pop/v6"
)

var Connection *pop.Connection

func init() {
	pop.AddLookupPaths("../../..")

	if err := pop.LoadConfigFile(); err != nil {
		log.Panic(err)
	}

	conn, err := pop.Connect("")
	if err != nil {
		log.Panic(err)
	}

	Connection = conn
	log.Println("Conected to database")
}
