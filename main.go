package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

//  Test data that will eventually be handled by backend
var todos  = []string{
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

type templateData struct{
	Todos []string
	MetaData pageVariables
}

type pageVariables struct{
	Date string
	Time string
}

var now = time.Now()
var homePageVars = pageVariables{
	Date: now.Format("02-01-2006"),
	Time: now.Format("15:04:05"),
}

var appData = templateData{
	Todos: todos,
	MetaData: homePageVars,
}


func main() {
	fmt.Println(appData.MetaData.Date)
	//tpl returns pointer to a template if no err
	tpl, err := template.ParseFiles("temps/tpl.gohtml", "temps/footnote.gohtml", "temps/header.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	//container := []interface{}{ todos}
	//which then's gives us the ability to call Execute!!
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer nf.Close()
	//Instead of execute use executeTemplates and pass temp names for
	//specificity's sake. and then instead of parseFiles use parseTemplates
	err = tpl.Execute(nf, appData)
	if err != nil {
		log.Fatalln(err)
	}
}
