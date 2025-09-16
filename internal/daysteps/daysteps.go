package daysteps

import (
	"fmt"
	"time"
	"strings"
	"strconv"

	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

type DaySteps struct {
	Steps int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	vals := strings.Split(datastring, ",")
	if len(vals) != 2 {
		return fmt.Errorf("conversion error: String contains more values. Example: 678,0h50m")
	}

	counter, err := strconv.Atoi(vals[0])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}

	duration, err := time.ParseDuration(vals[1])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}

	if counter <= 0 || duration <= 0 {
		return fmt.Errorf("null error: Empty variable value")
	}

	ds.Steps = counter
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Weight <= 0 || ds.Height <= 0 {
		return "", fmt.Errorf("null error: Empty variable value")
	}

	calory, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, spentenergy.Distance(ds.Steps, ds.Height), calory), nil
}
