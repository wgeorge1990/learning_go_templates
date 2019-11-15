package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	//tpl returns pointer to a template if no err
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	quotes := []string{
		"To be or not to be",
		"Dont worry be happy",
		"Tomatoes Potatoes",
		"build me a bridge and then cross it",
	}

	weekDays := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	container := []interface{}{weekDays, quotes}
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
