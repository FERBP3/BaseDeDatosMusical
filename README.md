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
### Consultas
Las consultas funcionan de la siguiente manera:
Para buscar por título se escribe:
```
T:Black Diamond
```
seguido de la cadena que se quiere buscar, en ese caso de buscó el título Black Diamond.
Los mismo para buscar por Álbum o por Interprete:
```
A:Imaginaerum
I:Nightwish
```
Sólo puede buscar una categoría cada vez.
Cualquier cadena que no reconozca automáticamente mostrará todas las rolas.

### DETALLES
La carpeta donde se leerán los archivos MP3 por omisión es 
```
~/Music.  
```
Al definir un interprete como persona se debe poner en el campo StageName el nombre del campo artista de la rola que está seleccionada en la vista. Lo mismo para definir como grupo, se debe poner en el campo Name el nombre del campo Artista de la rola que está seleccionada.

## Autor
* Brigido Pablo José Fernando
* 314146171