package main

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/FERBP3/BaseDeDatosMusical/modelos"
)

type View struct {
	TreeView *gtk.TreeView
	ListStore *gtk.ListStore
	Rolas []*modelos.Rola
}

type EntradaRola struct{
	Title *gtk.Entry
	Track *gtk.Entry
	Year *gtk.Entry
	Genre *gtk.Entry
}

type EntradaAlbum struct{
	Name *gtk.Entry
	Year *gtk.Entry
}

type EntradaPersona struct {
	StageName *gtk.Entry
	RealName *gtk.Entry
	DateBirth *gtk.Entry
	DateDeath *gtk.Entry
}

type EntradaGrupo struct {
	Name *gtk.Entry
	StartDate *gtk.Entry
	EndDate *gtk.Entry
}
