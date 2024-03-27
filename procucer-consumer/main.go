package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func main() {
	// seed the random numer generator
	rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzaria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzaria(pizzaJob)

	// create and run consumer

	// print out the ending message
}

func pizzaria(pizzaMaker *Producer) {
	for {

	}
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}
