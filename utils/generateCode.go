package utils

import (
	"fmt"
	"time"
)

// GenerateEmployeeCode generates a new employee code in the format yymm + running number
func GenerateEmployeeCode(latestCode string) string {
	currentTime := time.Now()
	year := currentTime.Year() % 100   // get last two digits of the year
	month := int(currentTime.Month()) // get the month as an integer

	var newRunningNumber int
	if latestCode != "" {
		// Extract the running number from the latest code
		latestRunningNumber := latestCode[4:]
		fmt.Sscanf(latestRunningNumber, "%d", &newRunningNumber)
	}

	// Increment the running number
	newRunningNumber++

	// Format the new code as yymm + running number
	newCode := fmt.Sprintf("%02d%02d%04d", year, month, newRunningNumber)
	 fmt.Println(newCode)
	return newCode
}
