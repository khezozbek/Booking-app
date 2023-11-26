package main

import (
	"booking-app/actions"
	"fmt"
	"strings"
)

var conf_name string = "Go Conference"
const conf_tickets = 50
var remainingTickets uint = 50
var bookings = []string{}
var city string
var name string
var last_name string
var tichektscount uint
var email string

func main() {

	greetUsers()

	fmt.Printf("Enter your location: ")
	fmt.Scan(&city)

	switch city {
		case "Tashkent":
			for {
				isValidName, isVaildEmail, isVaildTick := actions.ValidateUserInputs(name, last_name, tichektscount, email, remainingTickets)
				name, last_name, tichektscount, email := userInputs()
				if isValidName && isVaildEmail && isVaildTick{
					bookingtick()
				} else {
					if !isValidName {
						fmt.Printf("Your first name or last name has a proble check it %v %v", name, last_name)
						continue
					}
					if !isVaildEmail {
						fmt.Printf("Your email %v is incorrect\n", email)
					}
					if !isVaildTick{
						fmt.Printf("Your input data is incorrect")
					} else {
						fmt.Printf("Dear, %v,  %v %v is incorrect\n", name, email, tichektscount)
					}
				}
			}
		case "Singapure", "London":
			fmt.Printf("Sorry, right now we have no tickets for %v", city)
			break
		case "Berlin":
			fmt.Printf("Sorry, right now we have no tickets for %v", city)
			break
		default:
			fmt.Print("No valid city")
	}
}

func greetUsers() {
	fmt.Printf("Welcom to %v our app\n", conf_name)
	fmt.Println("We have total of", conf_tickets, "tickets and", remainingTickets, "are still avabile\n")	
}

func frisnameofUsers() []string {
	firstnames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstnames = append(firstnames, names[0])
	}
	return firstnames
}


func userInputs() (string, string, uint, string) {
	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)
	
	fmt.Printf("Enter your last name: ")
	fmt.Scan(&last_name)
	
	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)
	
	fmt.Printf("how many tickets do you want to buy? ")
	fmt.Scan(&tichektscount)	
	return name, last_name, tichektscount, email
}

func bookingtick() (uint, []string, string, string, uint, string, string) {
	remainingTickets = remainingTickets - tichektscount
	bookings = append(bookings,name + " " + last_name)
	fmt.Printf("Thank you %v %v for booking %v tickets. You have to coniform your email at %v\n", name, last_name, tichektscount, email)
	firstanames := frisnameofUsers()
	fmt.Printf("%v, booking a %v ticket!\n", firstanames, tichektscount)
	fmt.Printf("%v tichkets remaining for %v\n", remainingTickets, conf_name)

	return remainingTickets, bookings, name, last_name, tichektscount, email, conf_name
}