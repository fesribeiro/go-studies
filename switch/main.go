package main

import "fmt"


func getWeekDay(number int8) string {

	switch number {
		case 1:
			return "Monday"
		case 2:
			return "Tuesday"
		case 3:
			return "Wednesday"
		case 4:
			return "Thursday"
		case 5:
			return "Friday"
		case 6:
			return "Saturday"
		case 7:
			return "Sunday"
		default:
			return "Invalid day"
	}
}


func getWeekDay2(number int8) string {
	var weekDay string
	switch {
		case number == 1:
			weekDay = "Monday"
			fallthrough // fallthrough allows the next case to be executed even if the condition is true (just for demonstration)
		case number == 2:
			weekDay = "Tuesday"
		case number == 3:
			weekDay = "Wednesday"
		case number == 4:
			weekDay = "Thursday"
		case number == 5:
			weekDay = "Friday"
		case number == 6:
			weekDay = "Saturday"
		case number == 7:
			weekDay = "Sunday"
		default:
			weekDay = "Invalid day"
	}
	return weekDay
}

func main() {
	var day int8
	fmt.Print("Enter a day number (1-7): ")
	fmt.Scanln(&day)
	weekDay := getWeekDay(day)
	weekDay2 := getWeekDay2(day)
	fmt.Println("Using switch with case:", weekDay)
	fmt.Println("Using switch with conditions:", weekDay2)

}
