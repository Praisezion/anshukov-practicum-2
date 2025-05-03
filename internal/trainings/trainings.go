package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 3 {
		return fmt.Errorf("wrong data format input")
	}
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("wrong steps count input %v", err)
	}
	if steps <= 0 {
		return fmt.Errorf("steps must be positive")
	}
	t.Steps = steps

	t.TrainingType = strings.TrimSpace(slice[1])

	duration, err := time.ParseDuration(strings.TrimSpace(slice[2]))
	if err != nil {
		return fmt.Errorf("wrong duration input %v", err)
	}
	if duration <= 0 {
		return fmt.Errorf("duration must be positive")
	}
	t.Duration = duration

	return nil

}

func (t Training) ActionInfo() (string, error) {
	dis := spentenergy.Distance(t.Steps, t.Height)
	mean := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}

	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %v", err)
	}

	durationHours := t.Duration.Hours()

	info := fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n",
		t.TrainingType,
		durationHours,
		dis,
		mean,
		calories)

	return info, nil
}
