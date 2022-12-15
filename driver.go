package main

import (
	"bytes"
	"encoding/json"
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

func getDriver(id string) ([]Driver, error) {
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/driver/"+id, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
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

func createDriver(driver Driver) error {
	client := &http.Client{}
	postBody, _ := json.Marshal(driver)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/driver", resBody); err == nil {
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

func updateDriver(driver Driver) error {
	client := &http.Client{}
	postBody, _ := json.Marshal(driver)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest(http.MethodPut, fmt.Sprint("http://localhost:5000/api/v1/driver/", driver.Driver_Id), resBody); err == nil {
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

func updateDriverAvilability(driver Driver) error {
	client := &http.Client{}
	postBody, _ := json.Marshal(driver)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest(http.MethodPut, fmt.Sprint("http://localhost:5000/api/v1/driver/is_available/", driver.Driver_Id), resBody); err == nil {
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
