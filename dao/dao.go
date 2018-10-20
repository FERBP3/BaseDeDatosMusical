package dao

import(
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
	"github.com/FERBP3/BaseDeDatosMusical/modelos"
    "github.com/FERBP3/BaseDeDatosMusical/Util"
    //"fmt"
)

func InsertaPerformer(rola *modelos.Rola) (int64,error) {
	database := get()
	defer database.Close()

    var id_performer int64
    var name string
    row := database.QueryRow(Util.SeleccionaPerformerName, rola.Interprete)
    err := row.Scan(&name)
    if err == sql.ErrNoRows{
        result, err := database.Exec(Util.InsertaPerformer,2,rola.Interprete)
        if err != nil {
            return id_performer, err
        }
        id_performer, errId := result.LastInsertId()
        if errId != nil {
            return id_performer, err
        }
        //fmt.Printf("Se agregó el intérprete %+v a la base de datos con %v como ID\n", rola.Interprete, id_performer)
        return id_performer,nil
    } else if err == nil {
        row = database.QueryRow(Util.SeleccionaPerformer, rola.Interprete)
        row.Scan(&id_performer,&name)
        return id_performer, nil
    }else{
        return id_performer, err
    }
}

func InsertaAlbum(rola *modelos.Rola) (int64, error) {
	database := get()
	defer database.Close()

    var id_album int64
    var name string
    row := database.QueryRow(Util.SeleccionaAlbumName, rola.Album)
    err := row.Scan(&name)
    if err == sql.ErrNoRows {
        result, err := database.Exec(Util.InsertaAlbum,rola.Path,rola.Album,rola.FechaGrabacion)
        if err != nil {
            return id_album, err
        }
        id_album, err = result.LastInsertId() 
        if err != nil {
            return id_album, err
        }
        //fmt.Printf("Se agregó el album %+v a la base de datos con %v como ID\n", rola.Album, id_album)
        return id_album,nil
    } else if err == nil {
        row = database.QueryRow(Util.SeleccionaAlbum, rola.Album)
        row.Scan(&id_album,&name)
        return id_album,nil
    } else {
        return id_album,err
    }
}

func InsertaRola(rola *modelos.Rola, id_performer int64, id_album int64) (error) {
	database := get()
	defer database.Close()

    var path string
    var titulo string
    row := database.QueryRow(Util.SeleccionaRola, rola.Path, rola.Album)
    err := row.Scan(&path, &titulo)
    if err == sql.ErrNoRows {
        _, err = database.Exec(Util.InsertaRola,id_performer, id_album, rola.Path, rola.Titulo, rola.Track, rola.FechaGrabacion, rola.Genero) 
        if err != nil {
            return err
        }
            //fmt.Printf("Se agregó la rola %+v a la base de datos\n",rola.Titulo)
    }
    return nil
}

func InsertaPersona(persona *modelos.Person, rola *modelos.Rola) {
	database := get()
	defer database.Close()

    var stageName string
    var realName string
    var DateBirth string
    var DateDeath string
    row := database.QueryRow(Util.SeleccionaPersona, persona.StageName, persona.RealName, persona.DateBirth, persona.DateDeath)
    err := row.Scan(&stageName, &realName, &DateBirth, &DateDeath)
    if err == sql.ErrNoRows {
        _, err = database.Exec(Util.InsertaPersona, persona.StageName, persona.RealName, persona.DateBirth, persona.DateDeath) 
        if err != nil {
           // fmt.Println(err)
        }
    }

    _, err = database.Exec(Util.ActualizaInterprete, 0, rola.Interprete)
    if err != nil {
    	//fmt.Println(err)
    }
}

func InsertaGrupo(grupo *modelos.Group, rola *modelos.Rola) {
	database := get()
	defer database.Close()

    var name string
    var start string
    var end string
    row := database.QueryRow(Util.SeleccionaGrupo, grupo.Name, grupo.StartDate, grupo.EndDate)
    err := row.Scan(&name, &start, &end)
    if err == sql.ErrNoRows {
        _, err = database.Exec(Util.InsertaGrupo, grupo.Name, grupo.StartDate, grupo.EndDate) 
        if err != nil {
            //fmt.Println(err)
        }
    }

    _, err = database.Exec(Util.ActualizaInterprete, 1, rola.Interprete)
    if err != nil {
    	//fmt.Println(err)
    }
}

func GetAllRolas() ([]*modelos.Rola) {
	database := get()
	defer database.Close()

	renglones, err := database.Query(Util.SeleccionaVistaRola)
	if err != nil {
		//fmt.Println("No se pudo hacer la consulta: ", err)
	}
	var rolas []*modelos.Rola
	var rola*modelos.Rola
	var title string 
	var artist string
	var album string
	var genre string
	var track string
	var fechaGrabacion string
	var path string
	for renglones.Next() {
		err = renglones.Scan(&title, &artist, &album, &genre, &track, &fechaGrabacion, &path)
		if err != nil {
        //fmt.Println("Error al leer el renglon: ", err)
        continue
		}
		rola = &modelos.Rola{
        	Titulo: title,
        	Interprete: artist,
        	Album: album,
        	Genero: genre,
        	Track: track,
        	FechaGrabacion: fechaGrabacion,
        	Path: path,
		}
		rolas = append(rolas, rola)
	}
	return rolas
}

func SaveRola(rola *modelos.Rola) {
	database := get()
	defer database.Close()

	_, err := database.Exec(Util.ActualizaRola, rola.Titulo, rola.Track, rola.FechaGrabacion, rola.Genero, rola.Path)
	if err != nil {
		//fmt.Println(err)
	}
}

func SaveAlbum(rola *modelos.Rola) {
	database := get()
	defer database.Close()
	_, err := database.Exec(Util.ActualizaAlbum, rola.Album, rola.FechaGrabacion, rola.Path)
	if err != nil {
		//fmt.Println(err)
	}
}

func GetTypePerformer(rola *modelos.Rola) (string) {
	database := get()
	defer database.Close()

	var tipo string
	renglon := database.QueryRow(Util.SeleccionaTipoPerformer, rola.Path)
	err := renglon.Scan(&tipo)
	if err != nil {
		//fmt.Println(err)
	}
	return tipo

}

func GetPerson(rola *modelos.Rola) (person *modelos.Person){
	database := get()
	defer database.Close()

	var stageName string
	var realName string
	var dateBirth string
	var dateDeath string

	renglon := database.QueryRow(Util.SeleccionaPersonaFromStageName, rola.Interprete)
	err := renglon.Scan(&stageName, &realName, &dateBirth, &dateDeath)
	if err != nil {
		//fmt.Println(err)
	}
	//fmt.Println(rola.Interprete)

	person = &modelos.Person{
		StageName: stageName,
		RealName: realName,
		DateBirth: dateBirth,
		DateDeath: dateDeath,
	}
	return person
}

func GetGroup(rola *modelos.Rola) (group *modelos.Group) {
	database := get()
	defer database.Close()

	var name string
	var start string
	var end string

	renglon := database.QueryRow(Util.SeleccionaGrupoFromName, rola.Interprete)
	err := renglon.Scan(&name, &start, &end)
	if err != nil {
		//fmt.Println(err)
	}

	grupo := &modelos.Group{
		Name: name,
		StartDate: start,
		EndDate: end,
	}
	return grupo
}

func Ejecuta(consulta string) ([]*modelos.Rola) {
    database := get()
    defer database.Close()

    renglones, err := database.Query(consulta)
    if err != nil {
        //fmt.Println("No se pudo hacer la consulta: ", err)
    }
    var rolas []*modelos.Rola
    var rola *modelos.Rola
    var title string 
    var artist string
    var album string
    var genre string
    var track string
    var fechaGrabacion string
    var path string
    for renglones.Next() {
        err = renglones.Scan(&title, &artist, &album, &genre, &track, &fechaGrabacion, &path)
        if err != nil {
        //fmt.Println("Error al leer el renglon: ", err)
        continue
        }
        rola = &modelos.Rola{
            Titulo: title,
            Interprete: artist,
            Album: album,
            Genero: genre,
            Track: track,
            FechaGrabacion: fechaGrabacion,
            Path: path,
        }
        rolas = append(rolas, rola)
    }
    return rolas
}
