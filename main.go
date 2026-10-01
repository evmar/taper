package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Performance struct {
	URL    string  `toml:"url"`
	Title  string  `toml:"title"`
	Date   string  `toml:"date"`
	Tracks []Track `toml:"tracks"`
}

type Track struct {
	Time string `toml:"time"`
	Name string `toml:"name"`
}

func run() error {
	flag.Usage = func() {
		fmt.Printf("Usage: %s <input.toml>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}

	data, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		return err
	}

	var performance Performance
	if err := toml.Unmarshal(data, &performance); err != nil {
		return err
	}

	fmt.Printf("%+v\n", performance)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
