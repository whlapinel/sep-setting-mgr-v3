package demodata

import (
	"errors"
	"sep_setting_mgr/internal/domain/models"
)

func (ds *demoDataService) createDemoClasses() ([]*models.Class, error) {
	nameStrings := []string{
		"Math 1",
		"Biology",
		"History",
		"English",
		"Spanish",
		"Chemistry",
		"Earth Science",
		"Art",
		"Music",
		"PE",
		"Health",
	}
	users := ds.demoData.users
	if len(users) != len(nameStrings) {
		return nil, errors.New("number of users and classes do not match")
	}
	for _, user := range users {
		for j := 1; j < 4; j++ {
			class, err := models.NewClass(nameStrings[j], j, user.ID)
			if err != nil {
				return nil, err
			}
			err = ds.classesRepo.Store(class)
			if err != nil {
				return nil, err
			}
		}
	}
	classes, err := ds.classesRepo.All()
	if err != nil {
		return nil, err
	}
	return classes, nil

}
