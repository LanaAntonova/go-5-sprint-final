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

// Функция принимает количество шагов, вес и рост, продолжительность бега и 
// возвращает количество калорий с корректирующим коэффициентом, error
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Empty variable value")
	} else {
		return ((weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH) * walkingCaloriesCoefficient, nil
	}
}

// Функция принимает количество шагов, вес и рост, продолжительность бега и 
// возвращает количество калорий, error
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Empty variable value")
	} else {
		return (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH, nil
	}
}

// Функция принимает количество шагов, рост пользователя и продолжительность 
// активности и возвращает среднюю скорость
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	} else { 
		return Distance(steps, height) / duration.Hours()
	}
}

// Функция принимает количество шагов и рост пользователя в метрах и 
// возвращает дистанцию в километрах
func Distance(steps int, height float64) float64 {
	return float64(steps) * (height * stepLengthCoefficient) / mInKm
}
