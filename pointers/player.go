package main

// unsigned can only be positive
import "fmt"

type Player struct {
	name string
	health, maxHealth uint
	energy, maxEnergy uint
}

func (p *Player) add(health uint) {
	p.health += health
	if p.health > p.maxHealth {
		p.health = p.maxHealth
	}
	fmt.Println(p.name, "add", health, "health -> ", p.health)
}

func (p *Player) apply(damage uint) {
	if p.health - damage > p.health {
		p.health = 0
	} else {
		p.health -= damage
	}
	fmt.Println(p.name, "Damage", damage, "->", p.health)
}

func (p *Player) recharge(energy uint) {
	p.energy += energy
	if p.energy > p.maxEnergy {
		p.energy = p.maxEnergy
	}
	fmt.Println(p.name, "add", energy, "energy -> ", p.energy)
}

func (p *Player) consume(energy uint) {
	if p.energy - energy > p.energy {
		p.energy = 0
	} else {
		p.energy -= energy
	}
	fmt.Println(p.name, "Consume", energy, "->", p.energy)
}

func main() {
	p := Player{
		name: "Cecil",
		health: 100,
		maxHealth: 100,
		energy: 50,
		maxEnergy: 50,
	}

	p.apply(99)
	p.add(10)
	p.consume(20)
	p.recharge(5)

	p.consume(100)
}