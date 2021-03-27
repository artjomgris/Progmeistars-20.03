package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Data struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Lname string `json:"lastname"`
	Age   int    `json:"age"`
}

func main() {
	for i := 1; i <= 11; i++ {
		fmt.Println(removeData(i))
	}

}

func getData() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000", nil)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	rec, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(rec)
}

func changeData(data Data) string {
	client := &http.Client{}

	var body bytes.Buffer

	out, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	body.Write(out)
	req, err := http.NewRequest("POST", "http://127.0.0.1:5000", &body)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	rec, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	return string(rec)
}

func createData(data Data) string {
	client := &http.Client{}

	var body bytes.Buffer

	out, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	body.Write(out)
	req, err := http.NewRequest("PUT", "http://127.0.0.1:5000", &body)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	rec, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	return string(rec)
}

func removeData(id int) string {
	client := &http.Client{}

	var body bytes.Buffer

	body.Write([]byte(strconv.Itoa(id)))
	req, err := http.NewRequest("DELETE", "http://127.0.0.1:5000", &body)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	rec, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	return string(rec)
}
