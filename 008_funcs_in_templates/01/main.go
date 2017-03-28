package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

// "create a FuncMap to register functions.
// "uc" - what the func will be called in the template
// "uc" - the ToUpper function in the strings package
// "ft" - what the func will be called in the template
// "ft" - slices a string and returns the first three chars
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}
	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed",
	}
	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}
	muhammed := sage{
		Name:  "Muhammed",
		Motto: "To overcome evil with good is good",
	}

	sages := []sage{buddha, gandhi, mlk, jesus, muhammed}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
