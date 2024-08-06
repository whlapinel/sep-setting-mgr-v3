package demodata

import (
	"fmt"
	"math/rand"
	"sep_setting_mgr/internal/domain/models"
	"strconv"
	"time"
)

func (ds *demoDataService) createDemoTestEvents() ([]*models.TestEvent, error) {
	classes := ds.demoData.classes
	var randomDate = func() *time.Time {
		date := time.Now().AddDate(0, 0, rand.Intn(90))
		return &date
	}
	var testEvents []*models.TestEvent
	for _, class := range classes {
		for i := 0; i < 3; i++ {
			testEvent, err := models.NewTestEvent(fmt.Sprintf("Unit %s Test", strconv.Itoa(i)), class, randomDate())
			if err != nil {
				return nil, err
			}
			err = ds.testEventsRepo.Store(testEvent)
			if err != nil {
				return nil, err
			}
			testEvents = append(testEvents, testEvent)
		}
	}
	return testEvents, nil
}
