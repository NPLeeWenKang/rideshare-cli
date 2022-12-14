package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Trip struct {
	Trip_Id      int          `json:"trip_id"`
	Passanger_Id int          `json:"passanger_id"`
	Pick_Up      string       `json:"pick_up"`
	Drop_Off     string       `json:"drop_off"`
	Start        sql.NullTime `json:"start"`
	End          sql.NullTime `json:"end"`
}

type Trip_Filter_Passanger struct {
	Trip_Id      int            `json:"trip_id"`
	Passanger_Id int            `json:"passanger_id"`
	Pick_Up      string         `json:"pick_up"`
	Drop_Off     string         `json:"drop_off"`
	Start        sql.NullTime   `json:"start"`
	End          sql.NullTime   `json:"end"`
	Status       sql.NullString `json:"status"`
}

type Trip_Assignment struct {
	Trip_Id         int       `json:"trip_id"`
	Driver_Id       int       `json:"driver_id"`
	Status          string    `json:"status"`
	Assign_Datetime time.Time `json:"assign_datetime"`
}

func getTripFilterPassangerId(passangerId string) ([]Trip_Filter_Passanger, error) {
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/trip?passanger_id="+passangerId, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				var allTrip []Trip_Filter_Passanger
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

func createTrip(trip Trip) error {
	client := &http.Client{}
	postBody, _ := json.Marshal(trip)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/trip", resBody); err == nil {
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
