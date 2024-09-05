package consumer_credit

import (
	"testing"

	"github.com/garrettladley/garrettladley/reports/pkg/types"
)

func TestParse(t *testing.T) {
	t.Parallel()

	s := `
        2024 Q2
        1,744,342.61
        1,565,116.90
      `

	var (
		actual = parse(s)

		expectedStudentLoans = types.TimeSeries{
			QuarterDate: types.QuarterDate{
				Year:    2024,
				Quarter: types.Quarter2,
			},
			Data: 1_744_342.61,
		}

		expectedMotorVehicleLoans = types.TimeSeries{
			QuarterDate: types.QuarterDate{
				Year:    2024,
				Quarter: types.Quarter2,
			},
			Data: 1_565_116.90,
		}
	)
	if actual.StudentLoans != expectedStudentLoans {
		t.Fatalf("expected studentLoans to be %v, got %v", expectedStudentLoans, actual.StudentLoans)
	}

	if actual.MotorVehicleLoans != expectedMotorVehicleLoans {
		t.Fatalf("expected motorVehicleLoans to be %v, got %v", expectedMotorVehicleLoans, actual.MotorVehicleLoans)
	}
}
