package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Reading custom errors

type Error struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

func (e Error) Error() string {
	fs := "HTTP: %d, Code: %d, Message: %s"
	return fmt.Sprintf(fs, e.HTTPCode, e.Code, e.Message)
}

func get(u string) (*http.Response, error) {
	res, err := http.Get(u)
	if err != nil {
		return res, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		if res.Header.Get("Content-Type") != "application/json" {
			sm := "Unknown error. HTTP status: %s"
			return res, fmt.Errorf(sm, res.Status)
		}
		b, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		var data struct {
			Err Error `json:"error"`
		}
		fmt.Println("b", string(b))
		err = json.Unmarshal(b, &data)
		if err != nil {
			sm := "Unable to parse json: %s. HTTP status: %s"
			return res, fmt.Errorf(sm, err, res.Status)
		}
		data.Err.HTTPCode = res.StatusCode
		return res, data.Err
	}
	return res, nil
}

func main() {
	res, err := get("http://localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println(b)
}
