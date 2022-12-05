package main

import (
	"bufio"
	"fmt"
	"os"
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
			} else if option == "777" {
				createUserPassanger()
				return
			} else if option == "888" {

			}
			tempUserId, err := confirmUser(option)
			if err != nil {
				fmt.Println("Error occured while retrieving users")
			} else if tempUserId == "" {
				fmt.Println("Error occured: No user with ID")
			}
			userId = tempUserId
			fmt.Println("Sign In Successful")
		} else {
			userType := strings.ToLower(userId[0:1])
			if userType == "p" {
				option := menuPassanger()
				if option == "1" {
					updateInformationPassanger()
				} else if option == "999" {
					userId = ""
				} else {
					fmt.Println("Invalid option")
				}
			} else if userType == "d" {
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
	fmt.Println(" 777. Create Passanger")
	fmt.Println(" 888. Create Driver")
	fmt.Println(" 999. Quit")
	fmt.Print("Enter a user to sign in as (eg. c1): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	option := scanner.Text()
	return option
}

func createUserPassanger() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n========== Create User (Passanger) ==========")
	fmt.Println(`Type "esc" for any option go back to menu`)

	fmt.Print("First Name: ")
	scanner.Scan()
	firstName := scanner.Text()
	if strings.ToLower(firstName) == "esc" {
		return
	}

	fmt.Print("Last Name: ")
	scanner.Scan()
	lastName := scanner.Text()
	if strings.ToLower(lastName) == "esc" {
		return
	}

	fmt.Printf("Email: ")
	scanner.Scan()
	email := scanner.Text()
	if strings.ToLower(email) == "esc" {
		return
	}

	fmt.Printf("Mobile No: ")
	scanner.Scan()
	mobileNo := scanner.Text()
	if strings.ToLower(mobileNo) == "esc" {
		return
	}

	fmt.Print("Confirm Create? (y/n): ")
	scanner.Scan()
	confirmUpdate := scanner.Text()
	if strings.ToLower(confirmUpdate) == "esc" {
		return
	}

	if strings.ToLower(confirmUpdate) == "y" || strings.ToLower(confirmUpdate) == "yes" {
		err := createPassanger(Passanger{First_Name: firstName, Last_Name: lastName, Email: email, Mobile_No: mobileNo})
		if err == nil {
			fmt.Println("Passanger successfully updated")
		} else {
			fmt.Println("Error occured while updating passanger")
		}
	}
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
	id := strings.ReplaceAll(userId, userId[0:1], "")
	passangers, err := getPassanger(id)
	if err != nil {
		fmt.Println("Error occured while retrieving users")
		return ""
	} else if len(passangers) != 1 {
		fmt.Println("Error occured: No user with ID")
		return ""
	}
	onlyPassanger := passangers[0]

	fmt.Println("========== Ride Share ==========")
	fmt.Printf("Passanger Id: %d\n", onlyPassanger.Passanger_Id)
	fmt.Printf("First Name: %s\n", onlyPassanger.First_Name)
	fmt.Printf("Last Name: %s\n", onlyPassanger.Last_Name)
	fmt.Printf("Email: %s\n", onlyPassanger.Email)
	fmt.Printf("Mobile No: %s\n\n", onlyPassanger.Mobile_No)
	fmt.Println("Passanger Console")
	fmt.Println(" 1. Update information")
	fmt.Println(" 2. Display trips")
	fmt.Println(" 3. Start a new trip")
	fmt.Println(" 999. Log out")
	fmt.Print("Enter a user to sign in as (eg. c1): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	option := scanner.Text()
	return option
}

func updateInformationPassanger() {
	id := strings.ReplaceAll(userId, userId[0:1], "")
	passangers, err := getPassanger(id)
	if err != nil {
		fmt.Println("Error occured while retrieving users")
		return
	} else if len(passangers) != 1 {
		fmt.Println("Error occured: No user with ID")
		return
	}
	onlyPassanger := passangers[0]

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("========== Update Information (Passanger) ==========")
	fmt.Println(`Type "esc" for any option go back to menu`)
	fmt.Printf("Passanger Id: %d\n", onlyPassanger.Passanger_Id)

	fmt.Printf("First Name (%s): ", onlyPassanger.First_Name)
	scanner.Scan()
	firstName := scanner.Text()
	if strings.ToLower(firstName) == "esc" {
		return
	}

	fmt.Printf("Last Name (%s): ", onlyPassanger.Last_Name)
	scanner.Scan()
	lastName := scanner.Text()
	if strings.ToLower(lastName) == "esc" {
		return
	}

	fmt.Printf("Email (%s): ", onlyPassanger.Email)
	scanner.Scan()
	email := scanner.Text()
	if strings.ToLower(email) == "esc" {
		return
	}

	fmt.Printf("Mobile No (%s): ", onlyPassanger.Mobile_No)
	scanner.Scan()
	mobileNo := scanner.Text()
	if strings.ToLower(mobileNo) == "esc" {
		return
	}

	fmt.Print("Confirm Update? (y/n): ")
	scanner.Scan()
	confirmUpdate := scanner.Text()
	if strings.ToLower(confirmUpdate) == "esc" {
		return
	}

	if strings.ToLower(confirmUpdate) == "y" || strings.ToLower(confirmUpdate) == "yes" {
		err := updatePassanger(Passanger{Passanger_Id: onlyPassanger.Passanger_Id, First_Name: firstName, Last_Name: lastName, Email: email, Mobile_No: mobileNo})
		if err == nil {
			fmt.Println("Passanger successfully updated")
		} else {
			fmt.Println("Error occured while updating passanger")
		}
	}
}

func menuDriver() string {
	fmt.Println("========== Ride Share ==========")
	fmt.Println("Driver Console")
	fmt.Println(" 1. Update information")
	fmt.Println(" 2. Change availability status (to get allocated trips)")
	fmt.Println(" \nTrip Status")
	fmt.Println(" 999. Log out")
	fmt.Print("Enter a user to sign in as (eg. c1): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	option := scanner.Text()
	return option
}
