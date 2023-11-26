package actions

import "strings"

func ValidateUserInputs(name string, last_name string, tichektscount uint, email string, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(name) >=2 && len(last_name) >= 2
	isVaildEmail := strings.Contains(email, "@")
	isVaildTick := tichektscount > 0 && tichektscount <= remainingTickets
	return isVaildEmail, isVaildTick, isValidName
}
