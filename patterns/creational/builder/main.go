package main

import "fmt"

func main() {
	builder := newNotificationBuilder()

	builder.SetTitle("New Notification")
	builder.SetIcon("image.jpg")
	builder.SetSubTitle("Cool image")
	builder.SetPriority(5)
	builder.SetMessage("This is a basic Notification with some text in it")
	builder.SetType("alert")

	notification, err := builder.Build()

	if err != nil {
		fmt.Println("error creating the notification", err)
		return
	}

	fmt.Printf("Notification  %+v", notification)
}
