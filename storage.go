package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var reminders []Reminder

func saveReminders() {
	file, err := os.Create("reminders.json")
	if err != nil {
		fmt.Println("❌ Error saving reminders:", err)
		return
	}
	defer file.Close()

	json.NewEncoder(file).Encode(reminders)
}

func loadReminders() {
	file, err := os.Open("reminders.json")
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("❌ Error loading reminders:", err)
		}
		return
	}
	defer file.Close()

	json.NewDecoder(file).Decode(&reminders)
}
