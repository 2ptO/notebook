package jsonexample

import (
	"encoding/json"
	"fmt"
)

/*
 * Examples showing the basic use of json package in Go.
 */

type bio struct {
	Name string
	Age  int
}

//Batsman ..exported for JSON usage
type Batsman struct {
	bio
	Runs int
}

func (b Batsman) score() int {
	return b.Runs
}

func (b *Batsman) addRuns(runs int) {
	b.Runs += runs
}

//Bowler2 ...exported for JSON usage.
type Bowler2 struct {
	bio
	Wickets int
}

func (b Bowler2) score() int {
	return b.Wickets
}

//Player ..exported for JSON usage
type Player interface {
	score() int
}

func main() {
	sachin := Batsman{
		bio: bio{
			Name: "sachin",
			Age:  32,
		},
		Runs: 100,
	}

	ashwin := Bowler2{
		bio: bio{
			Name: "ashwin",
			Age:  28,
		},
		Wickets: 5,
	}

	players := []Player{sachin, ashwin}
	fmt.Println(players)

	if playersByteSlice, err := json.Marshal(players); err == nil {
		fmt.Println(string(playersByteSlice))
	} else {
		fmt.Println(err)
	}

	playersJSON := []byte(`[{"Name":"sachin","Age":32,"Runs":100}]`)
	var playersFromJSON []Batsman
	if err := json.Unmarshal(playersJSON, &playersFromJSON); err == nil {
		fmt.Println("Decoded from JSON:", playersFromJSON)
	} else {
		fmt.Println(err)
	}

}
