package db

type Cmd struct {
	Id      string
	Commond string
	Agent   []Agent
}
type Agent struct {
	Id       string
	Ip       string
	HostName string
}
