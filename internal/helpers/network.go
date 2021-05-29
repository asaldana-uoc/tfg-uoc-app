package helpers

import (
	"log"
	"net"
)

// GetLocalIP Funció auxiliar que retorna l'adreça IP de l'equip on s'està executant el codi
func GetLocalIP() (string, error) {
	// Busquem totes les adreces IP del sistema, en cas d'error (permisos, etc..) registem l'error
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Print("Error al obtenir les adreces IPs locals", err.Error())
		return "", err
	}

	// Iterem sobre totes les adreces IP trobades en el sistema i s'omet l'adreça IP de loopback 127.0.0.01
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