package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getUserInput(reader *bufio.Reader, key string) (string, error) {
	fmt.Printf("Enter %v: ", key)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	trimmedString := strings.TrimSpace(text)
	if trimmedString == "" {
		return "", fmt.Errorf("input cannot be empty")
	}
	println(len(trimmedString))
	return trimmedString, nil
}

func checkGrade(grade string) error {
	switch grade {
	case "A1", "B2", "B3", "C4", "C5", "C6", "D7", "E8", "F9":
		return nil
	default:
		return fmt.Errorf("Invalid grade entered. Accepted grades are A1, B2, B3, C4, C5, C6, D7, E8, F9")
	}
}

func recordGrade(subject string, reader *bufio.Reader, user *Composite) {
	var grade string
	var subjectError error
	var err error

	for {
		grade, subjectError = getUserInput(reader, fmt.Sprintf("%v grade", strings.ToTitle(subject)))
		grade = strings.ToUpper(grade)
		err = checkGrade(grade)
		if err != nil {
			fmt.Println("Error: ", err)
		} else if subjectError != nil {
			fmt.Println("Error: ", subjectError)
		} else {
			break
		}
	}
	user.updateSubjectScore(subject, grade)
}

func main() {
	user := createComposite()
	reader := bufio.NewReader(os.Stdin)

	//Save the username
	var name string
	var err error

	for {
		name, err = getUserInput(reader, "your username")
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			break
		}
	}
	user.updateName(name)

	//record grades
	recordGrade("english", reader, &user)
	recordGrade("maths", reader, &user)
	recordGrade("chemistry", reader, &user)
	recordGrade("biology", reader, &user)
	recordGrade("physics", reader, &user)

	scores := map[string]int32{
		"A1": 6,
		"B2": 5,
		"B3": 4,
		"C4": 3,
		"C5": 2,
		"C6": 1,
		"E8": 0,
		"F9": 0,
	}

	mathScore := scores[user.maths]
	chemistryScore := scores[user.chemistry]
	biologyScore := scores[user.biology]
	physicsScore := scores[user.physics]
	englishScore := scores[user.english]

	totalScore := mathScore + chemistryScore + biologyScore + physicsScore + englishScore

	println(totalScore)

	fmt.Println(user)

}
