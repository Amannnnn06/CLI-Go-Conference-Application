package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUsersInput()

	isValueName, isValidEmail, isValidTicketsNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValueName && isValidEmail && isValidTicketsNumber {

		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)

		firstNamelist := getFirstNames()
		fmt.Printf("The First Names of Bookings are : %v\n ", firstNamelist)
		//call the function for firstname print

		noTicketsRemaining := remainingTickets == 0

		if noTicketsRemaining {
			//end Program
			fmt.Printf("Our conference is booked out , come back next years.")
			//break
		}

	} else {
		if !isValueName {
			fmt.Println("First name or last name your entered is too short ")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign ")
		}
		if !isValidTicketsNumber {
			fmt.Println("Number Of Tickets you entered is Invalid ")
		}

	}
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome To Our %v !!!\n", conferenceName)
	fmt.Printf("We have Total of %v tickets and %v are still avaiable. \n", conferenceTickets, remainingTickets)
	fmt.Printf("Get Your Tickets here to attend\n ")

}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func getUsersInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your First Name : ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last  Name : ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address  : ")
	fmt.Scan(&email)

	fmt.Printf("Enter your tickets  : ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	//create a map for user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is %v \n ", bookings)

	fmt.Printf("Thank you ,\n %v %v for booking  %v tickets. \n You will recieve a confirmation email at  %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)

	var ticket = fmt.Sprintf("%v Tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending Tickets %v to email address %v ", ticket, email)
	fmt.Println("########################")
	wg.Done()
}
