package db

type Agent struct {
	Id       string
	Ip       string
	HostName string
}
type Cmd struct {
	Id      string
	Commond string
	Agent   []Agent
}
