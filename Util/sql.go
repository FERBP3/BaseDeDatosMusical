package Util

const (
Esquema = `
CREATE TABLE types (
id_type       INTEGER PRIMARY KEY,
description   TEXT
);

INSERT INTO types VALUES(0,'Person');
INSERT INTO types VALUES(1,'Group');
INSERT INTO types VALUES(2,'Unknown');


CREATE TABLE performers (
    id_performer  INTEGER PRIMARY KEY,
    id_type       INTEGER,
    name          TEXT,
    FOREIGN KEY   (id_type) REFERENCES types(id_type)
);

CREATE TABLE persons (
    id_person     INTEGER PRIMARY KEY,
    stage_name    TEXT,
    real_name     TEXT,
    birth_date    TEXT,
    death_date    TEXT
);

CREATE TABLE groups (
    id_group      INTEGER PRIMARY KEY,
    name          TEXT,
    start_date    TEXT,
    end_date      TEXT
);

CREATE TABLE albums (
    id_album      INTEGER PRIMARY KEY,
    path          TEXT,
    name          TEXT,
    year          INTEGER
);

CREATE TABLE rolas (
    id_rola       INTEGER PRIMARY KEY,
    id_performer  INTEGER,
    id_album      INTEGER,
    path          TEXT,
    title         TEXT,
    track         INTEGER,
    year          INTEGER,
    genre         TEXT,
    FOREIGN KEY   (id_performer) REFERENCES performers(id_performer),
    FOREIGN KEY   (id_album) REFERENCES albums(id_album)
);

CREATE TABLE in_group (
    id_person     INTEGER,
    id_group      INTEGER,
    PRIMARY KEY   (id_person, id_group),
    FOREIGN KEY   (id_person) REFERENCES persons(id_person),
    FOREIGN KEY   (id_group) REFERENCES  groups(id_group)
);
`
SeleccionaPerformerName = "SELECT name FROM performers WHERE name=$1"
SeleccionaPerformer = "SELECT id_performer, name FROM performers WHERE name=$1"
SeleccionaAlbumName = "SELECT name FROM albums WHERE name=$1"
SeleccionaAlbum = "SELECT id_album, name FROM albums WHERE name=$1"
SeleccionaRola = "SELECT path, title FROM rolas WHERE path=$1 AND title=$2"
SeleccionaPersona = `SELECT stage_name, real_name, birth_date, death_date 
                     FROM persons 
                     WHERE stage_name=$1 AND real_name = $2 AND birth_date = $3 AND death_date = $4`

SeleccionaPersonaFromStageName = `SELECT stage_name, real_name, birth_date, death_date 
                                  FROM persons
                                  WHERE stage_name = $1`

SeleccionaGrupoFromName = `SELECT name, start_date, end_date 
                           FROM groups
                           WHERE name = $1`

SeleccionaVistaRola = `SELECT title, performers.name, albums.name, genre, track, rolas.year, rolas.path
                       FROM rolas 
                       INNER JOIN albums ON rolas.id_album = albums.id_album
                       INNER JOIN performers ON rolas.id_performer = performers.id_performer`

SeleccionaTipoPerformer = `SELECT id_type
                           FROM rolas INNER JOIN performers 
                           ON rolas.id_performer = performers.id_performer AND path = $1`

SeleccionaGrupo = `SELECT name, start_date, end_date
                   FROM groups
                   WHERE name = $1 AND start_date = $2 AND end_date = $3`

InsertaPerformer = "INSERT INTO performers (id_type,name) VALUES($1,$2)"
InsertaAlbum = "INSERT INTO albums (path,name,year) VALUES($1,$2,$3)"
InsertaRola = "INSERT INTO rolas (id_performer, id_album, path, title, track, year, genre) VALUES($1,$2,$3,$4,$5,$6,$7)"
InsertaPersona = `INSERT INTO persons (stage_name, real_name, birth_date, death_date) VALUES($1,$2,$3,$4)`
InsertaGrupo = `INSERT INTO groups (name, start_date, end_date) VALUES($1,$2,$3)`

ActualizaRola = `UPDATE rolas
                 SET title = $1,
                 track = $2,
                 year = $3,
                 genre = $4
                WHERE path = $5`

ActualizaAlbum = `UPDATE albums
                 SET name = $1,
                 year = $2
                 WHERE path = $3`

ActualizaInterprete = `UPDATE performers
                        SET id_type = $1
                        WHERE name = $2`
)
