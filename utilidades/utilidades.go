package utilidades

import "net"

//GetMyLocalIP devuelve de la ip local de la maquina
func GetMyLocalIP() string {
	con, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer con.Close()
	localAddr := con.LocalAddr().(*net.UDPAddr)
	ip := localAddr.IP.String()

	return ip
}
