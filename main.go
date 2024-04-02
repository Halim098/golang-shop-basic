package main

import (
	"fmt"

	"github.com/Halim098/golang-shop-basic/entity"
	"github.com/Halim098/golang-shop-basic/helper"
)

func main() {
	var cart = entity.NewCart()
	var customer = entity.NewCustomer()
	var choose int
	var acc string
	var loop = true

	for loop {
		if acc == "" {
			helper.ClearScreen()
			fmt.Println("1. Login")
			fmt.Println("2. Register")
			fmt.Println("3. Exit")
			fmt.Println("==============")
			fmt.Print("Choose: ")
			fmt.Scanln(&choose)
			switch choose {
			case 1:
				helper.ClearScreen()
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
					acc = acc2
				}
				helper.Delay(3)

			case 2:
				helper.ClearScreen()
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
				helper.Delay(3)

			case 3:
				helper.ClearScreen()
				fmt.Println("Bye")
				loop = false
			default:
				fmt.Println("Invalid input")
				helper.Delay(3)
			}
		}

		if acc != "" {
			helper.ClearScreen()
			fmt.Println("Hi, ", acc)
			fmt.Println("==============")
			fmt.Println("1. Add to cart")
			fmt.Println("2. Show cart")
			fmt.Println("3. Remove from cart")
			fmt.Println("4. Checkout")
			fmt.Println("5. Logout")
			fmt.Println("==============")
			fmt.Print("Choose: ")
			fmt.Scanln(&choose)
			switch choose {
			case 1:
				helper.ClearScreen()
				fmt.Println("===== Product =====")
				for i := range entity.DataProduct {
					fmt.Println("Neme :", entity.DataProduct[i].Name, "\nPrice :", entity.DataProduct[i].Price, "\nStock :", entity.DataProduct[i].Stock)
					fmt.Println("==============")
				}
				fmt.Print("Enter product name: ")
				name := ""
				fmt.Scanln(&name)
				fmt.Print("Enter quantity: ")
				quantity := 0
				fmt.Scanln(&quantity)
				fmt.Println("")
				err := cart.AddItem(name, quantity)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Success add to cart")
				}
				helper.Delay(3)
			case 2:
				helper.ClearScreen()
				fmt.Println("===== Show cart =====")
				for _, v := range cart.GetCart() {
					fmt.Println("Product :", v.Product.Name, "\nPrice :", v.Product.Price, "\nQuantity :", v.Quantity)
					fmt.Println("==============")
				}
				helper.Delay(5)
			case 3:
				helper.ClearScreen()
				fmt.Println("===== Remove from cart =====")
				for i, v := range cart.GetCart() {
					fmt.Println(i+1, "Product :", v.Product.Name)
				}
				fmt.Print("")
				fmt.Print("Enter product name: ")
				name := ""
				fmt.Scanln(&name)
				msg, err := cart.RemoveItem(name)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(msg)
				}
				helper.Delay(3)
			case 4:
				helper.ClearScreen()
				fmt.Println("===== Checkout =====")
				payment := entity.NewPayment()
				err := payment.AddItem(cart)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Total :", payment.Total)
					fmt.Println("=============")
				}
				fmt.Print("Enter your money: ")
				money := 0
				fmt.Scanln(&money)
				change, msg, err := payment.Pay(money)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Change :", change)
					fmt.Println("")
					fmt.Println(msg)
					cart.ResetCart()
				}
				helper.Delay(3)
			case 5:
				helper.ClearScreen()
				fmt.Println("Thanks :)")
				acc = ""
				helper.Delay(3)
			default:
				fmt.Println("Invalid input")
				helper.Delay(3)
			}
		}
	}
}
