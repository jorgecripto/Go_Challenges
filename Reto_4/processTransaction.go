package main

func ProcessTransaction(transactions []Transaction, procesor func(float64) float64) float64 {
	var total float64
	for _, record := range transactions {
		total += procesor(record.Amount)

	}

	return total

}