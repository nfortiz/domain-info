package models

import (
	"time"
)

type Server struct {
	Address string `json:"address"`
	SSLGrade string `json:"ssl_grade"`
	Country string `json:"country"`
	Owner string `json:"owner"`
}

type Domain struct {
	Name string `json:"domain"`
	ServersChanged bool `json:"servers_changed"`
	SSLGrade string `json:"ssl_grade"`
	PreviousSSLGrade string `json:"previous_ssl_grade"`
	Logo string `json:"logo"`
	Title string `json:"title"`
	IsDown bool `json:"is_down"`
	Servers []Server `json:"servers"`
	DateChecked time.Time `json:"-"`
}

type History struct {
	Items []Domain `json:"items"`
}

type ServerSSLLabs struct {
	Address string `json:"ipAddress"`
	Grade string `json:"grade"`
}

type SSLLabs struct {
	Status string `json:"status"`
	Servers []ServerSSLLabs `json:"endpoints"`
}
