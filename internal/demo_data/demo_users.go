package demodata

import "sep_setting_mgr/internal/domain/models"

func (ds *demoDataService) createDemoUsers() ([]*models.User, error) {
	emailDomain := "@cms.k12.nc.us"
	demoUsers := []models.User{
		{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe" + emailDomain,
			Picture:   "https://randomuser.me",
		},
		{
			FirstName: "Jane",
			LastName:  "Smith",
			Email:     "jane.smith" + emailDomain,
			Picture:   "https://randomuser.me",
		},
		{
			FirstName: "Bob",
			LastName:  "Johnson",
			Email:     "bob.johnson" + emailDomain,
			Picture:   "https://randomuser.me",
		},
		{
			FirstName: "Sally",
			LastName:  "Brown",
			Email:     "sally.brown" + emailDomain,
			Picture:   "https://randomuser.me",
		},
		{
			FirstName: "Tom",
			LastName:  "Jones",
			Email:     "tom.jones" + emailDomain,
			Picture:   "https://randomuser.me",
		},
		{
			FirstName: "Mary",
			LastName:  "Davis",
			Email:     "mary.davis" + emailDomain,
			Picture:   "https://randomuser.me",
		},
		{
			FirstName: "Chris",
			LastName:  "Wilson",
			Email:     "chris.wilson" + emailDomain,
			Picture:   "https://randomuser.me",
		},
		{
			FirstName: "Lisa",
			LastName:  "Martinez",
			Email:     "lisa.martinez" + emailDomain,
			Picture:   "https://randomuser.me",
		},
		{
			FirstName: "David",
			LastName:  "Hernandez",
			Email:     "david.hernandez" + emailDomain,
			Picture:   "https://randomuser.me",
		},
		{
			FirstName: "Karen",
			LastName:  "Young",
			Email:     "karen.young" + emailDomain,
			Picture:   "https://randomuser.me",
		},
		{
			FirstName: "William",
			LastName:  "Lapinel",
			Email:     "williamh.lapinel" + emailDomain,
			Picture:   "https://lh3.googleusercontent.com/a/ACg8ocKZ7iCB2XYFNu0_1wE5nYiGtee0eNbyuk4WiS3N-HKd4mXwgQ=s96-c",
			Admin:     true,
			Teacher:   true,
		},
	}
	for _, user := range demoUsers {
		err := ds.usersRepo.Store(&user)
		if err != nil {
			return nil, err
		}
	}
	users, err := ds.usersRepo.All()
	if err != nil {
		return nil, err
	}
	return users, nil
}
