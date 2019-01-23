package db

type User struct {
	Id     int
	Name   string
	Passwd string
}

var Slice []User
var State = make(map[string]interface{})
