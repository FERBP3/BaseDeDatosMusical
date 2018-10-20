package main

import(
  "github.com/FERBP3/BaseDeDatosMusical/modelos"
  "github.com/FERBP3/BaseDeDatosMusical/dao"
  "github.com/bogem/id3v2"
  "path/filepath"
  "os/user"
  "os"
	"fmt"
	"log"
)

const MusicDir = "/Music"
//const MusicDir = "/Música/Test2"

func main(){

  rutas, err := Recorre("")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%v rutas por leer\n", len(rutas))

  rolas := CreaRolas(rutas)
  fmt.Println(len(rolas)," mp3 leídos correctamente")

   var id_performer int64
   var id_album int64

   for _,rola := range rolas {
    //SE AÑADEN LOS DATOS A LA TABLA PERFORMERS
       id_performer, err = dao.InsertaPerformer(rola)
       if err != nil {
           log.Fatal(err)
       }
    //SE AÑADEN LOS DATOS A LA TABLA ALBUM
       id_album, err = dao.InsertaAlbum(rola)
       if err != nil {
           log.Fatal(err)
       }
    //SE AÑADEN LOS DATOS A LA TABLA ROLAS
       err = dao.InsertaRola(rola,id_performer, id_album)
       if err != nil {
           log.Fatal(err)
       }
   }

}

func Recorre(directorio string) ([]string, error) {
  if directorio == ""{
    directorio, _ = GetDirMusic()
  }
  var rutas []string
  err := filepath.Walk(directorio, Agrega(&rutas))
  return rutas,err  
}

func GetDirMusic() (string, error){
  usr, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }
    rutaMusic := usr.HomeDir+MusicDir
    //fmt.Println(rutaMusic)
    return rutaMusic,err
}

func Agrega(rutas *[]string) filepath.WalkFunc {
  return func (path string, info os.FileInfo, err error) error {
    if err != nil {
      log.Fatal(err)
    }
    if filepath.Ext(path) != ".mp3"{
      return nil
    }
    *rutas = append(*rutas,path)
    return nil
  }
}

func CreaRolas(rutas []string) ([]*modelos.Rola){
    var rolas []*modelos.Rola
    var rola *modelos.Rola
    var tag *id3v2.Tag
    var err error
    for _, ruta := range rutas {
    tag, err = id3v2.Open(ruta, id3v2.Options{Parse:true})
    if err != nil{
        fmt.Println("Error al leer archivo mp3 "+ruta, err)
        continue
    }
    rola = &modelos.Rola{
            Interprete: GetTagInterprete(tag),
            Titulo: GetTagTitulo(tag),
            Album: GetTagAlbum(tag),
            FechaGrabacion: GetTagFechaGrabacion(tag),
            Genero: GetTagGenero(tag),
            Track: GetTagTrack(tag),
            Path: ruta,
    }
    rolas = append(rolas, rola)
    }

    tag.Close()
    return rolas    
}
