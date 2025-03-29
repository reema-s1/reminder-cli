package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func startScheduler(stopChan chan struct{}) { //start sched
	fmt.Println("⏳ Scheduler started...")

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			checkReminders()
		case <-stopChan:
			fmt.Println("⏹️ Scheduler stopped.")
			return
		}
	}
}

func checkReminders() {
	now := time.Now()
	for _, reminder := range reminders {
		if reminder.Time.Before(now) {
			fmt.Println("🔔 Reminder:", reminder.Message)
		}
	}
}

func handleSignals(stopChan chan struct{}) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\nReceived interrupt, shutting down...")
	close(stopChan)
	os.Exit(0)
}
