package observer

type observable interface {
	Register(obs DataListener)
	Unregister(obs DataListener)
	notifyAll()
}

type DataSubject struct {
	observers []DataListener
	field     string
}

func (ds *DataSubject) Register(obs DataListener) {
	ds.observers = append(ds.observers, obs)
}

func (ds *DataSubject) Unregister(obs DataListener) {

	var newObservers []DataListener

	count := 0
	for _, item := range ds.observers {
		if item.Name != obs.Name {
			newObservers = append(newObservers, item)
		}
		count++
	}

	ds.observers = newObservers
}

func (ds *DataSubject) notifyAll() {
	for _, obs := range ds.observers {
		obs.onUpdate(ds.field)
	}
}

func (ds *DataSubject) ChangeItem(data string) {
	ds.field = data
	ds.notifyAll()
}
