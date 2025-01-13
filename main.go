package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scores = map[string]int32{
	"A1": 6,
	"B2": 5,
	"B3": 4,
	"C4": 3,
	"C5": 2,
	"C6": 1,
	"E8": 0,
	"F9": 0,
}

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

func getUserName(reader *bufio.Reader) string {
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
	return name
}

func getUtmeScore(reader *bufio.Reader) int64 {
	var utmeScore int64
	for {
		value, err := getUserInput(reader, "your UTME score")
		i, strErr := strconv.ParseInt(value, 10, 32)
		utmeScore = i

		if utmeScore > 400 {
			fmt.Println("Invalid UTME score, score cannot be greater than 400")
			continue
		}
		if strErr != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			break
		}
	}
	return utmeScore
}

func getNumberOfSitting(reader *bufio.Reader) int32 {
	var numOfsitting int32
	for {

		value, err := getUserInput(reader, "number of sitting(s)")

		if err != nil {
			fmt.Println("Error: ", err)
		}

		i, strErr := strconv.ParseInt(value, 10, 32)

		if strErr != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch i {
		case 1, 2:
			numOfsitting = int32(i)
			break
		default:
			fmt.Println("Invalid input. Expected 1 or 2")
			continue
		}
		break
	}
	return numOfsitting
}

func printResult(c *Composite, total float64) {
	fmt.Println("Hi, ", c.name)
	fmt.Printf("Your total composite is %0.2f ", total)
}

func main() {
	user := createComposite()
	reader := bufio.NewReader(os.Stdin)

	//Get and save the username
	name := getUserName(reader)
	user.updateName(name)

	//record grades
	recordGrade("english", reader, &user)
	recordGrade("maths", reader, &user)
	recordGrade("chemistry", reader, &user)
	recordGrade("biology", reader, &user)
	recordGrade("physics", reader, &user)

	mathScore := scores[user.maths]
	chemistryScore := scores[user.chemistry]
	biologyScore := scores[user.biology]
	physicsScore := scores[user.physics]
	englishScore := scores[user.english]

	totalOlevelScore := mathScore + chemistryScore + biologyScore + physicsScore + englishScore

	// Get and save the utme score
	utmeScore := getUtmeScore(reader)
	user.updateUtmeScore(utmeScore)
	utmeScorePercent := ((float64(user.utme) / 400) * 100) * 0.6

	// Get and save number of sittings
	numOfsitting := getNumberOfSitting(reader)
	user.updateNumOfSitting(numOfsitting)

	var numOfSittingPercent int32
	if user.numOfSitting == 1 {
		numOfSittingPercent = 10
	} else {
		numOfSittingPercent = 0
	}

	// Get total composite
	totalScore := utmeScorePercent + float64(totalOlevelScore) + float64(numOfSittingPercent)

	printResult(&user, totalScore)
}
