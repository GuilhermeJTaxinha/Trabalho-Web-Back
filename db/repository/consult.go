package repository

import (
	"log"
	db "projeto-transportadora/db/database"
	"projeto-transportadora/models"
)

func GetClient() (c []models.Cliente, err error) {
	var respclient []models.Cliente
	conn, err := db.ConectDBPostgres()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`select * from clientInfo;`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var m models.Cliente
		if err := rows.Scan(
			&m.Id,
			&m.Data.Name,
			&m.Data.Address,
			&m.Data.Phone,
			&m.Date.Created,
		); err != nil {
			log.Printf("repository error: GetClient row scan error: %s", err.Error())
		}
		respclient = append(respclient, m)
	}

	return respclient, nil
}
