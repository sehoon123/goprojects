package main

import "fmt"

type Database interface {
	Get()
	Set()
}

type CDataBase struct {
}

func (c CDataBase) GetData() {
}

func (c CDataBase) SetData() {
}

type Adapter struct {
	cdb CDataBase
}

func (a Adapter) Get() {
	a.cdb.GetData()
}

func (a Adapter) Set() {
	a.cdb.SetData()
}

func main() {
	var  database Database
	database = Adapter{cdb: CDataBase{}}
	fmt.Println(database)
}