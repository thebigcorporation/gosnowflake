package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/snowflakedb/gosnowflake"
)

func main() {
	// get environment variables
	env := func(k string) string {
		if value := os.Getenv(k); value != "" {
			return value
		}
		log.Fatalf("%v environment variable is not set.", k)
		return ""
	}

	account := env("SNOWFLAKE_TEST_ACCOUNT")
	user := env("SNOWFLAKE_TEST_USER")
	password := env("SNOWFLAKE_TEST_PASSWORD")

	dsn := fmt.Sprintf("%v:%v@%v", user, password, account)
	db, err := sql.Open("snowflake", dsn)
	if err != nil {
		log.Fatalf("failed to connect. %v, err: %v", dsn, err)
	}
	query := "SELECT 1"
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Fatalf("failed to run a query. %v, err: %v", query, err)
	}
	var v int
	for rows.Next() {
		err := rows.Scan(&v)
		if err != nil {
			log.Fatalf("failed to get result. err: %v", err)
		}
		if v != 1 {
			log.Fatalf("failed to get 1. got: %v", v)
		}
		fmt.Printf("Congrats! You have successfully run %v with Snowflake DB!", query)
	}
}
