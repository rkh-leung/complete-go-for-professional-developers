package main

import (
	"fmt"
)

type Player struct {
	Name      string
	Inventory []Item
}

type Item struct {
	Name string
	Type string
}

// func main() {
// 	k := Player{
// 		Name:      "K",
// 		Inventory: []Item{},
// 	}
// 	sword := Item{
// 		Name: "Sword",
// 		Type: "Weapon",
// 	}
// 	potion := Item{
// 		Name: "Red pot",
// 		Type: "potion",
// 	}
//
// 	k.dropItem(sword.Name)
// 	k.pickUpItem(potion)
// 	k.pickUpItem(sword)
// 	k.dropItem(sword.Name)
// 	k.useItem(sword.Name)
// 	k.useItem(potion.Name)
// }

func (p *Player) pickUpItem(i Item) {
	p.Inventory = append(p.Inventory, i)
	fmt.Printf("%s picked up %s!\n", p.Name, i.Name)
}

func (p *Player) dropItem(itemName string) {
	for i, item := range p.Inventory { // doesn't run on empty slice
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Printf("%s dropped %s!\n", p.Name, itemName)
			return
		}
	}
	fmt.Printf("Oops, don't have %s in inventory\n", itemName)
}

func (p *Player) useItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			if item.Type == "potion" {
				fmt.Printf("%s powah!", itemName)
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			} else {
				fmt.Printf("%s used %s!\n", p.Name, itemName)
			}
			return
		}
	}
	fmt.Printf("Oops, don't have %s to use\n", itemName)
}
