package main

import (
	"fmt"
	"time"
)

type Reminder struct {
	Message string
	Time    time.Time
}

func addReminder() {
	fmt.Println("\n📝 Adding a new reminder...")
	fmt.Print("Enter reminder message: ")
	var message string
	_, err := fmt.Scanln(&message)
	if err != nil {
		fmt.Println("❌ Error reading message:", err)
		return
	}

	fmt.Print("Enter time (YYYY-MM-DD HH:MM): ")
	var datePart, timePart string
	_, err = fmt.Scanf("%s %s", &datePart, &timePart)
	if err != nil {
		fmt.Println("❌ Error reading time:", err)
		return
	}

	timeInput := datePart + " " + timePart
	reminderTime, err := time.Parse("2006-01-02 15:04", timeInput)
	if err != nil {
		fmt.Println("❌ Invalid time format! Use YYYY-MM-DD HH:MM")
		return
	}

	reminders = append(reminders, Reminder{Message: message, Time: reminderTime})
	saveReminders()

	fmt.Println("✅ Reminder added for:", reminderTime)
}

func listReminders() {
	fmt.Println("\n📜 Your Reminders:")
	if len(reminders) == 0 {
		fmt.Println("No reminders set.")
		return
	}

	for i, reminder := range reminders {
		fmt.Printf("%d. [%s] - %s\n", i+1, reminder.Time.Format("2006-01-02 15:04"), reminder.Message)
	}
}
