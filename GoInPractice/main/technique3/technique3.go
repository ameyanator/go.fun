package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kylelemons/go-gypsy/yaml"
	"gopkg.in/gcfg.v1"
)

type configuration struct {
	Enabled bool
	Path    string
}

func technique1() {
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}

	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(conf.Path)
}

func technique2() {
	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))
}

func technique3() {
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}

	err := gcfg.ReadFileInto(&config, "config.ini")
	if err != nil {
		fmt.Printf("Failed to parse config file: %s", err)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
}

func main() {
	// technique1()
	// technique2()
	technique3()
}
