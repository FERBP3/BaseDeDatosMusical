package main

import (
	"log"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/glib"
	"fmt"
)

const (
	COLUMNA_TITULO = iota
	COLUMNA_ARTISTA 
	COLUMNA_ALBUM
	COLUMNA_GENERO
)

func creaTreeView() (*gtk.TreeView, *gtk.ListStore) {
	treeView, err := gtk.TreeViewNew()
	if err != nil {
		log.Fatal("No se pudo crear el Tree View:",err)
	}
	treeView.AppendColumn(creaColumna("Título",COLUMNA_TITULO))
	treeView.AppendColumn(creaColumna("Artista",COLUMNA_ARTISTA))
	treeView.AppendColumn(creaColumna("Álbum",COLUMNA_ALBUM))
	treeView.AppendColumn(creaColumna("Género",COLUMNA_GENERO))

	listStore, err := gtk.ListStoreNew(glib.TYPE_STRING,glib.TYPE_STRING,glib.TYPE_STRING,glib.TYPE_STRING)
	if err != nil {
		log.Fatal("No se pudo crear la listStore: ", err)
	}
	treeView.SetModel(listStore)
	return treeView, listStore
}

func creaColumna(titulo string, id int) (*gtk.TreeViewColumn) {
	cellRenderer, err := gtk.CellRendererTextNew()
	if err != nil {
		log.Fatal("No se pudo crear el text cell renderer:", err)
	}
	columna, err := gtk.TreeViewColumnNewWithAttribute(titulo, cellRenderer, "text", id)
	if err != nil {
		log.Fatal("No se pudo crear la columna: ", err)
	}
	columna.SetExpand(true)
	return columna
}

func agregaRenglon(listStore *gtk.ListStore, titulo, artista, album, genero string){
	iterador := listStore.Append()
	err := listStore.Set(iterador,
		[]int{COLUMNA_TITULO,COLUMNA_ARTISTA,COLUMNA_ALBUM,COLUMNA_GENERO},
		[]interface{}{titulo,artista,album,genero})
	if err != nil {
		log.Fatal("No se pudo agregar el renglon: ",err)
	}
}

func creaVentanaScroll() (*gtk.ScrolledWindow){
	ventanaScroll, err := gtk.ScrolledWindowNew(nil,nil)
	if err != nil {
		log.Fatal("No se pudo crear la ventana Scroll: ", err)
	}
	return ventanaScroll
}

func creaBotonMenu() (*gtk.MenuButton) {
	botonMenu, err := gtk.MenuButtonNew()
	if err != nil {
		log.Fatal("No se pudo crear el boton menu: ", err)
	}
	menu := creaModeloMenu()
	botonMenu.SetMenuModel(&menu.MenuModel)
	return botonMenu
}

func creaModeloMenu() (*glib.Menu) {
	menu := glib.MenuNew()
	if menu == nil {
		log.Fatal("No se pudo crear el modelo del menu (nil)")
	}
	menu.Append("Editar Rola","app.editaRola")
	menu.Append("Editar Album", "app.editaAlbum")
	menu.Append("Editar Intérprete", "app.editaInterprete")
	return menu
}

func creaBarraBusqueda() (*gtk.SearchEntry) {
	barraBusqueda, err := gtk.SearchEntryNew()
	if err != nil {
		log.Fatal("No se pudo crear la barra de búsqueda: ",err)
	}
	barraBusqueda.Connect("activate", func() {
		fmt.Println(barraBusqueda.GetText())
	})
	return barraBusqueda
}

func creaDialog(titulo string) (*gtk.Dialog, *gtk.Box) {
	dialog, err := gtk.DialogNew()
	if err != nil {
		log.Fatal(err)
	}
	dialog.SetTitle(titulo)
	caja, err := dialog.GetContentArea()
	if err != nil {
		log.Fatal("No se pudo obtener el contentArea de Dialog: ",err)
	}
	dialog.Connect("destroy", func() {
		dialog.Destroy()
	})
	dialog.SetModal(true)
	return dialog,caja
}

func creaHeaderBar() (*gtk.HeaderBar){
	header, err := gtk.HeaderBarNew()
	if err != nil {
		log.Fatal("No se pudo crear el Header Bar:",err)
	}
	return header
}

func creaListBox() (*gtk.ListBox) {
	campos, err := gtk.ListBoxNew()
	if err != nil {
		log.Fatal("No se pudo crear la ListBox: ", err)
	}
	return campos
}

func creaRenglonListBox() (*gtk.ListBoxRow) {
	renglon, err := gtk.ListBoxRowNew()
	if err != nil {
		log.Fatal("No se pudo crear el renglon: ", err)
	}
	return renglon
}

func creaCaja() (*gtk.Box) {
	caja, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 80)
	if err != nil {
		log.Fatal("No se pudo crear la caja: ", err)
	}
	return caja
}

func creaEtiqueta(titulo string) (*gtk.Label) {
	etiqueta, err := gtk.LabelNew(titulo)
	if err != nil {
		log.Fatal("No se pudo crear la etiqueta: ", err)
	}
	return etiqueta
}

func creaEntrada(contenido string) (*gtk.Entry) {
	entrada, err := gtk.EntryNew()
	if err != nil {
		log.Fatal(err)
	}
	entrada.SetText(contenido)
	return entrada
}
