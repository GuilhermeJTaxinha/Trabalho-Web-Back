package repository

import (
	db "projeto-transportadora/db/database"
	"projeto-transportadora/models"
)

func DeleteClientId(id int64) (c models.Cliente, err error) {
	conn, err := db.ConectDBPostgres()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`DELETE FROM clientinfo WHERE id = $1;`, id)

	err = row.Scan(id)

	return models.Cliente{}, err
}
