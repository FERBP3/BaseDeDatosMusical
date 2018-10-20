package main

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/FERBP3/BaseDeDatosMusical/modelos"
    "github.com/FERBP3/BaseDeDatosMusical/dao"
)

func camposPersona(view *View, esNueva bool) (*gtk.ListBox, *EntradaPersona){
	campos  := creaListBox()
	renglon := creaRenglonListBox()
	
	var persona *modelos.Person
	if esNueva {
		persona = &modelos.Person{}
	} else {
		persona = getPersonfromView(view)
	}

	EntradaPersona := &EntradaPersona{}
	caja := creaCaja()
	etiqueta := creaEtiqueta("Stage name: ")
	entrada := creaEntrada(persona.StageName)
	caja.PackStart(etiqueta,false,false,0)
	caja.PackStart(entrada,true,true,0)
	renglon.Add(caja)
	campos.Add(renglon)
	EntradaPersona.StageName = entrada

	renglon = creaRenglonListBox()
	caja = creaCaja()
	etiqueta = creaEtiqueta("Real name: ")
	entrada = creaEntrada(persona.RealName)
	caja.PackStart(etiqueta,false,false,0)
	caja.PackStart(entrada,true,true,0)
	renglon.Add(caja)
	campos.Add(renglon)
	EntradaPersona.RealName = entrada

	renglon = creaRenglonListBox()
	caja = creaCaja()
	etiqueta = creaEtiqueta("Date of birth: ")
	entrada = creaEntrada(persona.DateBirth)
	caja.PackStart(etiqueta,false,false,0)
	caja.PackStart(entrada,true,true,0)
	renglon.Add(caja)
	campos.Add(renglon)
	EntradaPersona.DateBirth = entrada

	renglon = creaRenglonListBox()
	caja = creaCaja()
	etiqueta = creaEtiqueta("Date of death: ")
	entrada = creaEntrada(persona.DateDeath)
	caja.PackStart(etiqueta,false,false,0)
	caja.PackStart(entrada,true,true,0)
	renglon.Add(caja)
	campos.Add(renglon)
	EntradaPersona.DateDeath = entrada

	return campos, EntradaPersona
}

func camposGrupo(view *View, esNueva bool) (*gtk.ListBox, *EntradaGrupo){
	campos  := creaListBox()

	var grupo *modelos.Group
	if esNueva {
		grupo = &modelos.Group{}
	} else {
		grupo = getGroupFromView(view)
	}

	EntradaGrupo := &EntradaGrupo{}
	renglon := creaRenglonListBox()
	caja := creaCaja()
	etiqueta := creaEtiqueta("Name: ")
	entrada := creaEntrada(grupo.Name)
	caja.PackStart(etiqueta,false,false,0)
	caja.PackStart(entrada,true,true,0)
	renglon.Add(caja)
	campos.Add(renglon)
	EntradaGrupo.Name = entrada

	renglon = creaRenglonListBox()
	caja = creaCaja()
	etiqueta = creaEtiqueta("Start date: ")
	entrada = creaEntrada(grupo.StartDate)
	caja.PackStart(etiqueta,false,false,0)
	caja.PackStart(entrada,true,true,0)
	renglon.Add(caja)
	campos.Add(renglon)
	EntradaGrupo.StartDate = entrada

	renglon = creaRenglonListBox()
	caja = creaCaja()
	etiqueta = creaEtiqueta("End date: ")
	entrada = creaEntrada(grupo.EndDate)
	caja.PackStart(etiqueta,false,false,0)
	caja.PackStart(entrada,true,true,0)
	renglon.Add(caja)
	campos.Add(renglon)
	EntradaGrupo.EndDate = entrada

	return campos, EntradaGrupo
}

func camposRola(view *View) (*gtk.ListBox,*EntradaRola) {
	rola := getRolaFromView(view)

	entradaRola := &EntradaRola{}
	campos  := creaListBox()
	renglon := creaRenglonListBox()
	caja := creaCaja()
	etiqueta := creaEtiqueta("Título:")
	entrada := creaEntrada(rola.Titulo)
	caja.PackStart(etiqueta,false,false,0)
	caja.PackStart(entrada,true,true,0)
	renglon.Add(caja)
	campos.Add(renglon)
	entradaRola.Title = entrada

	renglon = creaRenglonListBox()
	caja = creaCaja()
	etiqueta = creaEtiqueta("Track:")
	entrada = creaEntrada(rola.Track)
	caja.PackStart(etiqueta,false,false,0)
	caja.PackStart(entrada,true,true,0)
	renglon.Add(caja)
	campos.Add(renglon)
	entradaRola.Track = entrada

	renglon = creaRenglonListBox()
	caja = creaCaja()
	etiqueta = creaEtiqueta("Año:")
	entrada = creaEntrada(rola.FechaGrabacion)
	caja.PackStart(etiqueta,false,false,0)
	caja.PackStart(entrada,true,true,0)
	renglon.Add(caja)
	campos.Add(renglon)
	entradaRola.Year = entrada

	renglon = creaRenglonListBox()
	caja = creaCaja()
	etiqueta = creaEtiqueta("Género:")
	entrada = creaEntrada(rola.Genero)
	caja.PackStart(etiqueta,false,false,0)
	caja.PackStart(entrada,true,true,0)
	renglon.Add(caja)
	campos.Add(renglon)
	entradaRola.Genre = entrada

	return campos, entradaRola
}

func camposAlbum(view *View) (*gtk.ListBox, *EntradaAlbum) {
	rola := getRolaFromView(view)
	entradaAlbum := &EntradaAlbum{}

	campos  := creaListBox()
	renglon := creaRenglonListBox()
	caja := creaCaja()
	etiqueta := creaEtiqueta("Nombre:")
	entrada := creaEntrada(rola.Album)
	caja.PackStart(etiqueta,false,false,0)
	caja.PackStart(entrada,true,true,0)
	renglon.Add(caja)
	campos.Add(renglon)
	entradaAlbum.Name = entrada

	renglon = creaRenglonListBox()
	caja = creaCaja()
	etiqueta = creaEtiqueta("Año:")
	entrada = creaEntrada(rola.FechaGrabacion)
	caja.PackStart(etiqueta,false,false,0)
	caja.PackStart(entrada,true,true,0)
	renglon.Add(caja)
	campos.Add(renglon)
	entradaAlbum.Year = entrada

	return campos, entradaAlbum
}

func getRolaFromView(view *View) (*modelos.Rola) {
	s, _ := view.TreeView.GetSelection()
	_,iter,_ := s.GetSelected()
	path, _ := view.ListStore.GetPath(iter)
	indice := path.GetIndices()[0]
	return view.Rolas[indice]
}

func getPersonfromView(view *View) (*modelos.Person) {
	rola := getRolaFromView(view)
	person := dao.GetPerson(rola)
	return person
}

func getGroupFromView(view *View) (*modelos.Group) {
	rola := getRolaFromView(view)
	group := dao.GetGroup(rola)
	return group
}

func boxExistingPerson(view *View, esNueva bool) (*gtk.Box){
	box := creaCaja()
	combo,_ := gtk.ComboBoxTextNew()

	if !esNueva {
		group := getGroupFromView(view)
		for _, member := range group.Members {
		combo.AppendText(member.StageName)
		}
	}

	box.PackStart(combo, false, true, 0)
	return box
}



