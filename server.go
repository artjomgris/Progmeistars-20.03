package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"strconv"
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

			update, err := db.Query(qwery)
			if err != nil {
				panic(err.Error())
			}
			defer update.Close()
			w.Write([]byte("Updated successfully!"))
		}

	} else if req.Method == "PUT" {

		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			panic(err.Error())
		}

		var ins Data
		err = json.Unmarshal(data, &ins)
		if err != nil {
			panic(err.Error())
		}

		results, err := db.Query(fmt.Sprintf("SELECT `id` FROM maindata WHERE `id` = %v", ins.Id))
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		cntres := 0
		for results.Next() {
			err := results.Scan(&cntres)
			if err != nil {
				panic(err.Error())
			}
		}

		if cntres > 0 {
			w.Write([]byte("ID already exists!"))
		} else {
			if ins.Id == 0 {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO `maindata`(`name`, `lname`, `age`) VALUES ('%v', '%v', %v);", ins.Name, ins.Lname, ins.Age))
				if err != nil {
					panic(err.Error())
				}
				defer insert.Close()
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO maindata VALUES (%v, '%v', '%v', %v);", ins.Id, ins.Name, ins.Lname, ins.Age))
				if err != nil {
					panic(err.Error())
				}
				defer insert.Close()
			}
			w.Write([]byte("Inserted successfully!"))
		}

	} else if req.Method == "DELETE" {

		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			panic(err.Error())
		}

		ins, err := strconv.Atoi(string(data))
		if err != nil {
			panic(err.Error())
		}

		if ins == 0 {
			w.Write([]byte("Invalid Id"))
		} else {
			remove, err := db.Query(fmt.Sprintf("DELETE FROM `maindata` WHERE `id` = %v;", ins))
			if err != nil {
				panic(err.Error())
			}
			defer remove.Close()
			w.Write([]byte("Removed successfully!"))
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
