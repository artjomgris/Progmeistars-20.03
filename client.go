package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Lname string `json:"lastname"`
	Age   int    `json:"age"`
}

func main() {
	fmt.Println(getData())
	/*person := Data{
		Id:  2,
		Age: 32,
	}
	fmt.Println(changeData(person))*/

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
