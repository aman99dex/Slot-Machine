package main

import (
	"fmt"
)

func GetName() string {
	name := ""
	fmt.Println("Welcome to Aman's Casino.....!!!!! ")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s, Let's Play the Game\n", name)
	return name
}

func GetBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Enter your bet, or 0 to quit (balance =$%d): ", balance)
		fmt.Scan(&bet)
		//fix the bug i.e. if user enters a string, the program will crash
		if bet > balance {
			fmt.Printf("You cannot bet more than you have in your balance\n")
		} else {
			break
		}
	}
	return bet
}
