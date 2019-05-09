package conexion

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//GetConexion funcion que retorna una instancia de bd a un error
func GetConexion() (db *sql.DB, e error) {
	usuario := "root"
	password := ""
	host := "tcp(127.0.0.1:3306)"
	nombreBD := "tour_utec"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, password, host, nombreBD))

	if err != nil {
		return nil, err
	}
	return db, nil
}
