package main

import(
  "github.com/FERBP3/ReproductorMP3/modelos"
  "github.com/FERBP3/ReproductorMP3/dao"
  "github.com/bogem/id3v2"
  rbt "github.com/emirpasic/gods/trees/redblacktree"
  "path/filepath"
  "os/user"
  "os"
	"fmt"
	"log"
)

const MusicDir = "/Música/Test2"

func main(){

  arbol, err := Recorre("")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%v rutas por leer\n",arbol.Size())

  rolas := CreaRolas(arbol)
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

    //rolas = dao.GetRolas()
    //for _,rola := range rolas {
        //fmt.Println(rola.Titulo, rola.Interprete, rola.Album, rola.Genero)
    //}
}

func Recorre(directorio string) (*rbt.Tree, error) {
  if directorio == ""{
    directorio, _ = GetDirMusic()
  }
  arbol := rbt.NewWithStringComparator()
  err := filepath.Walk(directorio, Agrega(arbol))
  return arbol,err  
}

func GetDirMusic() (string, error){
  usr, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }
    rutaMusic := usr.HomeDir+MusicDir
    fmt.Println(rutaMusic)/////////-------------------------////////
    return rutaMusic,err
}

func Agrega(arbol *rbt.Tree) filepath.WalkFunc {
  return func (path string, info os.FileInfo, err error) error {
    if err != nil {
      log.Fatal(err)
    }
    if filepath.Ext(path) != ".mp3"{
      return nil
    }
    arbol.Put(path,0)
    return nil
  }
}

func CreaRolas(arbol *rbt.Tree) ([]*modelos.Rola){
    var rolas []*modelos.Rola
    var rola *modelos.Rola
    var ruta string
    var tag *id3v2.Tag
    var err error
    iterador := arbol.Iterator()
    for iterador.Next() {  
    ruta = iterador.Key().(string)  
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
