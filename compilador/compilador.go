package compilador 

import (
	"strings"
)

var SeleccionaVistaRola = `SELECT title, performers.name, albums.name, genre, track, rolas.year, rolas.path
                       FROM rolas 
                       INNER JOIN albums ON rolas.id_album = albums.id_album
                       INNER JOIN performers ON rolas.id_performer = performers.id_performer`


func Compila(entrada string) (string) {
	entrada = strings.TrimSpace(entrada)
	var salida string

	if len(entrada) < 3 {
		return ""
	}
	if entrada[1] != ':' {
		return ""
	}
	palabra := strings.TrimSpace(entrada[2:])

	switch entrada[0] {
	case 'T':
		salida = SeleccionaVistaRola+" WHERE title LIKE '%"+palabra+"%'"
	case 'I':
		salida = SeleccionaVistaRola+" WHERE performers.name LIKE '%"+palabra+"%'"
	case 'A':
		salida = SeleccionaVistaRola+" WHERE albums.name LIKE '%"+palabra+"%'"
	default:
		salida = ""
	}
	return salida
}