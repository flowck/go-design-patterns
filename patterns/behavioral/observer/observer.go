package observer

import "fmt"

type observer interface {
	onUpdate(data string)
}

type DataListener struct {
	Name string
}

func (d *DataListener) onUpdate(data string) {
	fmt.Println("Listener:", d.Name, "got a data change:", data)
}
