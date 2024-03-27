package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount float64
}

var wg sync.WaitGroup

func main() {
	var bankBalance float64
	var mutex sync.Mutex

	fmt.Printf("Initial account balance: $%f", bankBalance)
	fmt.Println()

	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	} 

	wg.Add(len(incomes))

	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := range 52 {
				mutex.Lock()
				
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp

				mutex.Unlock()	

				fmt.Printf("On week %d, you earned $%f from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}
	wg.Wait()

	fmt.Printf("Final bank balance: $%.2f", bankBalance)
}