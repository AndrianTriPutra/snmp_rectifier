package domain

type Data struct {
	ID        string    `json:"id"`
	TS        string    `json:"ts"`
	AC_Dist   AC_Dist   `json:"ac_distribution"`
	DC_Dist   DC_Dist   `json:"dc_distribution"`
	Rectifier Rectifier `json:"rectifier"`
}

type AC_Dist struct {
	Voltage_R float32 `json:"voltage_R"`
	Voltage_S float32 `json:"voltage_S"`
	Voltage_T float32 `json:"voltage_T"`
	Current_R float32 `json:"current_R"`
	Current_S float32 `json:"current_S"`
	Current_T float32 `json:"current_T"`
}

type DC_Dist struct {
	Bus_Voltage       float32 `json:"bus_voltage"`
	LoadCurrent_Total float32 `json:"loadcurrent_total"`
	Battery_Current   float32 `json:"battery_current"`
}

type Rectifier struct {
	Output_Voltage float32 `json:"output_voltage_1"`
	Output_Current float32 `json:"output_current_1"`
}
