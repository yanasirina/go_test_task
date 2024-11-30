package service

import (
	"github.com/yanasirina/umbrella_corp_test_task/pkg"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (service *Service) GetDaysBefore2025() int {
	neededDate := pkg.Date(2025, 1, 1)
	durationToDate := time.Until(neededDate)
	daysBeforeDate := int(durationToDate.Hours()) / 24
	return daysBeforeDate
}
