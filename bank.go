package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceAccountText = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)                          //creates new string
	os.WriteFile(balanceAccountText, []byte(balanceText), 0644) //storing balance in a file

}

func getBalanceFromFile() (float64, error) { //reading value from value
	data, err := os.ReadFile(balanceAccountText)

	if err != nil {
		return 100, errors.New("Failed to read from file")
	}
	balanceText := string(data)                         //converting byte to string
	balance, err := strconv.ParseFloat(balanceText, 64) //converting string to float64

	if err != nil {
		return 100, errors.New("Failed to parse stored balance values")
	}
	return balance, nil

}

func main() {
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}
	fmt.Println("Welcome to the bank!")

	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit money ")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		//wantsCheckBalance := choice == 1

		//using switch statement

		/*switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New Amount is: ", accountBalance)
		case 3:
			fmt.Print("Your withdraw: ")
			var withdraw float64
			fmt.Scan(&withdraw)

			if withdraw <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			if withdraw > accountBalance {
				fmt.Println("Insuifficient balance")
				continue
			}
			accountBalance += withdraw
			fmt.Println("Balance updated! New Amount is: ", accountBalance)
		default:
			fmt.Println("Closing")

		}*/

		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)

		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New Amount is: ", accountBalance)
		} else if choice == 3 {
			fmt.Print("Your withdraw: ")
			var withdraw float64
			fmt.Scan(&withdraw)

			if withdraw <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			if withdraw > accountBalance {
				fmt.Println("Insuifficient balance")
				continue
			}
			accountBalance += withdraw
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New Amount is: ", accountBalance)
		} else {
			fmt.Println("Closing")
			break
		}

	}

	fmt.Println("Thanks for visiting!")

}
