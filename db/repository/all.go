package repository

import (
	db "projeto-transportadora/db/database"
	"projeto-transportadora/models"
)

func GetAllClient(id int64) (clie []models.Cliente, err error) {
	conn, err := db.ConectDBPostgres()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`select id, "name", address, phone, createdat from clientInfo WHERE id = $1;`, id)
	if err != nil {
		return
	}

	for rows.Next() {
		var m models.Cliente
		err = rows.Scan(
			&m.Id,
			&m.Data.Name,
			&m.Data.Address,
			&m.Data.Phone,
			&m.Date.Created,
		)
		clie = append(clie, m)
	}

	return clie, nil
}
