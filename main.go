package main

import (
	"fmt"
)

func main() {
	fmt.Println("📅 Reminder Scheduler Started...")

	stopChan := make(chan struct{})

	go startScheduler(stopChan)

	go handleSignals(stopChan)

	for {
		fmt.Println("\n--- Reminder Manager ---")
		fmt.Println("1. Add Reminder")
		fmt.Println("2. List Reminders")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		_, err := fmt.Scanf("%d\n", &choice)
		if err != nil {
			fmt.Println("❌ Invalid input! Please enter 1-3.")
			continue
		}

		switch choice {
		case 1:
			addReminder()
		case 2:
			listReminders()
		case 3:
			fmt.Println("Exiting...")
			close(stopChan)
			return
		default:
			fmt.Println("❌ Invalid choice! Please enter 1-3.")
		}
	}
}
