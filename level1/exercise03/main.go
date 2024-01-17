package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

var users map[string]map[string]interface{}
var specialties map[string]map[string]map[string][]int
var appointments map[string]map[string]interface{}

func main() {
	InitializeUserData()
	InitializeSpecialtyData()
	InitializeAppointmentData()

	username, err := AuthenticateUser()
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		return
	}

	specialty, _ := GetTextFromUser("Enter the specialty (general medicine, emergency care, clinical analysis, " +
		"cardiology, neurology, nutrition, physiotherapy, traumatology, and internal medicine): ")

	timePreference, _ := GetTextFromUser("When would you like to schedule the medical appointment for? (morning, afternoon): ")

	doctorsAndHours, err := GetDoctorsAndHours(specialty, timePreference)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		return
	}

	DisplayAvailableDoctors(doctorsAndHours, timePreference)
	selectedDoctor, _ := GetTextFromUser("Please enter the doctor you wish to schedule an appointment with: ")
	if !SelectedDoctorIsValid(selectedDoctor, doctorsAndHours) {
		fmt.Println("An error occurred: The doctor entered is not valid.")
		return
	}

	selectedHour, err := GetHourFromUser(selectedDoctor, doctorsAndHours)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		return
	}

	createAppointment(username, selectedDoctor, selectedHour)
}

func InitializeUserData() {
	users = make(map[string]map[string]interface{})

	users["user01"] = map[string]interface{}{
		"password":            "password01",
		"failedLoginAttempts": 0,
	}

	users["user02"] = map[string]interface{}{
		"password":            "password02",
		"failedLoginAttempts": 0,
	}
}

func InitializeSpecialtyData() {
	specialties = make(map[string]map[string]map[string][]int)

	specialties["general medicine"] = map[string]map[string][]int{
		"emily turner": {
			"availableHours": {9, 10, 11},
		},
		"benjamin hayes": {
			"availableHours": {13, 14, 15},
		},
		"olivia foster": {
			"availableHours": {16, 17, 18},
		},
	}

	specialties["emergency care"] = map[string]map[string][]int{
		"ava mitchell": {
			"availableHours": {9, 10, 11},
		},
		"ethan parker": {
			"availableHours": {13, 14, 15},
		},
		"liam reynolds": {
			"availableHours": {16, 17, 18},
		},
	}

	specialties["cardiology"] = map[string]map[string][]int{
		"sophia anderson": {
			"availableHours": {9, 10, 11},
		},
		"jackson carter": {
			"availableHours": {13, 14, 15},
		},
		"isabella taylor": {
			"availableHours": {16, 17, 18},
		},
	}
}

func InitializeAppointmentData() {
	appointments = make(map[string]map[string]interface{})
}

func AuthenticateUser() (string, error) {
	for {
		username, _ := GetTextFromUser("Enter your username: ")
		password, _ := GetTextFromUser("Enter your password: ")

		_, exists := users[username]
		if !exists {
			return "", errors.New("User not found")
		}

		credentialsAreValid := IsValidLogin(username, password)
		if !credentialsAreValid {
			fmt.Printf("Invalid credentials.\n\n")

			IncreaseFailedLoginAttempts(username)
			if UserShouldBeLockedOut(username) {
				return "", errors.New("The maximum number of failed login attempts has been reached. The user has been locked out")
			}

			continue
		}

		ResetFailedLoginAttempts(username)
		return username, nil
	}
}

func IsValidLogin(username string, password string) bool {
	return password == users[username]["password"]
}

func GetTextFromUser(message string) (string, error) {
	fmt.Print(message)

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	line = strings.TrimSuffix(line, "\n")

	return line, nil
}

func GetIntFromUser(message string) int {
	fmt.Print(message)
	var num int
	fmt.Scan(&num)

	return num
}

func IncreaseFailedLoginAttempts(username string) {
	attempts := users[username]["failedLoginAttempts"].(int)
	users[username]["failedLoginAttempts"] = attempts + 1
}

func ResetFailedLoginAttempts(username string) {
	users[username]["failedLoginAttempts"] = 0
}

func UserShouldBeLockedOut(username string) bool {
	return users[username]["failedLoginAttempts"].(int) >= 3
}

func DisplayAvailableDoctors(doctorsAndHours map[string][]int, timePreference string) {
	fmt.Printf("Doctors and available hours for %s appointments:\n", timePreference)
	for doctor, hours := range doctorsAndHours {
		fmt.Printf("%s: %v\n", doctor, hours)
	}
}

func GetDoctorsAndHours(specialty string, preference string) (map[string][]int, error) {
	result := make(map[string][]int)

	doctors, exists := specialties[specialty]
	if !exists {
		return nil, errors.New("Specialty not found.")
	}

	for doctor, schedule := range doctors {
		var hours []int

		switch preference {
		case "morning":
			for _, hour := range schedule["availableHours"] {
				if hour >= 9 && hour <= 12 {
					hours = append(hours, hour)
				}
			}
		case "afternoon":
			for _, hour := range schedule["availableHours"] {
				if hour >= 13 && hour <= 18 {
					hours = append(hours, hour)
				}
			}
		default:
			return nil, errors.New("Invalid preference. Please choose 'morning' or 'afternoon'.")
		}

		if len(hours) > 0 {
			result[doctor] = hours
		}
	}

	return result, nil
}

func SelectedDoctorIsValid(doctor string, doctorsAndHours map[string][]int) bool {
	_, exists := doctorsAndHours[doctor]

	return exists
}

func GetHourFromUser(doctor string, doctorsAndHours map[string][]int) (int, error) {
	availableHours := doctorsAndHours[doctor]
	message := fmt.Sprintf("Please select an hour %v: ", availableHours)
	selectedHour := GetIntFromUser(message)

	if !isInSlice(selectedHour, availableHours) {
		return 0, errors.New("The selected hour is not valid")
	}

	return selectedHour, nil
}

func isInSlice(target int, slice []int) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}

func createAppointment(patient string, doctor string, hour int) {
	appointmentId := uuid.New().String()

	appointments[appointmentId] = map[string]interface{}{
		"patient": patient,
		"doctor":  doctor,
		"hour":    hour,
	}
}
