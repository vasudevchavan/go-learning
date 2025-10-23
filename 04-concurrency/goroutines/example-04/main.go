package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzaMade, pizzaFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1

		color.Blue("Received order #%d", pizzaNumber)

		rnd := rand.Intn(5) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzaFailed++
		} else {
			pizzaMade++
		}

		total++

		color.Yellow("Making pizza #%d, it will take %d seconds...", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd < 2 {
			msg = fmt.Sprintf("❌ Pizza order #%d failed: over-heating", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("❌ Pizza order #%d failed: ingredient issue", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("✅ Pizza order #%d is ready!", pizzaNumber)
		}

		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
	}
	return nil
}

func pizzeria(pizzaMaker *Producer) {
	var i = 0
	for {
		curPizza := makePizza(i)
		if curPizza == nil {
			return
		}
		i = curPizza.pizzaNumber
		select {
		case pizzaMaker.data <- *curPizza:
		case quitChan := <-pizzaMaker.quit:
			close(pizzaMaker.data)
			close(quitChan)
			return
		}
	}
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	color.Cyan("🍕 Pizza shop is open!")
	color.Cyan("----------------------")

	// Create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// Start the pizzeria goroutine
	go pizzeria(pizzaJob)

	// Receive results
	for i := 0; i < NumberOfPizzas; i++ {
		order := <-pizzaJob.data
		if order.success {
			color.Green(order.message)
		} else {
			color.Red(order.message)
		}
	}

	color.Cyan("----------------------")
	color.Cyan("🍕 Pizza shop is closed.")
	color.Cyan("✅ Total pizzas made: %d", pizzaMade)
	color.Cyan("❌ Total pizzas failed: %d", pizzaFailed)
	color.Cyan("📦 Total orders processed: %d", total)
}
