package enums

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Months int

const (
	January Months = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func (day Weekday) Weekend() bool {
	switch day {
	case Sunday, Saturday:
		return true
	default:
		return false
	}
}

func Run() {
	fmt.Println(Sunday)
	fmt.Println(Saturday)

	fmt.Printf("Is Saturday a weekend day? %t\n", Saturday.Weekend())

	fmt.Println(January)
	fmt.Println(December)
}
