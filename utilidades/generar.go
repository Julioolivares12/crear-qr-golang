package utilidades

import (
	"fmt"
	"log"

	"github.com/Julioolivares12/CreaQr/models"

	qrcode "github.com/skip2/go-qrcode"
)

//GenerarQREdificios genera codigo QR a partir de una lista
func GenerarQREdificios(edificios []models.Edificio) {
	for _, e := range edificios {
		ip := GetMyLocalIP()
		id := e.IDEdificio
		url := fmt.Sprintf("http://%s:8080/api/Edificios/%d", ip, id)
		edificio := fmt.Sprintf("%s.png", e.Nombre)
		err := qrcode.WriteFile(url, qrcode.Medium, 256, edificio)
		if err != nil {
			log.Fatal(err)
		}
	}
}

//GenerarQRBiografias genera codigo qr de biografias
func GenerarQRBiografias(lista []models.Biografia) {
	for _, bio := range lista {
		ip := GetMyLocalIP()
		id := bio.IDBiografia
		Nombre := fmt.Sprintf("%s.png", bio.Nombre)
		url := fmt.Sprintf("http://%s:8080/api/Biografia/%d", ip, id)
		err := qrcode.WriteFile(url, qrcode.Medium, 256, Nombre)
		if err != nil {
			panic(err)
		}
	}
}
