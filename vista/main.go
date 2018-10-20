package main

import (
	"log"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/glib"
	"os"
    "github.com/FERBP3/BaseDeDatosMusical/dao"
    "fmt"
)

func main() {
	const idApp = "Modelado.Proyecto.ReproductorMP3"
	aplicacion, err := gtk.ApplicationNew(idApp,glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatal("No se pudo crear la aplicación: ",err)
	}
	aplicacion.Connect("activate", onActivate)
	os.Exit(aplicacion.Run(os.Args))
}

func onActivate(aplicacion *gtk.Application) {
	//minero.Mina()
	ventanaApp, err := gtk.ApplicationWindowNew(aplicacion)
	if err != nil {
		log.Fatal("No se pudo crear la ventana de la aplicacion: ",err)
	}
	ventanaApp.SetTitle("ventana")
	ventanaApp.SetDefaultSize(1200, 800)
	ventanaApp.SetPosition(gtk.WIN_POS_CENTER)

	header := creaHeaderBar()
	TreeView, ListStore := creaTreeView()
	ventanaScroll := creaVentanaScroll()
	botonMenu := creaBotonMenu()
	barraBusqueda := creaBarraBusqueda()

	rolas := dao.GetRolas()
	for _,rola := range rolas {
		agregaRenglon(ListStore,rola.Titulo,rola.Interprete, rola.Album, rola.Genero)
	}

	view := &View{TreeView: TreeView, ListStore: ListStore, Rolas: rolas}

	editaRola := glib.SimpleActionNew("editaRola", nil)
	editaRola.Connect("activate", func() {
		NuevaVentanaEditaRola(view)
	})
	aplicacion.AddAction(editaRola)

	editaAlbum := glib.SimpleActionNew("editaAlbum", nil)
	editaAlbum.Connect("activate", func() {
		NuevaVentanaEditaAlbum(view)
	})
	aplicacion.AddAction(editaAlbum)

	editaInterprete := glib.SimpleActionNew("editaInterprete", nil)
	editaInterprete.Connect("activate", func() {
		tipo := dao.GetTypePerformer(getRolaFromView(view))
		fmt.Println(getRolaFromView(view))
		fmt.Println(tipo)
		NuevaVentanaEditaInterprete(view, tipo )
	})
	aplicacion.AddAction(editaInterprete)

	header.SetTitle("Reproductor")
	header.SetSubtitle("Music database")
	header.SetShowCloseButton(true)

	header.PackStart(barraBusqueda)
	header.PackStart(botonMenu)

	ventanaApp.SetTitlebar(header)

	ventanaScroll.Add(TreeView)
	ventanaApp.Add(ventanaScroll)

	ventanaApp.ShowAll()
}
