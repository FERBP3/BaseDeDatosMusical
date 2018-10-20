package main

import (
	"github.com/gotk3/gotk3/gtk"
)

func NuevaVentanaEditaRola(view *View) {
	campos, entradas := camposRola(view)
	dialog,dialogCaja := creaDialog("Editar Rola")
	botonGuardarCambios := NuevoButonGuardaCambiosRola("Guardar rola", dialog, entradas, view)
	botonNoGuardarCambios := NuevoBotonNoGuardaCambios("No guardar rola", dialog)
	header := creaHeaderBar()
	header.PackStart(botonGuardarCambios)
	header.PackEnd(botonNoGuardarCambios)
	dialogCaja.Add(header)
	dialogCaja.Add(campos)
	dialog.SetModal(true)
	dialog.ShowAll()
	//dialog.Run()
}

func NuevaVentanaEditaAlbum(view *View) {
	dialog, dialogCaja := creaDialog("Editar Álbum")
	campos, entradas := camposAlbum(view)
	botonGuardarCambios := NuevoButonGuardaCambiosAlbum("Guardar Album", dialog, entradas, view)
	botonNoGuardarCambios := NuevoBotonNoGuardaCambios("No guardar Album", dialog)
	header := creaHeaderBar()
	header.PackStart(botonGuardarCambios)
	header.PackEnd(botonNoGuardarCambios)
	dialogCaja.Add(header)
	dialogCaja.Add(campos)
	dialog.SetModal(true)
	dialog.ShowAll()
	//dialog.Run()
}

func NuevaVentanaEditaInterprete(view *View, tipo string) {
	dialog, dialogCaja := creaDialog("Edita Interprete") 
	botonNoGuardarCambios := NuevoBotonNoGuardaCambios("Don't Save", dialog)
	notebook,_ := gtk.NotebookNew()

	if tipo == "2" {

	camposPersona, entradasPersona := camposPersona(view, true)
	camposGrupo, entradasGrupo := camposGrupo(view, true)

	botonGuardarPersona := NuevoButonGuardaPersona("Save Person",dialog, dialog, entradasPersona, view)

	headerPersona := creaHeaderBar()
	headerPersona.PackStart(botonGuardarPersona)
	headerPersona.PackStart(botonNoGuardarCambios)

	botonGuardarGrupo := NuevoButonGuardaGrupo("Save as Group", dialog, view, entradasGrupo)

	headerGrupo := creaHeaderBar()
	headerGrupo.PackStart(botonGuardarGrupo)

	camposPersona.Prepend(headerPersona)
	camposGrupo.Prepend(headerGrupo)

	tabPersona := creaEtiqueta("Person")
	tabGrupo := creaEtiqueta("Group")

	notebook.AppendPage(camposPersona,tabPersona)
	notebook.AppendPage(camposGrupo,tabGrupo)

	dialogCaja.Add(notebook)

	} else if tipo == "0" {
	campos, entradas := camposPersona(view, false)

	botonGuardarPersona := NuevoButonGuardaPersona("Save Person",dialog, dialog, entradas, view)

	headerPersona := creaHeaderBar()
	headerPersona.PackStart(botonGuardarPersona)
	headerPersona.PackStart(botonNoGuardarCambios)
	campos.Prepend(headerPersona)

	tabPersona := creaEtiqueta("Person")

	notebook.AppendPage(campos,tabPersona)

	dialogCaja.Add(notebook)

	} else {
	camposGrupo, entradasGrupo := camposGrupo(view, false)
	botonGuardarGrupo := NuevoButonGuardaGrupo("Save as Group", dialog, view, entradasGrupo)
	botonAgregaPersona := NuevoButonAgregaPersona("Add Person", dialog, view, true)

	headerGrupo := creaHeaderBar()
	headerGrupo.PackStart(botonGuardarGrupo)
	headerGrupo.PackStart(botonAgregaPersona)
	headerGrupo.PackStart(botonNoGuardarCambios)

	camposGrupo.Prepend(headerGrupo)
	tabGrupo := creaEtiqueta("Group")
	notebook.AppendPage(camposGrupo,tabGrupo)
	dialogCaja.Add(notebook)
	}

	dialog.SetModal(true)
	dialog.ShowAll()
	//dialog.Run()
}

func NuevaVentanaAgregaPersona(dialogPadre *gtk.Dialog, view *View, esNueva bool) {
	dialog,dialogCaja := creaDialog("Add Person")
	botonNoGuardarCambios := NuevoBotonNoGuardaCambios("Don't add", dialog)
	header := creaHeaderBar()

	tabAddPerson := creaEtiqueta("New person")
	tabExistingPerson := creaEtiqueta("Existing person")

	boxExistingPerson := boxExistingPerson(view, true)
	campos, entradas := camposPersona(view, esNueva)

	botonGuardarCambios := NuevoButonGuardaPersona("Add", dialogPadre, dialog, entradas, view)
	header.PackStart(botonGuardarCambios)
	header.PackEnd(botonNoGuardarCambios)
	dialogCaja.Add(header)

	notebook,_ := gtk.NotebookNew()
	notebook.AppendPage(campos, tabAddPerson)
	notebook.AppendPage(boxExistingPerson, tabExistingPerson)

	dialogCaja.Add(notebook)
	dialog.ShowAll()
	//dialog.Run()
}


