package main

import (
	"log"
	"testing"
)

func TestNotificationBuilder_Build(t *testing.T) {
	builder := newNotificationBuilder()

	builder.SetTitle("Login failed")
	builder.SetSubTitle("Email and Password are missing")
	builder.SetPriority(10)

	notification, err := builder.Build()

	if err != nil {
		t.Logf("Expected an error because priority must be between 0 5: %s", err)
	} else {
		t.Fatal("Expected an error instead got nil")
	}

	if notification == nil {
		t.Log("Notification must be nil because an error was also returned")
	} else {
		log.Fatalf("notification is expected to be nil instead got %+v", notification)
	}
}
