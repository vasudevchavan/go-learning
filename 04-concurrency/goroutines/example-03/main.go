package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

func main() {

	var bankBalance int
	var wg sync.WaitGroup
	var balanceMutex sync.Mutex
	fmt.Printf("Initial Balance in account: $%d.00", bankBalance)
	fmt.Println()

	incomes := []Income{
		{Source: "Main job", Amount: 10},
		{Source: "Gift job", Amount: 1},
		{Source: "Stock job", Amount: 2},
	}
	// wg.Add(len(incomes))

	for _, singleincome := range incomes {
		wg.Add(1)
		go func(singleincome Income) {
			defer wg.Done()
			for week := 1; week < 11; week++ {
				balanceMutex.Lock()
				bankBalance += singleincome.Amount
				balanceMutex.Unlock()
				fmt.Printf("Earning for week %d  from %d is %s\n", week, singleincome.Amount, singleincome.Source)
			}
		}(singleincome)
	}
	wg.Wait()

	fmt.Printf("Final Bank Balance is %d", bankBalance)

}
