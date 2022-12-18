package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Driver struct {
	Driver_Id    int    `json:"driver_id"`
	First_Name   string `json:"first_name"`
	Last_Name    string `json:"last_name"`
	Mobile_No    string `json:"mobile_no"`
	Email        string `json:"email"`
	Id_No        string `json:"id_no"`
	Car_No       string `json:"car_no"`
	Is_Available bool   `json:"is_available"`
}

func getAllDriver() ([]Driver, error) {
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/driver", nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				if res.StatusCode == http.StatusBadRequest {
					err = errors.New("ERROR: Bad Request")
					return nil, err
				}
				var allDriver []Driver
				json.Unmarshal(body, &allDriver)
				return allDriver, nil
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

// Gets drivers based on a single driver id. This returns a array as it is easier to deal with empty array instead of null data. To check whether the the query works use len().
func getDriver(id string) ([]Driver, error) {
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/driver/"+id, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				if res.StatusCode == http.StatusBadRequest {
					err = errors.New("ERROR: Bad Request")
					return nil, err
				}
				var allDriver []Driver
				json.Unmarshal(body, &allDriver)
				return allDriver, nil
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

// Creates a driver based on the driver object provided. The driver id is auto assigned on the database.
func createDriver(driver Driver) error {
	client := &http.Client{}
	postBody, _ := json.Marshal(driver)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/driver", resBody); err == nil {
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

// Updates driver information based on the object provided. Although indetification number can be included in the object, this attribute does not get updated in the API.
func updateDriver(driver Driver) error {
	client := &http.Client{}
	postBody, _ := json.Marshal(driver)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest(http.MethodPut, fmt.Sprint("http://localhost:5000/api/v1/driver/", driver.Driver_Id), resBody); err == nil {
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

// Updates driver availability, it uses the driver object to pass the data to the API but only the is_available attribute is used.
func updateDriverAvilability(driver Driver) error {
	client := &http.Client{}
	postBody, _ := json.Marshal(driver)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest(http.MethodPut, fmt.Sprint("http://localhost:5000/api/v1/driver/is_available/", driver.Driver_Id), resBody); err == nil {
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
