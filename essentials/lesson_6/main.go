package main

import "fmt"

func main() {

	accountBalance := 1000.00
	fmt.Println("Welcome to the bank app !")

	for {

		fmt.Println("What do you want to do ? ")
		fmt.Println("1. Check Account Balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdrawl money")
		fmt.Println("4. Exit")

		var choose int
		fmt.Print("Choose one option(1-4):")
		fmt.Scan(&choose)

		switch choose {

		case 1:
			fmt.Println("accountBalance:", accountBalance)

		case 2:
			var deposit float64
			fmt.Print("Enter money you want to deposit: ")
			fmt.Scan(&deposit)

			if deposit <= 0 {
				fmt.Print("Invalid Amount! should be greater that 0")
				continue
			}

			accountBalance = accountBalance + deposit
			fmt.Println("Deposit amount:", deposit)
			fmt.Println("Total amount:", accountBalance)

		case 3:
			var withdrawl float64
			fmt.Print("Enter withdrawl amount: ")
			fmt.Scan(&withdrawl)

			if withdrawl <= 0 {
				fmt.Printf("Invalid Amount! should be greater that 0")
				continue
			}

			if withdrawl > accountBalance {
				fmt.Printf("You cannot withdraw more than account balance. AccountBalance: %v", accountBalance)
				continue
			}

			accountBalance = accountBalance - withdrawl
			fmt.Println("Withdrawl amount:", withdrawl)
			fmt.Println("Total amount:", accountBalance)

		default:
			fmt.Println("Bye, Have a good day!")
			fmt.Println("Dubara ana bank mai !")
			// break
			return

		}

		// 	if choose == 1 {

		// 		fmt.Println("accountBalance:", accountBalance)
		// 	} else if choose == 2 {

		// 		var deposit float64
		// 		fmt.Print("Enter money you want to deposit: ")
		// 		fmt.Scan(&deposit)

		// 		if deposit <= 0 {
		// 			fmt.Print("Invalid Amount! should be greater that 0")
		// 			continue
		// 		}

		// 		accountBalance = accountBalance + deposit
		// 		fmt.Println("Deposit amount:", deposit)
		// 		fmt.Println("Total amount:", accountBalance)
		// 	} else if choose == 3 {

		// 		var withdrawl float64
		// 		fmt.Print("Enter withdrawl amount: ")
		// 		fmt.Scan(&withdrawl)

		// 		if withdrawl <= 0 {
		// 			fmt.Printf("Invalid Amount! should be greater that 0")
		// 			continue
		// 		}

		// 		if withdrawl > accountBalance {
		// 			fmt.Printf("You cannot withdraw more than account balance. AccountBalance: %v", accountBalance)
		// 			continue
		// 		}

		// 		accountBalance = accountBalance - withdrawl
		// 		fmt.Println("Withdrawl amount:", withdrawl)
		// 		fmt.Println("Total amount:", accountBalance)

		// 	} else if choose == 4 {

		// 		fmt.Println("Bye, Have a good day!")
		// 		break
		// 	}

	}

}
