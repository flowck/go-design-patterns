package publication

type IPublication interface {
	setName(name string)
	setPages(pages int)
	setPublisher(publisher string)
	Name() string
	Pages() int
	Publisher() string
}

type publication struct {
	name      string
	pages     int
	publisher string
}

func (p *publication) setName(name string) {
	p.name = name
}

func (p *publication) setPages(pages int) {
	p.pages = pages
}

func (p *publication) setPublisher(publisher string) {
	p.publisher = publisher
}

func (p *publication) Name() string {
	return p.name
}
func (p *publication) Pages() int {
	return p.pages
}

func (p *publication) Publisher() string {
	return p.publisher
}
