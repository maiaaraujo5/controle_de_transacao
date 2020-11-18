package main

import (
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/fx/module/runner"
	"log"
)

func main() {
	err := runner.RunApplication()
	if err != nil {
		log.Fatal(err)
	}
}
