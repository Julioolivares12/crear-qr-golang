package models

import "database/sql"

//Edificio struct de edificio
type Edificio struct {
	IDEdificio int    `json:"id"`
	Nombre     string `json:"nombre"`
}

//GetAllEdificios devulve una lista de edificios
func GetAllEdificios(db *sql.DB) []Edificio {
	sql := "select  id_edificio , nombre from tb_edificio;"
	row, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer row.Close()
	lista := []Edificio{}
	edificio := Edificio{}
	for row.Next() {
		err := row.Scan(&edificio.IDEdificio, &edificio.Nombre)
		if err != nil {
			panic(err)
		}
		lista = append(lista, edificio)
	}
	return lista
}
