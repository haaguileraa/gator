package main

import (
	"errors"
	"github.com/lib/pq"
	"log"
)

const duplicateObject = "23505"

func logSqlError(err error) {
	var pqError *pq.Error
	if errors.As(err, &pqError) {
		if pqError.Code == duplicateObject {
			return
		}
	}
	log.Println("error adding post to the database:", err)
}
