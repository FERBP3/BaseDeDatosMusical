package dao

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "github.com/FERBP3/BaseDeDatosMusical/Util"
    "os/user"
    "log"
)

func get() (*sql.DB) {
    usr, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }
    ruta  := usr.HomeDir+"/.local/rolas.db"
    //os.Remove(ruta)
    database, err := sql.Open("sqlite3", ruta)
    _, err = database.Exec(Util.Esquema)
    return database
}	
