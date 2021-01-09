package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

var configPath string
var conf config

var wg sync.WaitGroup

func main() {
	// Load config
	configPath = *flag.String("config", "./config.yml", "Config path.")
	flag.Parse()
	conf = loadConfig(configPath)
	conf.Address = *flag.String("a", conf.Address, "Address.")
	flag.Parse()

	fmt.Println("Succesfully loaded config!")

	//Enumerate jobs
	for _, v := range conf.Jobs {
		go work(v)
	}

	if conf.Address != "" {
		err := generateIndex(conf.Jobs)
		if err != nil {
			panic(err)
		}
		//Start webserver
		err = host()
		if err != nil {
			panic(err)
		}
	} else {
		wg.Add(1)
		wg.Wait()
	}
}

func generateIndex(js []job) error {
	os.Create("./static/index.html")
	f, err := os.OpenFile("./static/index.html", os.O_WRONLY, 0664)
	defer f.Close()
	if err != nil {
		return err
	}

	var index string = "<!DOCTYPE html><html><title>CovidGraph</title><body>"

	for _, v := range js {
		index += fmt.Sprintf("<img src='%s'>", v.Path[1:])
	}

	index += "</body></html>"

	f.WriteString(index)
	return nil
}

func work(j job) {
	defer wg.Done()
	dur, err := time.ParseDuration(j.Time)
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println("Started work on " + j.Title)
		data, err := getDayOne(j.Country, j.Status)
		if err != nil {
			panic(err)
		}

		err = drawGraph(data, j.Title, j.Line, j.Path)
		if err != nil {
			panic(err)
		}

		time.Sleep(dur)
	}
}
