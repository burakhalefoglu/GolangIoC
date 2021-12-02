package main

import (
	"fmt"
	"github.com/golobby/container/v3"
)

type IDatabase interface{
	Connect()
}
var Database IDatabase


type MySQL struct {

}

func NewMySQL() *MySQL {
	return &MySQL{}
}

func (m *MySQL) Connect() {
	fmt.Println("Connected. İt is works")
}

type OtherInterface interface {
	SomeMetod()
}
var Other OtherInterface

type OtherClass struct {
	Database IDatabase
}

func (o *OtherClass)SomeMetod(){
	o.Database.Connect()
	fmt.Println("Some metod!!!")
}

func NewOtherClass(database IDatabase) *OtherClass {
	return &OtherClass{Database: database}
}

func main() {

	_ = container.Singleton(func() IDatabase {
		return NewMySQL()
	})
	_ = container.Resolve(&Database)


	_ = container.Singleton(func() OtherInterface {
		return NewOtherClass(Database)
	})
	_ = container.Resolve(&Other)

	Other.SomeMetod()

}
