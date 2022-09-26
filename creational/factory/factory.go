package factory

import "errors"

func NewPublication(pubType string, name string, pg int, pub string) (IPublication, error) {
	switch pubType {
	case "magazine":
		return createMagazine(name, pg, pub), nil
	case "newspaper":
		return createNewsPaper(name, pg, pub), nil
	}

	return nil, errors.New("no such type of publication")
}

const (
	PublicationMagazine  = "magazine"
	PublicationNewspaper = "newspaper"
)
