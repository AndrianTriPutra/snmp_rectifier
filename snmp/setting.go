package snmp

import "time"

type Setting struct {
	TZ        string
	Host      string
	Community string
	Port      uint16
	Timeout   time.Duration
	OID       string
	OIDS      []string
}
