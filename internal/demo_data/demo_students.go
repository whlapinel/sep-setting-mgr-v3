package demodata

import (
	"math/rand"
	"sep_setting_mgr/internal/domain/models"
	"strings"
)

func (ds *demoDataService) createDemoStudents() ([]*models.Student, error) {
	classes := ds.demoData.classes
	var firstNamesArr = strings.Fields(firstNamesStr)
	var lastNamesArr = strings.Fields(lastNamesStr)
	var randomName = func(nameList []string) string {
		return nameList[rand.Intn(len(nameList))]
	}
	for _, class := range classes {
		for i := 0; i < 10; i++ {
			student, err := models.NewStudent(randomName(firstNamesArr), randomName(lastNamesArr), class.ID, i%10 == 0)
			if err != nil {
				return nil, err
			}
			err = ds.studentsRepo.Store(student)
			if err != nil {
				return nil, err
			}
		}

	}
	students, err := ds.studentsRepo.All()
	if err != nil {
		return nil, err
	}
	return students, nil
}

// create a list of 100 random first names
const firstNamesStr = `John Jane Bob Sally Tom Mary Chris Lisa David Sarah Mike Laura James Jennifer Robert 
Linda William Karen Richard Donna Charles Patricia Joseph Debra Daniel Amanda Steven Kelly Matthew Barbara Mark 
Elizabeth Donald Helen Paul Sandra George Nancy Kenneth Maria Edward Susan Brian Margaret Ronald Dorothy Kevin 
Jessica Jason Debra Timothy Michelle Jeffrey Cynthia Frank Angela Scott Amy Eric Shirley Stephen Kathleen Gary 
Carolyn Gregory Ruth Joshua Anna Jerry Virginia Dennis Rebecca Walter Janet Peter Rachel Harold Christine Ethan 
Marie Raymond Joyce Benjamin Janice Roy Emma Harry Julie Howard Heather Jesse Theresa Lawrence Beverly Nicholas 
Frances Alan Joan Philip Alice Eugene Jean Carl Victoria Bryan Grace Samuel Lauren Louis Martha Jordan Cheryl 
Randy Megan Wayne Jacqueline Adam Sara Jack Rose Billy Janice Jonathan Julia Brandon Pamela Russell Judith Alan 
Martha Geraldine Phillip Bonnie Craig Norma Aaron Lillian Bobby Kathleen Johnny Lori Terry Tammy Gerald Tamara 
Arthur Diane Roger Amanda Royce Marilyn Martin Janice Albert Kathryn Willie Rosemary Howard Cindy Earl Bonnie 
Stanley Jeanne Joe Gail Fred Joanne Eugene Maureen Samuel`

const lastNamesStr = `Smith Johnson Williams Jones Brown Davis Miller Wilson Moore Taylor Anderson Thomas Jackson White
Harris Martin Thompson Garcia Martinez Robinson Clark Rodriguez Lewis Lee Walker Hall Allen Young Hernandez King
Wright Lopez Hill Scott Green Adams Baker Gonzalez Nelson Carter Mitchell Perez Roberts Turner Phillips Campbell
Parker Evans Edwards Collins Stewart Sanchez Morris Rogers Reed Cook Morgan Bell Murphy Bailey Rivera Cooper
Richardson Cox Howard Ward Torres Peterson Gray Ramirez James Watson Brooks Kelly Sanders Price Bennett Wood
Barnes Ross Henderson Coleman Jenkins Perry Powell Long Patterson Hughes Flores Washington Butler Simmons
Foster Gonzales Bryant Alexander Russell Griffin Diaz Hayes Myers Ford Hamilton Graham Sullivan Wallace Woods
Cole West Jordan Owens Reynolds Fisher Ellis Harrison Gibson Mcdonald Cruz Marshall Ortiz Gomez Murray Freeman
Wells Webb Simpson Stevens Tucker Porter Hunter Hicks Crawford Henry Boyd Mason Morales Kennedy Warren Dixon
Ramos Reyes Burns Gordon Shaw Holmes Rice Robertson Hunt Black Daniels Palmer Mills Nichols Grant Knight`
