package utils

import (
	"github.com/crazy-max/WindowsSpyBlocker/app/whois"
)

func GetWhoIs(serverAddress string) whois.Whois {
	whoisInfo := whois.GetWhois(serverAddress)
	return whoisInfo
}
