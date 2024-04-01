package main

import (
	"fmt"
	"time"

	"github.com/Halim098/golang-shop-basic/entity"
)

func main() {
	var customer = entity.NewCustomer()
	var choose int
	var acc string
	var loop = true

	for loop {
		if acc == "" {
			fmt.Println("1. Login")
			fmt.Println("2. Register")
			fmt.Println("3. Exit")
			fmt.Println("==============")
			fmt.Print("Choose: ")
			fmt.Scanln(&choose)
			switch choose {
			case 1:
				fmt.Println("===== Login =====")
				fmt.Print("Enter your username: ")
				name := ""
				fmt.Scanln(&name)
				fmt.Print("Enter your password: ")
				pass := ""
				fmt.Scanln(&pass)
				fmt.Println("")
				acc2, msg, err := customer.Login(name, pass)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(msg)
					// acc = acc2
					fmt.Println("Hi, ", acc2)
				}
				time.Sleep(3 * time.Second)

			case 2:
				fmt.Println("===== Register =====")
				fmt.Print("Enter your username: ")
				name := ""
				fmt.Scanln(&name)
				fmt.Print("Enter your password: ")
				pass := ""
				fmt.Scanln(&pass)
				fmt.Println("")
				msg, err := customer.Register(name, pass)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(msg)
				}
				time.Sleep(3 * time.Second)

			case 3:
				fmt.Println("Bye")
				loop = false

			default:
				fmt.Println("Invalid input")
			}
		}

		if acc != "" {
			fmt.Println("Hi, ", acc)
			fmt.Println("==============")
			fmt.Println("1. Add to cart")
			fmt.Println("2. Show cart")
			fmt.Println("3. Checkout")
			fmt.Println("4. Logout")
			fmt.Println("==============")
			fmt.Print("Choose: ")
			fmt.Scanln(&choose)
			switch choose {
			case 1:

			case 2:

			case 3:

			case 4:
				fmt.Println("Thanks :)")
				acc = ""
			default:
				fmt.Println("Invalid input")
			}
		}
	}
}
