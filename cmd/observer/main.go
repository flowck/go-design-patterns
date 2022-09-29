package main

import "go-design-patterns/patterns/behavioral/observer"

func main() {
	list1 := observer.DataListener{Name: "Listener 1"}
	list2 := observer.DataListener{Name: "Listener 2"}

	subj := &observer.DataSubject{}

	subj.Register(list1)
	subj.Register(list2)

	subj.ChangeItem("Monday!")
	subj.ChangeItem("Wednesday!")

	subj.Unregister(list1)

	subj.ChangeItem("Friday!")
}
