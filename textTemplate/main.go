package main

import (
	"html/template"
	"os"
)

func main() {
	t1 := template.New("t1")

	t1, err := t1.Parse("Value is {{.}}\n")

	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	t1.Execute(os.Stdout, "Some Text")
	t1.Execute(os.Stdout, 35)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Python",
		"Java",
		"C++",
		"C",
	})

	create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Farid kaki"})

	t2.Execute(os.Stdout, map[string]string{
		"Name": "Asma azizi",
	})

	t3 := create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")

	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")

	t4.Execute(os.Stdout,
		[]string{
			"farid",
			"babak",
			"asma",
			"ahamd",
		})
}
