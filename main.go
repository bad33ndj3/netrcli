package main

import (
	"flag"
	"github.com/bgentry/go-netrc/netrc"
	"log"
)

const (
	PartName     = "name"
	PartLogin    = "login"
	PartPassword = "password"
	PartAccount  = "account"
)

func main() {
	var path string
	flag.StringVar(&path, "netrc", "", "Path to the netrc file")
	var machineName string
	flag.StringVar(&machineName, "machine", "", "Machine name")
	var part string
	flag.StringVar(&part, "part", "", "Part to print")
	flag.Parse()

	if path == "" {
		log.Fatal("Path to netrc file is required")
	}

	if machineName == "" {
		log.Fatal("Machine name is required")
	}

	if part == "" {
		log.Fatal("Part is required")
	}

	netrc, err := netrc.ParseFile(path)
	if err != nil {
		log.Fatal("Unable to read netrc")
	}

	machine := netrc.FindMachine(machineName)
	if machine == nil {
		log.Fatalf("Machine '%s' not found", machineName)
	}

	switch part {
	case PartLogin:
		log.Println(machine.Login)
	case PartPassword:
		log.Println(machine.Password)
	case PartAccount:
		log.Println(machine.Account)
	case PartName:
		log.Println(machine.Name)
	default:
		log.Fatalf("Invalid part: %s", part)
	}
}
