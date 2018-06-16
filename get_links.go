package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

type Links struct {
	linsk []string
}

func getLinks() *cobra.Command {
	return &cobra.Command{
		Use:  "union [enlace de union fansub]",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var cadena string
			cadena = args[0]
			s := strings.Contains(cadena, "&pid")
			if s {
				subcad := cadena[0:45]
				links := Scraper(subcad)
				var enlacesLinpios []string
				enlacesLinpios = clean(links)
				for _, en := range enlacesLinpios {
					fmt.Println(en)
				}
			} else {
				var links = Scraper(cadena)
				var enlacesLimpios []string
				enlacesLimpios = clean(links)
				fmt.Println(len(enlacesLimpios))
				for i := 0; i < len(enlacesLimpios); i++ {
					fmt.Println(enlacesLimpios[i])
				}
			}

			return nil
		},
	}
}
func Scraper(cadena string) []string {
	var slice []string
	/*
			respuesta, err := http.Get(cadena)
		if err != nil {
			log.Fatal(err)
		}
		defer respuesta.Body.Close()
	*/
	doc, err := goquery.NewDocument(cadena)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(doc.Html())
	doc.Find(".spoil section").Each(func(i int, selection *goquery.Selection) {
		linkTang := selection.Find("a")
		link, _ := linkTang.Attr("href")

		slice = append(slice, link)
	})
	return slice
}

func clean(array []string) []string {
	var newSlice []string
	//2var nuevaCad string
	//var substring string

	for i := 0; i < len(array); i++ {
		var limpiar string
		limpiar = array[i]
		//http://out.unionfansub.com/3096066/mega.co.nz/#!FEFQ0TbZ!SMCWX2akyIc0KPD18XSHRTcvmIU5KpWHYeAih_z680Y
		compararCad := "http://out.unionfansub.com/3096066/mega.co.nz/#!FEFQ0TbZ!SMCWX2akyIc0KPD18XSHRTcvmIU5KpWHYeAih_z680Y"
		if len(limpiar) >= len(compararCad) {
			var subCad = limpiar[35:len(limpiar)]
			var nueva = "http://" + subCad
			newSlice = append(newSlice, nueva)
		} else {
			newSlice = append(newSlice, limpiar)
		}

	}
	return newSlice
}
