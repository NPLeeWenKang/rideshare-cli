package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
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
				createUserDriver()
				return
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
				} else if option == "2" {
					displayPassangerTrips()
				} else if option == "3" {
					displayCreateTrip()
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

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 0, ' ', 0)
	fmt.Fprintln(w, "User Id", "\t", "User Type", "\t", "First Name", "\t", "Last Name")
	for _, v := range allPassanger {
		fmt.Fprintln(w, "p"+strconv.Itoa(v.Passanger_Id), "\t", "Passanger", "\t", v.First_Name, "\t", v.Last_Name)
	}

	// Get all drivers
	allDriver, err := getAllDriver()
	if err != nil {
		fmt.Println("Error occured while retrieving Passangers")
	}

	for _, v := range allDriver {
		fmt.Fprintln(w, "d"+strconv.Itoa(v.Driver_Id), "\t", "Driver", "\t", v.First_Name, "\t", v.Last_Name)
	}
	w.Flush()

	fmt.Println()
	fmt.Println("777. Create Passanger")
	fmt.Println("888. Create Driver")
	fmt.Println("999. Quit")

	fmt.Print("Enter an option (eg. 888) or a user id (eg. p1 or d1): ")

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
			fmt.Println("Trip successfully created")
		} else {
			fmt.Println("Error occured while creating trip")
		}
	}
}

func createUserDriver() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n========== Create User (Driver) ==========")
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

	fmt.Printf("Identification No: ")
	scanner.Scan()
	idNo := scanner.Text()
	if strings.ToLower(idNo) == "esc" {
		return
	}

	fmt.Printf("Car No: ")
	scanner.Scan()
	carNo := scanner.Text()
	if strings.ToLower(carNo) == "esc" {
		return
	}

	fmt.Print("Confirm Create? (y/n): ")
	scanner.Scan()
	confirmUpdate := scanner.Text()
	if strings.ToLower(confirmUpdate) == "esc" {
		return
	}

	if strings.ToLower(confirmUpdate) == "y" || strings.ToLower(confirmUpdate) == "yes" {
		err := createDriver(Driver{First_Name: firstName, Last_Name: lastName, Email: email, Mobile_No: mobileNo, Id_No: idNo, Car_No: carNo})
		if err == nil {
			fmt.Println("Driver successfully updated")
		} else {
			fmt.Println("Error occured while updating driver")
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

func displayPassangerTrips() {
	id := strings.ReplaceAll(userId, userId[0:1], "")
	trips, err := getTripFilterPassangerId(id)
	if err != nil {
		fmt.Println("Error occured while retrieving users")
		return
	} else if len(trips) == 0 {
		fmt.Println("No Trips")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 0, ' ', 0)
	fmt.Fprintln(w, "Trip Id", "\t", "Pickup", "\t", "Dropoff", "\t", "Start", "\t", "End", "\t", "Status")
	for _, v := range trips {
		fmt.Fprintln(w, v.Trip_Id, "\t", v.Pick_Up, "\t", v.Drop_Off, "\t", processSQLNullTime(v.Start), "\t", processSQLNullTime(v.End), "\t", processSQLNullString(v.Status))
	}
	w.Flush()
}

func displayCreateTrip() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n========== Create Trip ==========")
	fmt.Println(`Type "esc" for any option go back to menu`)

	fmt.Print("Pickup Location: ")
	scanner.Scan()
	pickUp := scanner.Text()
	if strings.ToLower(pickUp) == "esc" {
		return
	}

	fmt.Print("Dropoff Location: ")
	scanner.Scan()
	dropOff := scanner.Text()
	if strings.ToLower(dropOff) == "esc" {
		return
	}

	fmt.Print("Confirm Create? (y/n): ")
	scanner.Scan()
	confirmUpdate := scanner.Text()
	if strings.ToLower(confirmUpdate) == "esc" {
		return
	}

	if strings.ToLower(confirmUpdate) == "y" || strings.ToLower(confirmUpdate) == "yes" {
		id := strings.ReplaceAll(userId, userId[0:1], "")
		intId, _ := strconv.Atoi(id)
		err := createTrip(Trip{Passanger_Id: intId, Pick_Up: pickUp, Drop_Off: dropOff})
		if err == nil {
			fmt.Println("Passanger successfully updated")
		} else {
			fmt.Println("Error occured while updating passanger")
		}
	}
}

// ----------------------------------------------------------------------------
func menuDriver() string {
	id := strings.ReplaceAll(userId, userId[0:1], "")
	drivers, err := getDriver(id)
	if err != nil {
		fmt.Println("Error occured while retrieving users")
		return ""
	} else if len(drivers) != 1 {
		fmt.Println("Error occured: No user with ID")
		return ""
	}
	onlyDriver := drivers[0]

	fmt.Println("========== Ride Share ==========")
	fmt.Printf("Driver Id: %d\n", onlyDriver.Driver_Id)
	fmt.Printf("First Name: %s\n", onlyDriver.First_Name)
	fmt.Printf("Last Name: %s\n", onlyDriver.Last_Name)
	fmt.Printf("Email: %s\n", onlyDriver.Email)
	fmt.Printf("Mobile No: %s\n", onlyDriver.Mobile_No)
	fmt.Printf("Identification No: %s\n", onlyDriver.Id_No)
	fmt.Printf("Car No: %s\n\n", onlyDriver.Car_No)
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

func updateUserDriver() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n========== Create User (Driver) ==========")
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

	fmt.Printf("Identification No: ")
	scanner.Scan()
	idNo := scanner.Text()
	if strings.ToLower(idNo) == "esc" {
		return
	}

	fmt.Printf("Car No: ")
	scanner.Scan()
	carNo := scanner.Text()
	if strings.ToLower(carNo) == "esc" {
		return
	}

	fmt.Print("Confirm Create? (y/n): ")
	scanner.Scan()
	confirmUpdate := scanner.Text()
	if strings.ToLower(confirmUpdate) == "esc" {
		return
	}

	if strings.ToLower(confirmUpdate) == "y" || strings.ToLower(confirmUpdate) == "yes" {
		err := createDriver(Driver{First_Name: firstName, Last_Name: lastName, Email: email, Mobile_No: mobileNo, Id_No: idNo, Car_No: carNo})
		if err == nil {
			fmt.Println("Driver successfully updated")
		} else {
			fmt.Println("Error occured while updating driver")
		}
	}
}

func processSQLNullString(data sql.NullString) string {
	if data.Valid {
		return data.String
	} else {
		return "NULL"
	}
}

func processSQLNullTime(data sql.NullTime) string {
	if data.Valid {
		return fmt.Sprintf("%d/%d/%d %d:%d", data.Time.Day(), data.Time.Month(), data.Time.Year(), data.Time.Hour(), data.Time.Minute())
	} else {
		return "NULL"
	}
}
