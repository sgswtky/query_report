package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var header []string

func main() {
	var (
		dbUser = flag.String("user", "", "database user.")
		dbPass = flag.String("pass", "", "database password.")
		dbHost = flag.String("host", "", "database host.")
		dbName = flag.String("db", "", "database schema.")
		query = flag.String("query", "", "exec query.")
		resultKey = flag.String("key", "", "result key.")
		resultValue = flag.String("value", "", "result value.")
		interval = flag.Int("interval", 10, "interval seconds.")
	)
	flag.Parse()

	if dbUser == nil || dbPass == nil || dbHost == nil || dbName == nil || query == nil || resultKey == nil || resultValue == nil || interval == nil{
		fmt.Println("Insufficient parameters")
		os.Exit(2)
	}

	// db open
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", *dbUser, *dbPass, *dbHost, *dbName))
	if err != nil {
		panic(err.Error())
	}

	// exec loop
	for {
		m, h := exec(db, *query, *resultKey, *resultValue)

		if len(header) == 0 {
			header = append([]string{"datetime"}, h...)
			fmt.Println(strings.Join(header, ","))
		}

		dateformat := "2006/01/02 15:04:05"
		nowTime := time.Now().Format(dateformat)

		commaSepareted := []string{nowTime}
		for _, head := range header {
			if head == "datetime" {
				continue
			}
			commaSepareted = append(commaSepareted, m[head])
		}

		fmt.Println(strings.Join(commaSepareted, ","))

		time.Sleep(time.Duration(*interval) * time.Second)
	}
}

func exec(db *sql.DB, query string, keyName string, valName string) (map[string]string, []string) {

	// exec query
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	m := map[string]string{}
	var h []string

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		var valueOfKey string
		var valueOfValue string

		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			// append
			switch columns[i] {
			case keyName:
				valueOfKey = value
				h = append(h, value)
			case valName:
				valueOfValue = value
			}
		}
		m[valueOfKey] = valueOfValue
	}
	return m, h
}
