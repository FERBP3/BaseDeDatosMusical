# Base de Datos Musical

Sistema para crear y manipular una base datos musical

## Instalación

Se debe tener instalado Go.

## Dependencias

Se usaron las siguientes dependencias:
- gotk3:
https://github.com/gotk3/gotk3
- el driver para sqlite3:
https://github.com/mattn/go-sqlite3
- Biblioteca para leer las etiquetas de los MP3:
https://github.com/bogem/id3v2

Cada una se puede instalar con el comando go get:
```
$ go get github.com/gotk3/gotk3
$ go get github.com/mattn/go-sqlite3
$ go get github.com/bogem/id3v2 
```
## Intalación
```
$ go get github.com/FERBP3/BaseDeDatosMusical
```
### Ejecución
Estando en el $GOPATH ejecutamos lo siguiente para minar la base de datos primero:
```
$ go install github/FERBP3/BaseDeDatosMusical/minero
```
De esta forma se obtiene el binario en la carpeta src/ del $GOPATH y luego lo ejecutamos (estando situado en $GOPATH):
```
$./bin/minero
```
Después se genera el binario del programa:
```
$ go install github/FERBP3/BaseDeDatosMusical/vista
```
Después lo ejecutamos (de igual forma estando en $GOPATH):
```
$./bin/vista
```

## Autor
* Brigido Pablo José Fernando
* 314146171