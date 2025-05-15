package main

import (
	"fmt"
	fileops "lesson_2/fileops"

	"github.com/Pallinder/go-randomdata"
)

func main() {

	accountBalance, err := fileops.GetFloatFromFile()

	if err != nil {

		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------")
		// panic("cannot procced further ! Exiting ...")          ----> Better way than using return. exits with a message and line no. where the panic is called.
	}

	fmt.Println("Welcome to the bank app !")
	fmt.Println("reach us 24/7", randomdata.PhoneNumber())

	for {

		userOptions()

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
			fileops.WriteFloatToFile(accountBalance)

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
			fileops.WriteFloatToFile(accountBalance)

		default:
			fmt.Println("Bye, Have a good day!")
			fmt.Println("Dubara ana bank mai !")
			// break
			return

		}

	}

}
