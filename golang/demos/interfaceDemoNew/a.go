// Package main provides ...
package newDir

import "fmt"

type Humaner interface {
	FillEnergy(n uint)
	ConsumeEnergy(n uint) error
	GetEnergy() uint
}

type bodyChan struct {
	chanType int
	content  interface{}
}

type Body struct {
	Energy uint
}

func NewBody(life uint) *Body {
	return &Body{Energy: life}
}

func (b *Body) ConsumeEnergy(n uint) error {
	if b.Energy >= n {
		b.Energy -= n
		return nil
	}
	return fmt.Errorf("jack is weak")
}

func (b *Body) FillEnergy(n uint) {
	b.Energy += n
}

func (b *Body) GetEnergy() uint {
	return b.Energy
}

type Person struct {
	body       Humaner
	Name       string `json:"Name"`
	energyLock chan bodyChan
	energy     chan uint
	err        chan error
}

func NewPerson(h Humaner, n string) *Person {
	tp := &Person{Name: n,
		energyLock: make(chan bodyChan),
		energy:     make(chan uint),
		err:        make(chan error),
		body:       h,
	}
	tp.listenAllChannel()
	return tp
}
func (p *Person) listenAllChannel() {
	go func() {
		for {
			select {
			case n := <-p.energyLock:
				switch n.chanType {
				case 1:
					p.body.FillEnergy(n.content.(uint))
				case 2:
					p.err <- p.body.ConsumeEnergy(n.content.(uint))
				case 9:
					p.energy <- p.body.GetEnergy()
				}
			}
		}
	}()
}

func (p *Person) DoFillEnergy(n uint) {
	println(p.Name, "'s power up ", n)
	p.energyLock <- bodyChan{chanType: 1, content: n}
}

func (p *Person) DoConsumeEnergy(n uint) error {
	println(p.Name, "'s power down ", n)
	p.energyLock <- bodyChan{chanType: 2, content: n}
	return <-p.err
}

func (p *Person) DoGetEnergy() uint {
	// p.energyLock <- bodyChan{chanType: 9, content: nil}
	p.energyLock <- bodyChan{chanType: 9}
	return <-p.energy
}

func (p *Person) AllMember() {
	fmt.Printf("p: %+v\n", p.body)
}
