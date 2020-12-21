package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type config struct {
	Address string
	//Port	int
	Jobs []job
}

type job struct {
	Country string
	Status  string
	Path    string
	Title   string
	Line    string
	Time    string
}

func loadConfig(path string) config {
	c := config{}
	d, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(d, &c)
	if err != nil {
		panic(err)
	}
	return c
}
