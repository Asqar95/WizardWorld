package main

import (
	"WizardWorld/wizard"
	"fmt"
	"log"
	"time"
)

func main() {
	wizardClient, err := wizard.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}
	elixirs, err := wizardClient.GetElixirs()
	if err != nil {
		log.Fatal(err)
	}
	for _, el := range elixirs {
		fmt.Println(el.Info())
	}

	id, err := wizardClient.GetElixir("0106fb32-b00d-4d70-9841-4b7c2d2cca71")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id.Info())
}
