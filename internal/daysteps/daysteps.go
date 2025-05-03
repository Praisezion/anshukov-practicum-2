package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return fmt.Errorf("unexpected data format")
	}
	if strings.Contains(slice[0], " ") {
		return fmt.Errorf("steps format is invalid: has spaces: '%v'", slice[0])
	}

	stepsNoSpace := strings.TrimSpace(slice[0])

	steps, err := strconv.Atoi(stepsNoSpace)
	if err != nil {
		return fmt.Errorf("steps format is invalid: %v, error: %w", slice[0], err)
	}

	if steps <= 0 {
		return fmt.Errorf("steps must be positive, now: %d", steps)
	}

	ds.Steps = steps

	duration, err := time.ParseDuration(strings.TrimSpace(slice[1]))
	if err != nil {
		return fmt.Errorf("duration format is invalid %v, error: %w", slice[1], err)
	}

	if duration <= 0 {
		return fmt.Errorf("duration must be positive, now: %d", steps)
	}

	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("calories calculation failed: %w", err)
	}

	message := fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories)

	return message, nil
}
