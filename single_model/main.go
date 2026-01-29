package main

import (
	"fmt"
	"sync"
)

type Db struct {
}

var db *Db
var db22 *Db

func initDb(dsn string) *Db {
	return &Db{}
}
func GetDb() *Db {
	if db == nil {
		db = initDb("127.0.0.1:3306")
	}
	return db
}

var one sync.Once

func GetDb2() *Db {
	one.Do(func() {
		db22 = initDb("127.0.0.1:3306")
	})
	return db22
}
func main() {
	db1 := GetDb()
	db2 := GetDb()
	fmt.Printf("ptr %p\n", db1)
	fmt.Printf("ptr %p \n", db2)
	db3 := GetDb2()
	db4 := GetDb2()
	fmt.Printf("ptr3 %p\n", db3)
	fmt.Printf("ptr4 %p\n", db4)
}
