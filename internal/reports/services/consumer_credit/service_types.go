package consumer_credit

import (
	"github.com/garrettladley/garrettladley/internal/reports/types"
)

type Data struct {
	StudentLoans      types.TimeSeries
	MotorVehicleLoans types.TimeSeries
}
