package dto

import (
	"time"
)

type CostReportDto struct{
	Month string `json:"month"`
	Count float64 `json:"count"`
}

func NewCostReport() []CostReportDto{
	var month = []string{"January", "February", "March", "April", "May", "June","July", "August", "September", "October", "November", "December"}
	monthNow:= int(time.Now().Month())
	var costReport []CostReportDto
	for  i := 0; i <= monthNow-1;  i++ {
		costReport= append(costReport, CostReportDto{Month: month[i], Count: 0} )
	}
	return costReport
}

func (report CostReportDto) GetIntMonth() int{
	months := map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}
	return months[report.Month]
}

