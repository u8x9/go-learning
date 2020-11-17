package extension_test

import (
	"fmt"
	"testing"
)

type pet struct{}

func (p *pet) speak() {
	fmt.Println("...")
}
func (p *pet) speakTo(host string) {
	p.speak()
	fmt.Println(" ", host)
}

type dog struct {
	*pet
}

func (d *dog) speak() {
	fmt.Println("汪汪汪")
}

func TestDog(t *testing.T) {
	d := new(dog)
	d.speak()
	d.speakTo("foobar")
}
