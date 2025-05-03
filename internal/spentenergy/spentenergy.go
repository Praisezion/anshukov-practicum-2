package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("steps must be > 0, now %d", steps)
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be >= 0, now %.2f", height)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be >= 0, now %.2f", weight)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be != 0, now %v", duration)
	}

	mean := MeanSpeed(steps, height, duration)

	spent := (weight * mean * duration.Minutes()) / minInH * walkingCaloriesCoefficient

	return spent, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("steps must be > 0, now %d", steps)
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be >= 0, now %.2f", height)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be >= 0, now %.2f", weight)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be != 0, now %v", duration)
	}

	mean := MeanSpeed(steps, height, duration)

	spent := (weight * mean * duration.Minutes()) / minInH

	return spent, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration < 0 {
		return 0
	}
	distance := Distance(steps, height)
	hours := duration.Hours()

	if hours == 0 {
		return 0
	}

	return distance / hours
}

func Distance(steps int, height float64) float64 {
	if steps < 0 || height < 0 {
		return 0
	}
	return height * stepLengthCoefficient * float64(steps) / mInKm
}
