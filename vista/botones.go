package main

import(
	"log"
	"github.com/gotk3/gotk3/gtk"
  "github.com/FERBP3/BaseDeDatosMusical/dao"
 	"github.com/FERBP3/BaseDeDatosMusical/modelos"
	//"fmt"
)

func NuevoBotonNoGuardaCambios(etiqueta string, dialog *gtk.Dialog)  (*gtk.Button){
	boton, err := gtk.ButtonNewWithLabel(etiqueta)
	if err != nil {
		log.Fatal("No se pudo crear el boton Guardar cambios : ", err)
	}
	boton.Connect("clicked", func() {
		//fmt.Println("NO se guardaron los cambios de: ",etiqueta)
		dialog.Destroy()
		})
	return boton
}

func NuevoButonAgregaPersona(etiqueta string, dialog *gtk.Dialog, view *View, esNueva bool) (*gtk.Button) {
	boton, err := gtk.ButtonNewWithLabel(etiqueta)
	if err != nil {
		log.Fatal("No se pudo crear el boton Guardar cambios : ", err)
	}
	boton.Connect("clicked", func() {
		NuevaVentanaAgregaPersona(dialog, view, esNueva)
		})
	return boton

}

func NuevoButonGuardaPersona(etiqueta string, dialogPadre  *gtk.Dialog, dialogActual *gtk.Dialog, entradas *EntradaPersona, view *View) (*gtk.Button) {
	boton, err := gtk.ButtonNewWithLabel(etiqueta)
	if err != nil {
		log.Fatal("No se pudo crear el boton Guardar cambios : ", err)
	}
	boton.Connect("clicked", func() {
		//////-------GUARDA PERSONA A LA BASE DE DATOS---------//////
		stageName, _ := entradas.StageName.GetText()
		realName, _ := entradas.RealName.GetText()
		dateBirth, _ := entradas.DateBirth.GetText()
		dateDeath, _ := entradas.DateDeath.GetText()

		persona := &modelos.Person{
			StageName: stageName,
			RealName: realName,
			DateBirth: dateBirth,
			DateDeath: dateDeath,
		}

		dao.InsertaPersona(persona, getRolaFromView(view))
		dialogPadre.Destroy()
		dialogActual.Destroy()
		})
	return boton
}

func NuevoButonGuardaGrupo(etiqueta string, dialog *gtk.Dialog, view *View, entradas *EntradaGrupo) (*gtk.Button) {
	boton, err := gtk.ButtonNewWithLabel(etiqueta)
	if err != nil {
		log.Fatal("No se pudo crear el boton Guardar cambios : ", err)
	}
	boton.Connect("clicked", func() {
		name, _ := entradas.Name.GetText()
		start, _ := entradas.StartDate.GetText()
		end, _ := entradas.EndDate.GetText()

		grupo := &modelos.Group{
			Name: name, 
			StartDate: start,
			EndDate: end,
		}

		dao.InsertaGrupo(grupo, getRolaFromView(view))
		dialog.Destroy()
		})
	return boton
}

func NuevoButonGuardaCambios(etiqueta string, dialog *gtk.Dialog) (*gtk.Button) {
	boton, err := gtk.ButtonNewWithLabel(etiqueta)
	if err != nil {
		log.Fatal("No se pudo crear el boton Guardar cambios : ", err)
	}
	boton.Connect("clicked", func() {
		//fmt.Println("Se guardaron los cambios de: ",etiqueta)
		dialog.Destroy()
		})
	return boton
}

func NuevoButonGuardaCambiosAlbum(etiqueta string, dialog *gtk.Dialog, entradas *EntradaAlbum, view *View) (*gtk.Button) {
	boton, err := gtk.ButtonNewWithLabel(etiqueta)
	if err != nil {
		log.Fatal("No se pudo crear el boton Guardar cambios : ", err)
	}
	boton.Connect("clicked", func() {
		s, _ := view.TreeView.GetSelection()
		_,iter,_ := s.GetSelected()

		name, _ := entradas.Name.GetText()
		year, _ := entradas.Year.GetText()

		path, _ := view.ListStore.GetPath(iter)
		indice := path.GetIndices()[0]

		view.Rolas[indice].Album = name
		view.Rolas[indice].FechaGrabacion = year

		dao.SaveAlbum(view.Rolas[indice])//EL DAO GUARDA LOS CAMBIOS EN LA BASE DE DATOS

		err := view.ListStore.SetValue(iter, COLUMNA_ALBUM, name)
		if err != nil {
			//fmt.Println(err)
		}
		dialog.Destroy()
		})
	return boton
}

func NuevoButonGuardaCambiosRola(etiqueta string, dialog *gtk.Dialog, entradas *EntradaRola, view *View) (*gtk.Button) {
	boton, err := gtk.ButtonNewWithLabel(etiqueta)
	if err != nil {
		log.Fatal("No se pudo crear el boton Guardar cambios : ", err)
	}
	boton.Connect("clicked", func() {
		s, _ := view.TreeView.GetSelection()
		_,iter,_ := s.GetSelected()

		path, _ := view.ListStore.GetPath(iter)
		indice := path.GetIndices()[0]///indice

		title, _ := entradas.Title.GetText()
		genre, _ := entradas.Genre.GetText()
		track, _ := entradas.Track.GetText()
		year, _ := entradas.Year.GetText()

		view.Rolas[indice].Titulo = title
		view.Rolas[indice].FechaGrabacion = year
		view.Rolas[indice].Genero = genre
		view.Rolas[indice].Track = track
		
		dao.SaveRola(view.Rolas[indice])//EL DAO GUARDA LOS CAMBIOS EN LA BASE DE DATOS

		err := view.ListStore.SetValue(iter, COLUMNA_TITULO, title)
		err = view.ListStore.SetValue(iter, COLUMNA_GENERO, genre)

		if err != nil {
			//fmt.Println(err)
		}
		dialog.Destroy()
		})
	return boton
}
