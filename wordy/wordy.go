package wordy

import (
	"errors"
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	if !(question[:8] == "What is ") {
		return 0, false
	}
	remainingQuestion := question[8:]

	firstNum := remainingQuestion[:getCurrentWordEndIndx(remainingQuestion)]
	res, err := strconv.Atoi(firstNum)
	if err != nil {
		return 0, false
	}

	remainingQuestion = remainingQuestion[getCurrentWordEndIndx(remainingQuestion):]
	if remainingQuestion == "?" {
		return res, true
	}

	remainingQuestion = remainingQuestion[1:]

	currentWordEndIndx := getCurrentWordEndIndx(remainingQuestion)
	operation := remainingQuestion[:currentWordEndIndx]
	remainingQuestion = remainingQuestion[currentWordEndIndx+1:]

	result, isSuccess, remainingQuestion, err := executeOperation(operation, remainingQuestion, res)
	if err != nil {
		return 0, false
	}

	if len(remainingQuestion) < 2 {
		return result, isSuccess
	}

	currentWordEndIndx = getCurrentWordEndIndx(remainingQuestion)
	operation = remainingQuestion[:currentWordEndIndx]
	remainingQuestion = remainingQuestion[currentWordEndIndx+1:]
	result, isSuccess, remainingQuestion, err = executeOperation(operation, remainingQuestion, result)
	if err != nil {
		return 0, false
	}

	if len(remainingQuestion) < 2 {
		return result, isSuccess
	}

	return 0, true
}

func executeOperation(operation string, remainingQuestion string, res int) (int, bool, string, error) {
	switch operation {
	case "plus":
		{
			nextSpace := getCurrentWordEndIndx(remainingQuestion)
			if nextSpace <= 0 {
				return 0, false, remainingQuestion[nextSpace+1:], nil
			}
			secondNum, err := strconv.Atoi(remainingQuestion[0:nextSpace])
			if err != nil {
				return 0, false, remainingQuestion[nextSpace+1:], nil
			}
			return secondNum + res, true, remainingQuestion[nextSpace+1:], nil
		}
	case "minus":
		{
			nextSpace := getCurrentWordEndIndx(remainingQuestion)
			secondNum, err := strconv.Atoi(remainingQuestion[0:nextSpace])
			if err != nil {
				return 0, false, remainingQuestion[nextSpace+1:], nil
			}
			return res - secondNum, true, remainingQuestion[nextSpace+1:], nil
		}
	case "multiplied":
		{
			if remainingQuestion[:2] != "by" {
				return 0, false, remainingQuestion, errors.New("Incorrect")
			}

			remainingQuestion = remainingQuestion[3:]
			nextSpace := getCurrentWordEndIndx(remainingQuestion)
			secondNum, err := strconv.Atoi(remainingQuestion[0:nextSpace])
			if err != nil {
				return 0, false, remainingQuestion[nextSpace+1:], nil
			}
			return res * secondNum, true, remainingQuestion[nextSpace+1:], nil
		}
	case "divided":
		{
			if remainingQuestion[:2] != "by" {
				return 0, false, remainingQuestion, nil
			}

			remainingQuestion = remainingQuestion[3:]
			nextSpace := getCurrentWordEndIndx(remainingQuestion)
			secondNum, err := strconv.Atoi(remainingQuestion[0:nextSpace])
			if err != nil {
				return 0, false, remainingQuestion[nextSpace+1:], nil
			}
			return res / secondNum, true, remainingQuestion[nextSpace+1:], nil

		}
	}
	return 0, false, remainingQuestion, errors.New("Incorrect")
}

func getCurrentWordEndIndx(question string) int {
	nextSpace := strings.IndexByte(question, ' ')
	if nextSpace < 0 {
		nextSpace = strings.IndexByte(question, '?')
	}
	return nextSpace
}
