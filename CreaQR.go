package main

import (
	"github.com/Julioolivares12/CreaQr/conexion"
	"github.com/Julioolivares12/CreaQr/models"
	"github.com/Julioolivares12/CreaQr/utilidades"
)

func main() {

	db, err := conexion.GetConexion()
	if err != nil {
		panic(err)
	}
	lista := models.GetAllEdificios(db)
	utilidades.GenerarQREdificios(lista)
	lista2 := models.GetAllBiografias(db)
	utilidades.GenerarQRBiografias(lista2)
	//err := qrcode.Encode(content,qrcode.Medium , 256)
	//err := qrcode.WriteFile(content, qrcode.High, 256, edificio)
}
