package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
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

type Trip_Assignment_With_Passanger_Trip struct {
	Trip_Id      int            `json:"trip_id"`
	Driver_Id    sql.NullInt32  `json:"driver_id"`
	Status       sql.NullString `json:"status"`
	Passanger_Id int            `json:"passanger_id"`
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
	Passanger_Id int          `json:"passanger_id"`
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
	if req, err := http.NewRequest(http.MethodPut, fmt.Sprint("http://localhost:5000/api/v1/trip_assignment"), resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == http.StatusAccepted {
				return nil
			} else {
				return nil
			}
		} else {
			return err
		}
	} else {
		return err
	}
}

func getCurrentTripAssignmentWithMoreDataFilterPassangerId(passangerId string) ([]Trip_Assignment_With_Passanger_Trip, error) {
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/passanger/current_assignment/"+passangerId, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				var allTrip []Trip_Assignment_With_Passanger_Trip
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
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/driver/current_assignment/"+driverId, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
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
