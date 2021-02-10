package schema

import (
	"github.com/jmoiron/sqlx"
	"log"
)

//Initializing tables
func DBInit(database *sqlx.DB) {
	i := 1
	for _, ddl := range CreatingDDLs {
		_, err := database.Exec(ddl)
		if err != nil {
			log.Fatalf(" Error while creating table number %d. Error is: %s", i, err)
		}
		i++
	}
}