package repository

import (
	"fmt"
	"log"
	db "projeto-transportadora/db/database"
	"projeto-transportadora/models"
)

func InsertClientList(c models.Info) (models.Info, error) {
	conn, err := db.ConectDBPostgres()
	if err != nil {
		fmt.Println("error")
	}
	defer conn.Close()

	sql := `INSERT INTO clientInfo (Name, Address, Phone)
	VALUES ($1, $2, $3)RETURNING id;`
	var id int64

	err = conn.QueryRow(sql, c.Name, c.Address, c.Phone).Scan(&id)
	if err != nil {
		log.Println(err)
		return models.Info{}, err
	}
	log.Println("created client", id)

	return models.Info{Name: c.Name, Address: c.Address, Phone: c.Phone}, nil
}
