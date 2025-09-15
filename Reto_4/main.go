package main

import "fmt"

func main() {

	transactions := []Transaction{

		{ID: "TRX001", Amount: 150.0},
		{ID: "TRX002", Amount: 50.0},
		{ID: "TRX003", Amount: 250.0},
		{ID: "TRX004", Amount: 100.0},
		{ID: "TRX005", Amount: 300.0},
	}

	finalTotalImports := ProcessTransaction(transactions, totalImports)
	fmt.Printf("El total acumulado de los importes de las transacciones es: %.2f\n", finalTotalImports)

	finalDobleTotalImports := ProcessTransaction(transactions, dobleTotalImports)
	fmt.Printf("El total acumulado del doble de los importes de las transacciones es: %.2f\n", finalDobleTotalImports)

	finalHowManyOver100 := ProcessTransaction(transactions, howManyOver100)
	fmt.Printf("El numero de transacciones que superan el importe es: %.0f\n", finalHowManyOver100)

	howManyOver200 := constructor(250)
	finalHowManyOver200 := ProcessTransaction(transactions, howManyOver200)
	fmt.Printf("El número de transacciones que supera el umbral de 200 es: %.0f\n", finalHowManyOver200)


}

func totalImports(amount float64) float64 {

	return amount

}

func dobleTotalImports(amount float64) float64 {

	return amount * 2

}

func howManyOver100(amount float64) float64 {

	if amount > 100.0 {
		return 1
	}

	return 0

}

func constructor(threshold float64) func (float64) float64  {
	return func(amount float64) float64{
		if amount > threshold {
			return 1
		}
		return 0
	}
	
}