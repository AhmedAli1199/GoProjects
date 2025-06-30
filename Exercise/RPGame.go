package main

import( 
	"fmt"
)

type Item struct{
	Name string
	Type string
}


type Player struct{
	Name string
	Inventory []Item
}

func (p *Player) AddItem(item Item){
	p.Inventory = append(p.Inventory, item)
}

func (p *Player) DropItem(name string){
	for i,v := range p.Inventory{
		if v.Name == name{
			p.Inventory = append(p.Inventory[:i],p.Inventory[:i+1]... )
			return
		}
	}

}


func main(){

	fmt.Println("Welcome to the RP Game!")
	items := []Item{ { "RedDragon","Potion"  }, {"BlueMark","Eatable"} , {"LilyWater","Potion"},    }

	player1 := Player{Name: "Ahmed"}
	player1.AddItem(items[0])
	fmt.Printf("The item added: %v", items[:2])

	
}