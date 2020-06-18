package main

type Implant struct {
	Dbid        int    `json:"dbid"`
	Identifier  string `json:"id"`
	Ipaddress   string `json:"ipaddress"`
	Lastcheckin string `json:"lastcheckin`
}

var implants []Implant

type Task struct {
	Implantid string `json:"implantid"`
	Taskcode  string `json:"taskcode"`
	Args      string `json:"args"`
}

var tasks []Task
