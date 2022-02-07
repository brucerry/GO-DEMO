package main

import (
	"fmt"
	"math/rand"
	"time"

	"example.com/go-demo-1/mascot"
	"rsc.io/quote"
)

func main() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Opt())

	type Person struct {
		name string
		age  int
	}

	// our Team struct
	type Team struct {
		name    string
		players [2]Person
	}

	// declaring an empty 'Team'
	var myTeam Team
	fmt.Println(myTeam)

	myPlayers := [...]Person{{name: "Forrest", age: 18}, {name: "Gordon", age: 22}}
	// declaring a team with players
	celtic := Team{name: "Celtic FC", players: myPlayers}
	fmt.Println(celtic)

	fmt.Println("The time is", time.Now())

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random number is", rand.Intn(100))
}
