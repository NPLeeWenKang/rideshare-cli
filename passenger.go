package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Passenger struct {
	Passenger_Id int    `json:"passenger_id"`
	First_Name   string `json:"first_name"`
	Last_Name    string `json:"last_name"`
	Mobile_No    string `json:"mobile_no"`
	Email        string `json:"email"`
}

func getAllPassenger() ([]Passenger, error) {
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/passenger", nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				if res.StatusCode == http.StatusBadRequest {
					err = errors.New("ERROR: Bad Request")
					return nil, err
				}
				var allPassenger []Passenger
				json.Unmarshal(body, &allPassenger)
				return allPassenger, nil
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

func getPassenger(id string) ([]Passenger, error) {
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/passenger/"+id, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				if res.StatusCode == http.StatusBadRequest {
					err = errors.New("ERROR: Bad Request")
					return nil, err
				}
				var allPassenger []Passenger
				json.Unmarshal(body, &allPassenger)
				return allPassenger, nil
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

func createPassenger(passenger Passenger) error {
	client := &http.Client{}
	postBody, _ := json.Marshal(passenger)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/passenger", resBody); err == nil {
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

func updatePassenger(passenger Passenger) error {
	client := &http.Client{}
	postBody, _ := json.Marshal(passenger)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest(http.MethodPut, fmt.Sprint("http://localhost:5000/api/v1/passenger/", passenger.Passenger_Id), resBody); err == nil {
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
