package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"snmp_rectifier/snmp"
	"time"
)

func main() {
	flag.Usage = func() {
		log.Printf("Usage: go run . host methode")
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(flag.Args()) != 2 {
		flag.Usage()
		os.Exit(1)
	}

	host := flag.Args()[0]
	methode := flag.Args()[1]

	var oid = []string{
		// AC DISTRIBUTION
		".1.3.6.1.4.1.39553.10.1.1.1403.0", // voltage R (A)
		".1.3.6.1.4.1.39553.10.1.1.1404.0", // voltage S (A)
		".1.3.6.1.4.1.39553.10.1.1.1405.0", // voltage T (A)
		".1.3.6.1.4.1.39553.10.1.1.5.0",    // current R (V)
		".1.3.6.1.4.1.39553.10.1.1.6.0",    // current S (V)
		".1.3.6.1.4.1.39553.10.1.1.7.0",    // current T (V)

		// DC DISTRIBUTION
		".1.3.6.1.4.1.39553.10.3.1.1.0", // bus voltage      (V)
		".1.3.6.1.4.1.39553.10.3.1.2.0", // load current tot (A)
		".1.3.6.1.4.1.39553.10.3.1.4.0", // battery current  (A)

		// RECTIFIER
		".1.3.6.1.4.1.39553.10.2.1.1.0",   // output voltage (V)
		".1.3.6.1.4.1.39553.10.2.1.201.0"} // output current (A)

	setting := snmp.Setting{
		TZ:        "Asia/Jakarta",
		Host:      host,
		Community: "public",
		Port:      161,
		Timeout:   250 * time.Millisecond,
		OID:       ".1.3.6.1.4.1.39553.10.1.1",
		OIDS:      oid,
	}
	fmt.Print("setting:")
	fmt.Println(setting)

	if methode == "get" {
		log.Print("[snmp GET]")
		data, err := snmp.GET(setting)
		if err != nil {
			log.Printf("error snmp get:%s", err.Error())
		}
		jw, err := json.MarshalIndent(data, " ", " ")
		if err == nil {
			log.Printf("data:\n%s", string(jw))
		}
	} else if methode == "walk" {
		log.Println("[snmp WALK]")
		err := snmp.Walk(setting)
		if err != nil {
			log.Printf("error snmp walk:%s", err.Error())
		}
	}
}
