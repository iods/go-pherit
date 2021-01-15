package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func loops() {
	//text := "before: {{.}}\n{{range .}}{{end}}after: {{.}}\n"
	text := "before: {{.}}\n{{range .}}in: {{.}}\n{{end}}after: {{.}}\n"
	executeTemplate(text, []string{"one", "two", "three"})
}

func prices() {
	text := "Prices: \n{{range .}}${{.}}\n{{end}}"
	executeTemplate(text, []float64{33.3, 11, 9.8, 3.14})
}

func main() {
	text := "Here is my template start\nAction: {{.}}\nTemplate ended.\n"
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, "one")
	check(err)
	err = tmpl.Execute(os.Stdout, "two")
	check(err)
	executeTemplate("Dot is {{.}}!\n", "fuck yeah")
	executeTemplate("Dot is {{.}}\n", 199.55)
	executeTemplate("Dot is {{if .}}true{{end}}!\n", true)
	executeTemplate("Dot is {{if .}}true{{end}}\n", false)
	loops()
	prices()
}