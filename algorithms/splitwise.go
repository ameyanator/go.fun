package main

import (
	"fmt"
	"math"
)

type Transaction struct {
	Debtor   string
	Creditor string
	Amount   float64
}

func settleDebts(netBalances map[string]float64) []Transaction {
	var transactions []Transaction

	for {
		var debtor, creditor string
		minDebt := math.Inf(1)
		// Find debtor and creditor with minimum and maximum balances
		for user, balance := range netBalances {
			if balance < minDebt {
				minDebt = balance
				debtor = user
			}
		}

		maxCredit := math.Inf(-1)
		for user, balance := range netBalances {
			if balance > maxCredit && user != debtor {
				maxCredit = balance
				creditor = user
			}
		}

		if minDebt == 0 || maxCredit == 0 {
			break
		}
		fmt.Println("minDebt", minDebt, "maxCredit", maxCredit)
		amount := math.Min(math.Abs(minDebt), math.Abs(maxCredit))
		transactions = append(transactions, Transaction{Debtor: debtor, Creditor: creditor, Amount: amount})

		netBalances[debtor] += amount
		netBalances[creditor] -= amount
	}

	return transactions
}

func main() {
	netBalances := map[string]float64{
		"User1":  50.0,
		"User2":  -20.0,
		"User3":  -30.0,
		"User4":  0.0,
		"Ameya":  100.0,
		"Soumya": -30.0,
		"Minti":  -70.0,
	}

	transactions := settleDebts(netBalances)

	// Print transactions or perform further actions
	for _, transaction := range transactions {
		fmt.Printf("%s pays %.2f to %s\n", transaction.Debtor, transaction.Amount, transaction.Creditor)
	}
}
