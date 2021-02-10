package schema

import (
	"github.com/jmoiron/sqlx"
	"log"
)

//Dropping tables
func DBDrop(database *sqlx.DB) {
	i := 1
	for _, ddl := range DroppingDDLs {
		_, err := database.Exec(ddl)
		if err != nil {
			log.Fatalf(" Error while dropping table number %d. Error is: %s", i, err)
		}
	}
	i++
}
