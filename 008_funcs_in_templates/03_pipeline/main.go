package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func double(f float64) float64 {
	return f * 2
}

func square(f float64) float64 {
	return math.Pow(f, 2)
}

func squareRoot(f float64) float64 {
	return math.Sqrt(f)
}

var fm = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": squareRoot,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 50.0)
	if err != nil {
		log.Fatalln(err)
	}
}
