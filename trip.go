package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Trip struct {
	Trip_Id      int          `json:"trip_id"`
	Passenger_Id int          `json:"passenger_id"`
	Pick_Up      string       `json:"pick_up"`
	Drop_Off     string       `json:"drop_off"`
	Start        sql.NullTime `json:"start"`
	End          sql.NullTime `json:"end"`
}

type Trip_Filter_Passenger struct {
	Trip_Id      int            `json:"trip_id"`
	Passenger_Id int            `json:"passenger_id"`
	Pick_Up      string         `json:"pick_up"`
	Drop_Off     string         `json:"drop_off"`
	Start        sql.NullTime   `json:"start"`
	End          sql.NullTime   `json:"end"`
	Status       sql.NullString `json:"status"`
}

// Gets trips by passengers. It is ordered by trip id in descending order.
func getTripFilterPassengerId(passengerId string) ([]Trip_Filter_Passenger, error) {
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5001/api/v1/trip?passenger_id="+passengerId, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				if res.StatusCode == http.StatusBadRequest {
					err = errors.New("ERROR: Bad Request")
					return nil, err
				}
				var allTrip []Trip_Filter_Passenger
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

// Creates trip based on the object provided.
func createTrip(trip Trip) error {
	client := &http.Client{}
	postBody, _ := json.Marshal(trip)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5001/api/v1/trip", resBody); err == nil {
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
