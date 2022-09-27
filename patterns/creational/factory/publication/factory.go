package publication

import "errors"

func NewPublication(pubType string, name string, pg int, pub string) (IPublication, error) {
	switch pubType {
	case Magazine:
		return createMagazine(name, pg, pub), nil
	case Newspaper:
		return createNewspaper(name, pg, pub), nil
	}

	return nil, errors.New("no such type of publication")
}

const (
	Magazine  = "magazine"
	Newspaper = "newspaper"
)
