package main

import (
	"bytes"
	"errors"
	"fmt"
)

/*
import (
	"fmt"
    "time"
    "math/rand"
	"example.com/go-demo-1/mascot"
	"rsc.io/quote"
)
*/

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("")
	}
	row := (char - 'A' + 1) + (char - 'A')
	col := row
	buf := make([][]byte, row)
	for i := range buf {
		buf[i] = make([]byte, col)
	}
	c := byte('A')
	for i, j, k, l := byte(0), byte(col-1), col/2, col/2; i < row/2+1; i, j, k, l, c = i+1, j-1, k-1, l+1, c+1 {
		init_line := bytes.Repeat([]byte{' '}, int(col))
		buf[i], buf[j] = init_line, init_line
		buf[i][k], buf[i][l], buf[j][k], buf[j][l] = c, c, c, c
	}
	return string(bytes.Join(buf, []byte{'\n'})), nil
}

func main() {
	fmt.Println(Gen('G'))

	/*
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
	*/
}
