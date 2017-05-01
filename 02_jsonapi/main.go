package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/codenut247/golang_website/02_jsonapi/app/shared/database"
	"github.com/codenut247/golang_website/02_jsonapi/app/shared/jsonconfig"
	"github.com/codenut247/golang_website/02_jsonapi/app/shared/server"
)

type configuration struct {
	Database database.Info `json:"Database"`
	Server   server.Server `json:"Server"`
}

var config = &configuration{}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

func main() {
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)
	database.Connect(config.Database)

	f := foo{Name: "Mr. Anderson"}
	server.Run(&f, config.Server)
}

type foo struct {
	Name string
}

func (f *foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", f.Name)
}
