package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type LocaleMessage struct {
	Name    string `yaml:"name"`
	Message string `yaml:"message"`
}

func main() {
	var msgs []*LocaleMessage
	buf, err := ioutil.ReadFile("locale.yml")
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(buf, &msgs); err != nil {
		panic(err)
	}

	for _, msg := range msgs {
		log.Printf("%+v\n", msg)
	}
}
