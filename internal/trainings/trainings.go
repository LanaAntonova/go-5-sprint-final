package trainings

import (
	"fmt"
	"time"
	"strings"
	"strconv"

	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

type Training struct {
	Steps int
	TrainingType string 
	Duration time.Duration 
	personaldata.Personal 
}

func (t *Training) Parse(datastring string) (err error) {
	vals := strings.Split(datastring, ",")
	if len(vals) != 3 {
		return fmt.Errorf("conversion error: String contains more values. Example: 3456,Ходьба,3h00m")
	}

	counter, err := strconv.Atoi(vals[0])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}

	duration, err := time.ParseDuration(vals[2])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}

	if counter <= 0 || duration <= 0 || vals[1] <= "" {
		return fmt.Errorf("null error: Empty variable value")
	}

	t.Steps =  counter
	t.TrainingType = vals[1]
	t.Duration = duration
	
	return nil
}

func (t Training) ActionInfo() (string, error) {
	switch t.TrainingType {
	case "Ходьба":
		calory, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps, t.Height), spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration), calory), nil
	case "Бег":
		calory, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps, t.Height), spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration), calory), nil
	default:
		return "", fmt.Errorf("Unknown type of activity")
	}
}
