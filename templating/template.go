package main

import (
	"html/template"
	"os"
)

type RespJson struct {
	Token     string `json:"token"`
	ClientId  string `json:"clientId"`
	SessionId string `json:"sessionId"`
	Hostname  string
	User      string
	Password  template.HTML
}

func main() {
	var err error
	data := RespJson{"n2103j123n1293h921h.12n39123h2o13", "0001-0023-2223-1123", "be083af5-c8ee-41b7-52f3-c8a46946547d", "google.xom", "DCovPywTKiY5LgolLiYsKCI/MywlBRUTdxAAD24ZHwANGwAOCQ==", "op123j0123n2183h12093nl&&"}

	tpl, err := template.ParseFiles("settings.json.tmpl")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("settings.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tpl.Execute(f, data)
	if err != nil {
		panic(err)
	}

}
