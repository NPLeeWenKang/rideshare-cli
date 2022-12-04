package main

import (
	"fmt"
	"strings"
)

var userId string

func main() {
	for {
		fmt.Println()
		if len(userId) == 0 {
			option := menu()
			if option == "999" {
				return
			}
			tempUserId, err := confirmUser(option)
			if err != nil {
				fmt.Println("Error occured while retrieving users")
				return
			} else if tempUserId == "" {
				fmt.Println("Error occured: No user with ID")
				return
			}
			userId = tempUserId
		} else {
			userType := strings.ToLower(userId[0:1])
			if userType == "p" {
				fmt.Println("Sign In Successful")
				option := menuPassanger()
				if option == "1" {

				} else if option == "999" {
					userId = ""
				} else {
					fmt.Println("Invalid option")
				}
			} else if userType == "d" {
				fmt.Println("Sign In Successful")
				option := menuDriver()
				if option == "1" {

				} else if option == "999" {
					userId = ""
				} else {
					fmt.Println("Invalid option")
				}
			}
		}
	}
}
func menu() string {

	fmt.Println("========== Ride Share ==========")
	fmt.Println("Passanger & Driver Console")

	// Get all passangers
	allPassanger, err := getAllPassanger()
	if err != nil {
		fmt.Println("Error occured while retrieving Passangers")
	}
	for _, v := range allPassanger {
		fmt.Printf(" p%d. (Passanger) %s %s\n", v.Passanger_Id, v.First_Name, v.Last_Name)
	}

	// Get all drivers
	allDriver, err := getAllDriver()
	if err != nil {
		fmt.Println("Error occured while retrieving Passangers")
	}
	for _, v := range allDriver {
		fmt.Printf(" d%d. (Driver) %s %s\n", v.Driver_Id, v.First_Name, v.Last_Name)
	}
	fmt.Println(" 999. Quit")
	fmt.Print("Enter a user to sign in as (eg. c1): ")

	var option string
	// fmt.Scanf("%d", &option)
	fmt.Scanln(&option)
	return option
}

func confirmUser(userId string) (string, error) {
	userType := strings.ToLower(userId[0:1])
	if userType == "p" {
		id := strings.ReplaceAll(userId, userId[0:1], "")
		allPassanger, err := getPassanger(id)
		if err != nil {
			return "", err
		} else if len(allPassanger) != 1 {
			return "", nil
		} else {
			return userId, nil
		}
	} else if userType == "d" {
		id := strings.ReplaceAll(userId, userId[0:1], "")
		allDriver, err := getDriver(id)
		if err != nil {
			return "", err
		} else if len(allDriver) != 1 {
			return "", nil
		} else {
			return userId, nil
		}
	} else {
		return "", nil
	}
}

func menuPassanger() string {
	fmt.Println("========== Ride Share ==========")
	fmt.Println("Passanger Console")
	fmt.Println(" 1. Update information")
	fmt.Println(" 2. Display trips")
	fmt.Println(" 3. Start a new trip")
	fmt.Println(" 999. Log out")
	fmt.Print("Enter a user to sign in as (eg. c1): ")

	var option string
	fmt.Scanln(&option)
	return option
}

func menuDriver() string {
	fmt.Println("========== Ride Share ==========")
	fmt.Println("Driver Console")
	fmt.Println(" 1. Update information")
	fmt.Println(" 2. Change availability status (to get allocated trips)")
	fmt.Println(" \nTrip Status")
	fmt.Println(" 999. Log out")
	fmt.Print("Enter a user to sign in as (eg. c1): ")

	var option string
	fmt.Scanln(&option)
	return option
}
