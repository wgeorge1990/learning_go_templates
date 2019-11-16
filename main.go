package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)



func main() {
	//tpl returns pointer to a template if no err
	tpl, err := template.ParseFiles("temps/tpl.gohtml", "temps/footnote.gohtml", "temps/header.gohtml")
	//tpl := template.Must(template.New("new").Funcs(fm).ParseFiles("temps/tpl.gohtml", "temps/footnote.gohtml", "temps/header.gohtml"))
	//tpl, err := template.ParseFiles("temps/tpl.gohtml", "temps/footnote.gohtml", "temps/header.gohtml")
	//tpl, err := template.ParseGlob("temps/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	todos := []string{
		"Go sky diving",
		"Dont worry be happy",
		"Go deep sea diving",
		"Go to Santorini",
		"Bleach my hair",
		"Buy a race car",
		"Travel to the Netherlands",
		"Thru hike the Appalachian trail",
		"Go on a safari",
		"Learn  to fly a plane",

	}

	container := []interface{}{ todos}
	fmt.Println(container)
	//which then's gives us the ability to call Execute!!
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer nf.Close()
	//Instead of execute use executeTemplates and pass temp names for
	//specificity's sake. and then instead of parseFiles use parseTemplates
	//err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", quotes)
	err = tpl.Execute(nf, container)
	if err != nil {
		log.Fatalln(err)
	}
}
