package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB

type Data struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Lname string `json:"lastname"`
	Age   int    `json:"age"`
}

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		results, err := db.Query("SELECT * FROM maindata")
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output []Data
		for results.Next() {
			var res Data
			err = results.Scan(&res.Id, &res.Name, &res.Lname, &res.Age)
			if err != nil {
				panic(err.Error())
			}
			output = append(output, res)
		}
		out, err := json.Marshal(output)
		if err != nil {
			panic(err.Error())
		}
		w.Write(out)

	} else if req.Method == "POST" {
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%s\n", data)

		var upd Data
		err = json.Unmarshal(data, &upd)
		if err != nil {
			panic(err.Error())
		}

		if upd.Id == 0 {
			w.Write([]byte("Invalid id"))
		} else {
			cnt := 0
			qwery := "UPDATE `maindata` SET "
			if upd.Name != "" {
				qwery += fmt.Sprintf("`name` = '%v'", upd.Name)
				cnt++
			}
			if upd.Age != 0 {
				if cnt != 0 {
					qwery += fmt.Sprintf(", `age` = %v", upd.Age)
					cnt++
				} else {
					qwery += fmt.Sprintf("`age` = %v", upd.Age)
				}

			}
			if upd.Lname != "" {
				if cnt != 0 {
					qwery += fmt.Sprintf(", `lname` = '%v'", upd.Lname)
				} else {
					qwery += fmt.Sprintf("`lname` = '%v'", upd.Lname)
				}

			}

			qwery += fmt.Sprintf(" WHERE `id` = %v;", upd.Id)
			fmt.Println(qwery)

			update, err := db.Query(qwery)
			if err != nil {
				panic(err.Error())
			}
			defer update.Close()
		}

	} else {
		w.WriteHeader(405)
	}

}

func main() {
	d, err := sql.Open("mysql", "goserver:gotest123@tcp(127.0.0.1)/gohttp")
	if err != nil {
		panic(err.Error())
	}
	db = d
	defer db.Close()
	http.HandleFunc("/", Handler)

	err = http.ListenAndServe("127.0.0.1:5000", nil)
	if err != nil {
		panic(err.Error())
	}
}
