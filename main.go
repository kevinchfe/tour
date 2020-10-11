package main

import (
	"flag"
	"log"
)

var name string

func main() {
	flag.Parse()

	// fmt.Println(os.Args[:])

	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "go语言", "帮助信息")
	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	phpCmd.StringVar(&name, "name", "php语言", "帮助信息")

	args := flag.Args()
	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "php":
		_ = phpCmd.Parse(args[1:])
	}

	log.Printf("name: %s", name)
}
