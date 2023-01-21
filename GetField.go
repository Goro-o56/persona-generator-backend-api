package main

import _ "github.com/go-sql-driver/mysql"

func GetRandomFieldFromRecord(record []string, random_index int) string {
	return record[random_index]
}
