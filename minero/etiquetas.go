package main

import (
	"github.com/bogem/id3v2"
    "strings"
)

func GetTagInterprete(tag *id3v2.Tag) string {
    interprete := tag.GetTextFrame("TPE1").Text
    interprete = strings.TrimSpace(interprete)
    if interprete == "" {
        interprete = "UNKNOWN"
    }
    return interprete
}

func GetTagTitulo(tag *id3v2.Tag) string {
    titulo := tag.GetTextFrame("TIT2").Text
    titulo = strings.TrimSpace(titulo)
    if titulo == "" {
        titulo = "UNKNOWN"
    }
    return titulo
}

func GetTagAlbum(tag *id3v2.Tag) string {
    album := tag.GetTextFrame("TALB").Text
    album = strings.TrimSpace(album)
    if album == ""{
        album = "UNKNOWN"
    }
    return album
}

func GetTagFechaGrabacion(tag *id3v2.Tag) string {
    fechaGrabacion := tag.GetTextFrame("TDRC").Text
    fechaGrabacion = strings.TrimSpace(fechaGrabacion)
    if fechaGrabacion == "" {
        fechaGrabacion = "2018"
    }
    return fechaGrabacion
}

func GetTagGenero(tag *id3v2.Tag) string {
    genero := tag.GetTextFrame("TCON").Text
    genero = strings.TrimSpace(genero)
    if genero == "" {
        genero = "UNKNOWN"
    }
    return genero
}

func GetTagTrack(tag *id3v2.Tag) string {
    track := tag.GetTextFrame("TRCK").Text
    track = strings.TrimSpace(track)
    if track == ""{
        track = "0"
    }
    return track
}