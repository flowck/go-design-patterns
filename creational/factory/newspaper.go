package factory

import "fmt"

type newspaper struct {
	publication
}

func (n newspaper) String() string {
	return fmt.Sprintf("This is a newspaper named %s", n.name)
}

func createNewsPaper(name string, pages int, publisher string) IPublication {
	return &newspaper{publication{
		name:      name,
		pages:     pages,
		publisher: publisher,
	}}
}
