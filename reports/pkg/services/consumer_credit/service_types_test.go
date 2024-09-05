package consumer_credit

import (
	"testing"
)

func TestTimeSeriesFrom(t *testing.T) {
	s := `
        2024 Q2
        1,744,342.61
        1,565,116.90
      `

	var (
		studentLoans, motorVehicleLoans = TimeSeriesFrom(s)

		expectedStudentLoans = TimeSeries{
			Date: Date{
				Year:    2024,
				Quarter: Quarter2,
			},
			Data: 1_744_342.61,
		}

		expectedMotorVehicleLoans = TimeSeries{
			Date: Date{
				Year:    2024,
				Quarter: Quarter2,
			},
			Data: 1_565_116.90,
		}
	)
	if studentLoans != expectedStudentLoans {
		t.Fatalf("expected studentLoans to be %v, got %v", expectedStudentLoans, studentLoans)
	}

	if motorVehicleLoans != expectedMotorVehicleLoans {
		t.Fatalf("expected motorVehicleLoans to be %v, got %v", expectedMotorVehicleLoans, motorVehicleLoans)
	}
}
