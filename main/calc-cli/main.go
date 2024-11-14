package main

import (
	"flag"
	"log"
	"os"

	"github.com/smkartch/calc-apps/handlers"
	"github.com/smkartch/calc-lib"
)

func main() {
	var operation string
	flag.StringVar(&operation, "op", "+", "The mathematical operation.")
	flag.Parse()
	handler := handlers.NewHandler(os.Stdout, calculators[operation])
	err := handler.Handle(flag.Args())
	if err != nil {
		log.Fatal(err)
	}
}

var calculators = map[string]handlers.Calculator{
	"+": &calc.Addition{},
	"-": &calc.Subtraction{},
	"*": &calc.Multiplication{},
	"/": &calc.Division{},
}
