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
		if len(userId) == 0 { // Main menu
			option := menu()
			if option == "999" { // Exit program
				return
			} else if option == "000" { // Refresh the data
				continue
			} else if option == "777" { // Display create passanger flow
				createUserPassenger()
				continue
			} else if option == "888" { // Display create driver flow
				createUserDriver()
				continue
			}
			
			// User logs in using the user id. Eg. p1, d3
			tempUserId, err := confirmUser(option) // Checks with the database to determine whether user id is valid
			if err != nil {
				fmt.Println("Error occured while retrieving users")
			} else if tempUserId == "" {
				fmt.Println("Error occured: No user with ID")
			}
			userId = tempUserId // Saves user id (with user type) to global variable
			fmt.Println("Sign In Successful")
		} else { // User specific menu
			userType := strings.ToLower(userId[0:1])
			if userType == "p" { // Passanger menu
				option := menuPassenger()
				if option == "1" { // Display update UI
					updateInformationPassenger()
				} else if option == "2" { // Display passanger trips ordered in descending order
					displayPassengerTrips()
				} else if option == "3" { // Create trip
					displayCreateTrip()
				} else if option == "000" { // Refresh data
					continue
				} else if option == "999" { // Logout and return to main menu
					userId = ""
				} else {
					fmt.Println("Invalid option")
				}
			} else if userType == "d" { / Driver menu
				option := menuDriver()
				if option == "1" { // Display update UI
					updateInformationDriver()
				} else if option == "2" { // Display update availability UI
					updateUserDriverAvailability()
				} else if option == "3" || option == "4" || option == "5" || option == "6" { // Changing status of trip assignment
					id := strings.ReplaceAll(userId, userId[0:1], "")
					tripAssignments, err := getCurrentTripAssignmentWithMoreDataFilterDriverId(id)
					if err != nil {
						fmt.Println("Error occured while retrieving users")
						continue
					}
					onlyTripAssignment := tripAssignments[0]
					
					// Determines which status to update to
					if option == "3" && onlyTripAssignment.Status == "PENDING" {
						updateTripAssignment(Trip_Assignment{Trip_Id: onlyTripAssignment.Trip_Id, Driver_Id: onlyTripAssignment.Driver_Id, Status: "ACCEPTED"})
					} else if option == "4" && onlyTripAssignment.Status == "PENDING" {
						updateTripAssignment(Trip_Assignment{Trip_Id: onlyTripAssignment.Trip_Id, Driver_Id: onlyTripAssignment.Driver_Id, Status: "REJECTED"})
					} else if option == "5" && onlyTripAssignment.Status == "ACCEPTED" {
						updateTripAssignment(Trip_Assignment{Trip_Id: onlyTripAssignment.Trip_Id, Driver_Id: onlyTripAssignment.Driver_Id, Status: "DRIVING"})
					} else if option == "6" && onlyTripAssignment.Status == "DRIVING" {
						updateTripAssignment(Trip_Assignment{Trip_Id: onlyTripAssignment.Trip_Id, Driver_Id: onlyTripAssignment.Driver_Id, Status: "DONE"})
					} else {
						fmt.Println("Invalid option")
					}
				} else if option == "000" {
					continue
				} else if option == "999" {
					userId = ""
				} else {
					fmt.Println("Invalid option")
				}
			}
		}
	}
}

// Main menu
func menu() string {

	fmt.Println("========== Ride Share ==========")
	fmt.Println("Passenger & Driver Console")

	// Get all passengers
	allPassenger, err := getAllPassenger()
	if err != nil {
		fmt.Println("Error occured while retrieving Passengers")
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 0, ' ', 0)
	fmt.Fprintln(w, "User Id", "\t", "User Type", "\t", "First Name", "\t", "Last Name")
	for _, v := range allPassenger {
		fmt.Fprintln(w, "p"+strconv.Itoa(v.Passenger_Id), "\t", "Passenger", "\t", v.First_Name, "\t", v.Last_Name)
	}

	// Get all drivers
	allDriver, err := getAllDriver()
	if err != nil {
		fmt.Println("Error occured while retrieving Drivers")
	}

	for _, v := range allDriver {
		fmt.Fprintln(w, "d"+strconv.Itoa(v.Driver_Id), "\t", "Driver", "\t", v.First_Name, "\t", v.Last_Name)
	}
	w.Flush()

	fmt.Println()
	fmt.Println("000. Refresh")
	fmt.Println("777. Create Passenger")
	fmt.Println("888. Create Driver")
	fmt.Println("999. Quit")

	fmt.Print("Enter an option (eg. 888) or a user id (eg. p1 or d1): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	option := scanner.Text()
	return option
}

// Passanger creation menu
func createUserPassenger() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n========== Create User (Passenger) ==========")
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
		err := createPassenger(Passenger{First_Name: firstName, Last_Name: lastName, Email: email, Mobile_No: mobileNo})
		if err == nil {
			fmt.Println("Passenger successfully created")
		} else {
			fmt.Println("Error occured while creating passenger")
		}
	}
}

// Driver creation menu
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
			fmt.Println("Driver successfully created")
		} else {
			fmt.Println("Error occured while creating driver")
		}
	}
}

// Utility function confirm whether user input id is valid
func confirmUser(userId string) (string, error) {
	userType := strings.ToLower(userId[0:1])
	if userType == "p" { // Checks passanger table
		id := strings.ReplaceAll(userId, userId[0:1], "")
		allPassenger, err := getPassenger(id)
		if err != nil {
			return "", err
		} else if len(allPassenger) != 1 {
			return "", nil
		} else {
			return userId, nil
		}
	} else if userType == "d" { // Checks driver table
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

// Passanger menu
func menuPassenger() string {
	id := strings.ReplaceAll(userId, userId[0:1], "")
	passengers, err := getPassenger(id)
	if err != nil {
		fmt.Println("Error occured while retrieving users")
		return ""
	} else if len(passengers) != 1 {
		fmt.Println("Error occured: No user with ID")
		return ""
	}
	onlyPassenger := passengers[0]

	fmt.Println("========== Ride Share ==========")
	fmt.Printf("Passenger Id: %d\n", onlyPassenger.Passenger_Id)
	fmt.Printf("First Name: %s\n", onlyPassenger.First_Name)
	fmt.Printf("Last Name: %s\n", onlyPassenger.Last_Name)
	fmt.Printf("Email: %s\n", onlyPassenger.Email)
	fmt.Printf("Mobile No: %s\n\n", onlyPassenger.Mobile_No)
	fmt.Println("Passenger Console")
	fmt.Println(" 1. Update information")
	fmt.Println(" 2. Display trips")
	fmt.Println(" 3. Start a new trip")

	tripAssignments, err := getCurrentTripAssignmentWithMoreDataFilterPassengerId(id)
	if err != nil {
		fmt.Println("Error occured while retrieving users")
		return ""
	} else if len(tripAssignments) != 0 {
		fmt.Println(" \n========== Trip Status ==========")
	}
	for _, v := range tripAssignments {
		fmt.Printf("\nTrip Id: %d\n", v.Trip_Id)
		if !v.Driver_Id.Valid { // Trip has not been assigned before
			fmt.Printf("Pickup Location: %s\n", v.Pick_Up)
			fmt.Printf("Dropoff Location: %s\n", v.Drop_Off)
			fmt.Printf("Start Time: %s\n", processSQLNullTime(v.Start))
			fmt.Printf("End Time: %s\n", processSQLNullTime(v.End))
			fmt.Printf("Status: ASSIGNING...\n")
		} else {
			if processSQLNullString(v.Status) == "REJECTED" { // Trip has been assigned but rejected
				fmt.Printf("Pickup Location: %s\n", v.Pick_Up)
				fmt.Printf("Dropoff Location: %s\n", v.Drop_Off)
				fmt.Printf("Start Time: %s\n", processSQLNullTime(v.Start))
				fmt.Printf("End Time: %s\n", processSQLNullTime(v.End))
				fmt.Printf("Status: ASSIGNING...\n")
			} else { // Trip has been successfully assigned driver
				fmt.Printf("Driver: (%s) %s %s\n", processSQLNullInt(v.Driver_Id), processSQLNullString(v.First_Name), processSQLNullString(v.Last_Name))
				fmt.Printf("Mobile No: %s\n", processSQLNullString(v.Mobile_No))
				fmt.Printf("Car No: %s\n", processSQLNullString(v.Car_No))
				fmt.Printf("Pickup Location: %s\n", v.Pick_Up)
				fmt.Printf("Dropoff Location: %s\n", v.Drop_Off)
				fmt.Printf("Start Time: %s\n", processSQLNullTime(v.Start))
				fmt.Printf("End Time: %s\n", processSQLNullTime(v.End))
				fmt.Printf("Status: %s\n", processSQLNullString(v.Status))
			}
		}
	}

	fmt.Println("\n 000. Refresh Data")
	fmt.Println(" 999. Log out")
	fmt.Print("Enter an Option: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	option := scanner.Text()
	return option
}

func updateInformationPassenger() {
	id := strings.ReplaceAll(userId, userId[0:1], "")
	passengers, err := getPassenger(id)
	if err != nil {
		fmt.Println("Error occured while retrieving users")
		return
	} else if len(passengers) != 1 {
		fmt.Println("Error occured: No user with ID")
		return
	}
	onlyPassenger := passengers[0]

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("========== Update Information (Passenger) ==========")
	fmt.Println(`Type "esc" for any option go back to menu`)
	fmt.Printf("Passenger Id: %d\n", onlyPassenger.Passenger_Id)

	fmt.Printf("First Name (%s): ", onlyPassenger.First_Name)
	scanner.Scan()
	firstName := scanner.Text()
	if strings.ToLower(firstName) == "esc" {
		return
	}

	fmt.Printf("Last Name (%s): ", onlyPassenger.Last_Name)
	scanner.Scan()
	lastName := scanner.Text()
	if strings.ToLower(lastName) == "esc" {
		return
	}

	fmt.Printf("Email (%s): ", onlyPassenger.Email)
	scanner.Scan()
	email := scanner.Text()
	if strings.ToLower(email) == "esc" {
		return
	}

	fmt.Printf("Mobile No (%s): ", onlyPassenger.Mobile_No)
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
		err := updatePassenger(Passenger{Passenger_Id: onlyPassenger.Passenger_Id, First_Name: firstName, Last_Name: lastName, Email: email, Mobile_No: mobileNo})
		if err == nil {
			fmt.Println("Passenger successfully updated")
		} else {
			fmt.Println("Error occured while updating passenger")
		}
	}
}

func displayPassengerTrips() {
	id := strings.ReplaceAll(userId, userId[0:1], "")
	trips, err := getTripFilterPassengerId(id)
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
		err := createTrip(Trip{Passenger_Id: intId, Pick_Up: pickUp, Drop_Off: dropOff})
		if err == nil {
			fmt.Println("Trip successfully updated")
		} else {
			fmt.Println("Error occured while updating trip")
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
	fmt.Printf("\nDriver Id: %d\n", onlyDriver.Driver_Id)
	fmt.Printf("First Name: %s\n", onlyDriver.First_Name)
	fmt.Printf("Last Name: %s\n", onlyDriver.Last_Name)
	fmt.Printf("Email: %s\n", onlyDriver.Email)
	fmt.Printf("Mobile No: %s\n", onlyDriver.Mobile_No)
	fmt.Printf("Identification No: %s\n", onlyDriver.Id_No)
	fmt.Printf("Car No: %s\n", onlyDriver.Car_No)
	fmt.Printf("Availability: %t\n\n", onlyDriver.Is_Available)
	fmt.Println("Driver Console")
	fmt.Println(" 1. Update information")
	fmt.Println(" 2. Change availability status (to get allocated trips)")

	tripAssignments, err := getCurrentTripAssignmentWithMoreDataFilterDriverId(id)
	if err != nil {
		fmt.Println("Error occured while retrieving users")
		return ""
	}
	if len(tripAssignments) == 1 {
		onlyTripAssignment := tripAssignments[0]

		fmt.Println(" \n========== Trip Status ==========")
		fmt.Printf("Trip Id: %d\n", onlyTripAssignment.Trip_Id)
		fmt.Printf("Passenger: (%d) %s %s\n", onlyTripAssignment.Passenger_Id, onlyTripAssignment.First_Name, onlyTripAssignment.Last_Name)
		fmt.Printf("Mobile No: %s\n", onlyTripAssignment.Mobile_No)
		fmt.Printf("Pickup Location: %s\n", onlyTripAssignment.Pick_Up)
		fmt.Printf("Dropoff Location: %s\n", onlyTripAssignment.Drop_Off)
		fmt.Printf("Start Time: %s\n", processSQLNullTime(onlyTripAssignment.Start))
		fmt.Printf("End Time: %s\n", processSQLNullTime(onlyTripAssignment.End))
		fmt.Printf("Status: %s\n", onlyTripAssignment.Status)
		fmt.Println("\nTrip Console")
		if onlyTripAssignment.Status == "PENDING" {
			fmt.Println(" 3. Accept Trip")
			fmt.Println(" 4. Reject Trip")
		} else if onlyTripAssignment.Status == "ACCEPTED" {
			fmt.Println(" 5. Start Trip")
		} else if onlyTripAssignment.Status == "DRIVING" {
			fmt.Println(" 6. End Trip")
		}
	}

	fmt.Println("\n 000. Refresh Data")
	fmt.Println(" 999. Log out")
	fmt.Print("Enter an Option: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	option := scanner.Text()
	return option
}

func updateInformationDriver() {
	scanner := bufio.NewScanner(os.Stdin)

	id := strings.ReplaceAll(userId, userId[0:1], "")
	drivers, err := getDriver(id)
	if err != nil {
		fmt.Println("Error occured while retrieving users")
		return
	} else if len(drivers) != 1 {
		fmt.Println("Error occured: No user with ID")
		return
	}
	onlyDriver := drivers[0]

	fmt.Println("\n========== Update User (Driver) ==========")
	fmt.Println(`Type "esc" for any option go back to menu`)

	fmt.Printf("First Name (%s): ", onlyDriver.First_Name)
	scanner.Scan()
	firstName := scanner.Text()
	if strings.ToLower(firstName) == "esc" {
		return
	}

	fmt.Printf("Last Name (%s): ", onlyDriver.Last_Name)
	scanner.Scan()
	lastName := scanner.Text()
	if strings.ToLower(lastName) == "esc" {
		return
	}

	fmt.Printf("Email (%s): ", onlyDriver.Email)
	scanner.Scan()
	email := scanner.Text()
	if strings.ToLower(email) == "esc" {
		return
	}

	fmt.Printf("Mobile No (%s): ", onlyDriver.Mobile_No)
	scanner.Scan()
	mobileNo := scanner.Text()
	if strings.ToLower(mobileNo) == "esc" {
		return
	}

	fmt.Printf("Car No (%s): ", onlyDriver.Car_No)
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
		err := updateDriver(Driver{Driver_Id: onlyDriver.Driver_Id, First_Name: firstName, Last_Name: lastName, Email: email, Mobile_No: mobileNo, Car_No: carNo})
		if err == nil {
			fmt.Println("Driver successfully updated")
		} else {
			fmt.Println("Error occured while updating driver")
		}
	}
}

func updateUserDriverAvailability() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n========== Update Availability (Driver) ==========")
	fmt.Println(`Type "esc" for any option go back to menu`)

	var isAvailableBool bool
	for {
		fmt.Print("Availability (y/n): ")
		scanner.Scan()
		isAvailable := scanner.Text()
		if strings.ToLower(isAvailable) == "esc" {
			return
		} else if strings.ToLower(isAvailable) == "y" || strings.ToLower(isAvailable) == "yes" {
			isAvailableBool = true
			break
		} else if strings.ToLower(isAvailable) == "n" || strings.ToLower(isAvailable) == "no" {
			isAvailableBool = false
			break
		} else {
			fmt.Println("Wrong input, try y/n/esc")
		}
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
		err := updateDriverAvilability(Driver{Driver_Id: intId, Is_Available: isAvailableBool})
		if err == nil {
			fmt.Println("Driver successfully updated")
		} else {
			fmt.Println("Error occured while updating driver")
		}
	}
}

func processSQLNullInt(data sql.NullInt32) string {
	if data.Valid {
		return strconv.Itoa(int(data.Int32))
	} else {
		return "-"
	}
}

func processSQLNullString(data sql.NullString) string {
	if data.Valid {
		return data.String
	} else {
		return "-"
	}
}

func processSQLNullTime(data sql.NullTime) string {
	if data.Valid {
		return fmt.Sprintf("%d/%d/%d %d:%d UTC", data.Time.Day(), data.Time.Month(), data.Time.Year(), data.Time.Hour(), data.Time.Minute())
	} else {
		return "-"
	}
}

func validateTripConsoleOptions(onlyTripAssignment Trip_Assignment_With_Driver_Trip, option string) bool {
	if onlyTripAssignment.Status == "PENDING" && (option == "3" || option == "4") {
		return true
	} else if onlyTripAssignment.Status == "ACCEPTED" && (option == "5") {
		return true
	} else if onlyTripAssignment.Status == "DRIVING" && (option == "6") {
		return true
	} else {
		return false
	}
}
