package main

import "fmt"

func main() {
	records := make([][]string, 0)
	//student1
	student1 := make([]string, 4)
	student1[0] = "Jorge"
	student1[1] = "Saito"
	student1[2] = "100.00"
	student1[3] = "74.00"
	//store the records
	records = append(records, student1)
	//student2
	student2 := make([]string, 4)
	student2[0] = "Sandra"
	student2[1] = "Willy Lo"
	student2[2] = "99.00"
	student2[3] = "92.00"
	records = append(records, student2)
	fmt.Println(records)

	transactions := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {
		transaction := make([]int, 0)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
		fmt.Println(transactions)
	}

}
