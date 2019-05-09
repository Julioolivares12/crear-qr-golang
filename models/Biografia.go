package models

import "database/sql"

//Biografia modelo de biografia
type Biografia struct {
	IDBiografia int    `json:"id"`
	Nombre      string `json:"nombre"`
}

//GetAllBiografias debuelve una lista de biografias
func GetAllBiografias(db *sql.DB) []Biografia {
	slq := "select id_biografia,titulo from tb_biografia; "
	row, err := db.Query(slq)
	if err != nil {
		panic(err)
	}
	defer row.Close()
	lista := []Biografia{}
	biografia := Biografia{}
	for row.Next() {
		err := row.Scan(&biografia.IDBiografia, &biografia.Nombre)
		if err != nil {
			panic(err)
		}
		lista = append(lista, biografia)
	}
	return lista
}
