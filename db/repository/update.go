package repository

import (
	"fmt"
	"log"
	db "projeto-transportadora/db/database"
	"projeto-transportadora/models"
)

func UpdateCleintId(id int64, m models.Cliente) (int64, error) {
	conn, err := db.ConectDBPostgres()
	if err != nil {
		fmt.Println("Error conection")
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE clientinfo
	SET name = $2, address = $3, phone = $4
	WHERE id = $1;`, m.Id, m.Data.Name, m.Data.Address, m.Data.Phone)
	if err != nil {
		fmt.Println("Error in update")
	}
	if res == nil {
		log.Println("pushy user", res)
	}
	return res.RowsAffected()
}
