package professions

import (
	"fmt"
	"log"
	"time"
)

const customDateFormat string = "02-Jan-2006"

type EmployeeService interface {
	GetEmployeeInfo() string
	GetEmployeePosition() string
}

type Employee struct {
	ID         uint64 `json:"id" bson:"_id"`
	*Person    `json:"person" bson:"person"`
	HiringDate time.Time `json:"hiringDate,createdAt" bson:"hiringDate"`
	Salary     int       `json:"salary,omitempty" bson:"salary,omitempty"`
	FullTime   bool      `json:"fullTime,omitempty" bson:"fullTime,omitempty"`
	Position   string    `json:"position" bson:"position"`
}

func (e *Employee) ConvertEmployeeToPerson() *Person {
	return e.Person
}

/*func (employee *Employee) GetEmployeePosition() string {
	return employee.GetPersonInfo() + ", position: " + employee.Position
}

func (employee *Employee) GetEmployeeInfo() string {
	return employee.Person.GetPersonInfo() + "\nHiring date: " + string(employee.HiringDate.Format(customDateFormat)) + ", salary: $ " + strconv.Itoa(employee.Salary) + "\nFull time: " + strconv.FormatBool(employee.FullTime)
}*/

type PersonService interface {
	GetPersonInfo() string
}

type Person struct {
	ID   uint64 `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	/*LastName    string `json:"lastName" bson:"lastName"`
	Age         int    `json:"age,omitempty" bson:"age,omitempty"`
	Nationality string `json:"-" bson:"nationality"`
	Email       string `json:"-" bson:"email"`
	Phone       string `json:"phone,omitempty" bson:"phone,omitempty"`*/
}

func (p *Person) ConvertPersonToEmployee() (e *Employee, err error) {
	if p != nil {
		e = &Employee{Person: p}
		return
	} else {
		log.Fatal(fmt.Errorf("", "error ConvertPersonToEmployee: person is nil"))
		return e, err
	}
}

/*func (person *Person) GetPersonInfo() string {
	return person.Name + " " + person.LastName + ", " + strconv.Itoa(person.Age) + " years old"
}

func (person *Person) GetContacts() string {
	return person.Phone + ", " + person.Email
}

func (person *Person) SetAge(age int) bool {
	if person != nil && age >= 0 {
		person.Age = age
		return true
	}
	return false
}

func (person *Person) NewPersonWasBorned(name, lastName, nationality string) bool {
	if person == nil {
		person = new(Person)
		person.Name = name
		person.LastName = lastName
		person.Nationality = nationality
		return true
	}
	return false
}*/
