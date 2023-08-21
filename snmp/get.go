package snmp

import (
	"errors"
	"fmt"
	"snmp_rectifier/domain"
	"time"

	g "github.com/gosnmp/gosnmp"
)

func GET(setting Setting) (domain.Data, error) {
	var data domain.Data
	data.ID = setting.Host

	loc, _ := time.LoadLocation(setting.TZ)
	now := time.Now().In(loc)
	ts := now.Format(time.RFC3339)
	data.TS = ts

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
		return data, errN
	}
	defer params.Conn.Close()

	result, err := params.Get(setting.OIDS)
	if err != nil {
		errN := errors.New("failed get " + err.Error())
		return data, errN
	}

	var buffer [11]float32
	for i, variable := range result.Variables {
		fmt.Printf("oid: %s ", variable.Name)

		switch variable.Type {
		case g.OctetString:
			fmt.Printf("string : %s\n", string(variable.Value.([]byte)))
		case g.Uinteger32:
			fmt.Printf("uint32 : %d\n", variable.Value)
		case g.Integer:
			fmt.Printf("integer : %d\n", g.ToBigInt(variable.Value))
			integer := variable.Value.(int)
			buffer[i] = float32(integer)
		}
	}

	var ac domain.AC_Dist
	ac.Voltage_R = buffer[0] * 0.01
	ac.Voltage_S = buffer[1] * 0.01
	ac.Voltage_T = buffer[2] * 0.01
	ac.Current_R = buffer[3] * 0.01
	ac.Current_S = buffer[4] * 0.01
	ac.Current_T = buffer[5] * 0.01

	var dc domain.DC_Dist
	dc.Bus_Voltage = buffer[6] * 0.01
	dc.LoadCurrent_Total = buffer[7] * 0.01
	dc.Battery_Current = buffer[8] * 0.01

	var recti domain.Rectifier
	recti.Output_Voltage = buffer[9] * 0.01
	recti.Output_Current = buffer[10] * 0.01

	data.AC_Dist = ac
	data.DC_Dist = dc
	data.Rectifier = recti

	return data, nil
}
