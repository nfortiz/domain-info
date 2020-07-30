package utils

import (
	"truora-server/models"
)

func SetInfoServers(servers []models.ServerSSLLabs, resp *models.Domain) {
	serversInfo := make([]models.Server, len(servers))
	minGrade := "A+"
	for idx, server := range servers {
		whois := GetWhoIs(server.Address)
		serversInfo[idx] = models.Server{
			SSLGrade: server.Grade,
			Address: server.Address,
			Country: whois.Country,
			Owner: whois.Org,
		}

		if GetGrade(minGrade) > GetGrade(server.Grade) {
			minGrade = server.Grade
		}
	}

	resp.Servers = serversInfo
	resp.SSLGrade = minGrade
}
