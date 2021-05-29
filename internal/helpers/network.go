package helpers

import (
	"log"
	"net"
)

// GetLocalIP Funció auxiliar que retorna l'adreça IP de l'equip on s'està executant el codi
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Print("Error al obtenir les adreces IPs locals", err.Error())
		return "", err
	}

	for _, a := range addrs {
		if ip, ok := a.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				ipAddress := ip.IP.String()
				log.Printf("Indentificada l'adreça IP %s en el host local", ipAddress)
				return ipAddress, nil
			}
		}
	}
	return "", nil
}