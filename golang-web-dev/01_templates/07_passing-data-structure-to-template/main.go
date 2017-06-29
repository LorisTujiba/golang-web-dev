package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

//Human is exported
type Human struct {
	Name string
	Age  int
}

//Car is exporeted
type Car struct {
	Brand string
	Speed float64
}

//Collection is exported, composition
type Collection struct {
	H []Human
	C []Car
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	names := []string{"Maria", "Carissa", "Lidwina", "Tessa", "Claudia"}

	lists := map[string]string{
		"a": "A",
		"b": "B",
	}

	agus := Human{
		Name: "Agus",
		Age:  23,
	}

	loris := Human{
		Name: "Loris",
		Age:  22,
	}

	radeon := Car{
		Brand: "Radeon",
		Speed: 240.5,
	}

	blackBullet := Car{
		Brand: "Black Bullet",
		Speed: 260.23,
	}

	manusia := []Human{agus, loris}
	mobil := []Car{radeon, blackBullet}

	koleksi := Collection{manusia, mobil}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", names)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "withindex.gohtml", names)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "index.gohtml", lists)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", agus)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "rangestruct.gohtml", manusia)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "composition.gohtml", koleksi)
	if err != nil {
		log.Fatal(err)
	}
}
