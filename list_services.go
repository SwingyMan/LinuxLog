package main

import (
	"log"
	"os/exec"
	"strings"
)

func list_services() []string {
	command, err := exec.Command("systemctl", "--type=service").Output()
	if err != nil {
		log.Fatal(err)
	}
	services := strings.Split(string(command), "\n")
	return services[1 : len(services)-7]
}
func get_single_service(index int) string {
	list := list_services()
	for i, s := range list {
		if i == index {
			return s
		}
	}
	return ""
}
