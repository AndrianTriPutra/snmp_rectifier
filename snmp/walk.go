package snmp

import (
	"errors"
	"fmt"

	g "github.com/gosnmp/gosnmp"
)

func Walk(setting Setting) error {

	params := &g.GoSNMP{
		Target:    setting.Host,
		Port:      setting.Port,
		Community: setting.Community,
		Version:   g.Version2c,
		Timeout:   setting.Timeout,
	}

	err := params.Connect()
	if err != nil {
		errN := errors.New("failed connect " + err.Error())
		return errN
	}
	defer params.Conn.Close()

	err = params.BulkWalk(setting.OID, printValue)
	if err != nil {
		errN := errors.New("failed walk " + err.Error())
		return errN
	}

	return nil
}

func printValue(pdu g.SnmpPDU) error {
	fmt.Printf("oid:%s ", pdu.Name)

	switch pdu.Type {
	case g.OctetString:
		b := pdu.Value.([]byte)
		fmt.Printf("STRING: %s\n", string(b))
	default:
		fmt.Printf("TYPE %d: %d\n", pdu.Type, g.ToBigInt(pdu.Value))
	}
	return nil
}
