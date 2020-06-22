package main

type Implant struct {
	Dbid        int
	Identifier  string
	Ipaddress   string
	Lastcheckin string
}

var implants []Implant

type Task struct {
	Implantid string `json:"implantid"`
	Taskcode  string `json:"taskcode"`
	Args      string `json:"args"`
}

var tasks []Task

type Command struct {
	Id        int
	Task_code int
	Name      string
}

var commands []Command
