package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Trip_Assignment struct {
	Trip_Id         int       `json:"trip_id"`
	Driver_Id       int       `json:"driver_id"`
	Status          string    `json:"status"`
	Assign_Datetime time.Time `json:"assign_datetime"`
}

type Trip_Assignment_With_Passenger_Trip struct {
	Trip_Id      int            `json:"trip_id"`
	Driver_Id    sql.NullInt32  `json:"driver_id"`
	Status       sql.NullString `json:"status"`
	Passenger_Id int            `json:"passenger_id"`
	First_Name   sql.NullString `json:"first_name"`
	Last_Name    sql.NullString `json:"last_name"`
	Mobile_No    sql.NullString `json:"mobile_no"`
	Email        sql.NullString `json:"email"`
	Pick_Up      string         `json:"pick_up"`
	Drop_Off     string         `json:"drop_off"`
	Start        sql.NullTime   `json:"start"`
	End          sql.NullTime   `json:"end"`
	Car_No       sql.NullString `json:"car_no"`
}

type Trip_Assignment_With_Driver_Trip struct {
	Trip_Id      int          `json:"trip_id"`
	Driver_Id    int          `json:"driver_id"`
	Status       string       `json:"status"`
	Passenger_Id int          `json:"passenger_id"`
	First_Name   string       `json:"first_name"`
	Last_Name    string       `json:"last_name"`
	Mobile_No    string       `json:"mobile_no"`
	Email        string       `json:"email"`
	Pick_Up      string       `json:"pick_up"`
	Drop_Off     string       `json:"drop_off"`
	Start        sql.NullTime `json:"start"`
	End          sql.NullTime `json:"end"`
}

func updateTripAssignment(tripAssignment Trip_Assignment) error {
	client := &http.Client{}
	postBody, _ := json.Marshal(tripAssignment)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest(http.MethodPut, fmt.Sprint("http://localhost:5001/api/v1/trip_assignment"), resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == http.StatusAccepted {
				return nil
			} else {
				err = errors.New("ERROR: Bad Request")
				return err
			}
		} else {
			return err
		}
	} else {
		return err
	}
}

func getCurrentTripAssignmentWithMoreDataFilterPassengerId(passengerId string) ([]Trip_Assignment_With_Passenger_Trip, error) {
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5001/api/v1/current_trip_assignment/passenger/"+passengerId, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				if res.StatusCode == http.StatusBadRequest {
					err = errors.New("ERROR: Bad Request")
					return nil, err
				}
				var allTrip []Trip_Assignment_With_Passenger_Trip
				json.Unmarshal(body, &allTrip)
				return allTrip, nil
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func getCurrentTripAssignmentWithMoreDataFilterDriverId(driverId string) ([]Trip_Assignment_With_Driver_Trip, error) {
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5001/api/v1/current_trip_assignment/driver/"+driverId, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				if res.StatusCode == http.StatusBadRequest {
					err = errors.New("ERROR: Bad Request")
					return nil, err
				}
				var allTrip []Trip_Assignment_With_Driver_Trip
				json.Unmarshal(body, &allTrip)
				return allTrip, nil
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
