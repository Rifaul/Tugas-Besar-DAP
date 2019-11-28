/*This program code is created by Aulia Rahman Arif
	Student of Telkom University, Informatics Major, Class IF-43-INT
	This program is created to fullfill the requirement of the final project
of Basic Algorithm and Programming
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Creating new types to make it easier to organize
type Player struct {
	pos, score int
	name       string
}

type Ladder struct {
	bottom, top int
}

type Snake struct {
	head, tail int
}

type Crumb struct {
	pos, val int
}

type MOB struct {
	ladders [7]Ladder
	snakes  [10]Snake
	crumbs  [20]Crumb
}

var mobs MOB
var box [106]int // I used 106 for the length of the array, just in case the player overpass the 100 square
var user Player
var rolls int

// Initializing the Ladders randomly
func initLadder() {
	var x, y int

	for i := 0; i < 7; {
		rand.Seed(time.Now().UnixNano())
		x = rand.Intn(79) + 10
		y = rand.Intn(79) + 10
		if x > y {
			for (x - y) < 11 {
				x = rand.Intn(79) + 10
			}
			if box[y] == 0 {
				if box[x] == 0 {
					mobs.ladders[i].bottom = y
					mobs.ladders[i].top = x
					box[x] = 5
					box[y] = 1
					i++
				}
			}
		} else if y > x {
			for (y - x) < 11 {
				y = rand.Intn(79) + 10
			}
			if box[y] == 0 {
				if box[x] == 0 {
					mobs.ladders[i].bottom = x
					mobs.ladders[i].top = y
					box[x] = 1
					box[y] = 5
					i++
				}
			}
		}

	}
}

// Initializing the snakes randomly
func initSnakes() {
	var x, y int

	for i := 0; i < 10; {
		rand.Seed(time.Now().UnixNano())
		x = rand.Intn(64) + 25
		y = rand.Intn(64) + 25
		if x > y {
			for (x - y) < 11 {
				x = rand.Intn(64) + 25
			}
			if box[y] == 0 {
				if box[x] == 0 {
					mobs.snakes[i].tail = y
					mobs.snakes[i].head = x
					box[x] = 2
					box[y] = 5
					i++
				}
			}
		} else if y > x {
			for (y - x) < 11 {
				y = rand.Intn(64) + 25
			}
			if box[y] == 0 {
				if box[x] == 0 {
					mobs.snakes[i].tail = x
					mobs.snakes[i].head = y
					box[x] = 5
					box[y] = 2
					i++
				}
			}
		}

	}
}

// Initializing the crumbs randomly
func initCrumbs() {
	var x, y int

	for i := 0; i < 20; {
		rand.Seed(time.Now().UnixNano())
		x = rand.Intn(100)
		if box[x] == 0 {
			mobs.crumbs[i].pos = x
			y = rand.Intn(10) + 1
			mobs.crumbs[i].val = y
			box[x] = 3
			i++
		}
	}
}

// Initializing the Monsters randomly
func initMonsters() {
	var x int

	for i := 0; i < 10; {
		rand.Seed(time.Now().UnixNano())
		x = rand.Intn(100)
		if box[x] == 0 {
			box[x] = 4
			i++
		}
	}
}

// Initializing all of the snakes and ladders pion/mobs
func initialization() {
	for i := 0; i < 101; i++ {
		box[i] = 0
	}
	box[0] = 5
	initLadder()
	initSnakes()
	initCrumbs()
	initMonsters()
	user.pos = 0
	rolls = 0
	fmt.Println("Ladders: ")
	for i := 0; i < 7; i++ {
		fmt.Printf("|| %v >> %v ||\n", mobs.ladders[i].bottom, mobs.ladders[i].top)
	}

	fmt.Println("Snakes: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("|| %v >> %v ||\n", mobs.snakes[i].head, mobs.snakes[i].tail)
	}
	fmt.Println()
}

//Printing the welcome message to the player
func welcomeMessage(play *bool) {
	var n string

	fmt.Println("+----------------------------------------------------------------+")
	fmt.Println("|              Welcome to Monster Snakes and Ladder              |")
	fmt.Println("+----------------------------------------------------------------+")
	fmt.Scan()
	fmt.Printf("What is your Name? ")
	fmt.Scanln(&user.name)
	fmt.Printf("Okay %s, before we start, let me show you how to play this game\n", user.name)
	fmt.Scanln()
	fmt.Println("++------------------------------------------------------------------------------------++")
	fmt.Println("||                This game is basically the usual snakes and ladder.                 ||")
	fmt.Println("||       There are Snakes that will bring your down and Ladder to bring you up.       ||")
	fmt.Println("||              There are also crumbs, you get points if you collect them             ||")
	fmt.Println("|| But, there are also monsters that will decrese your points by 10% if you meet them ||")
	fmt.Println("++------------------------------------------------------------------------------------++")
	fmt.Printf("\nSo, are you ready to play(Y/N)? ")
	fmt.Scanln(&n)
	if n == "Y" {
		*play = true
	} else {
		*play = false
	}
}

// Used to roll the dice
func rollDice(die *int) {
	fmt.Print("|| Press ENTER to roll the die ||")
	fmt.Scanln()
	rand.Seed(time.Now().UnixNano())
	*die = rand.Intn(6) + 1
	rolls = rolls + 1
	fmt.Printf("You rolled the die %d\n", *die)
}

// Moving the player's position from the bottom ladder to the top
func moveLadder() {
	var i int = 0
	var found bool = false

	for i < 7 && !found {
		if mobs.ladders[i].bottom == user.pos {
			user.pos = mobs.ladders[i].top
			found = true
		}
		i++
	}
	fmt.Printf("You land on a Ladder at %d, You ascend to %d\n", mobs.ladders[i-1].bottom, user.pos)
}

// Moving the player's position from the snake's head to its tail
func moveSnake() {
	var i int = 0
	var found bool = false

	for i < 10 && !found {
		if mobs.snakes[i].head == user.pos {
			user.pos = mobs.snakes[i].tail
			found = true
		}
		i++
	}
	fmt.Printf("You hit a Snake head at %d, You descend to %d\n", mobs.snakes[i-1].head, user.pos)
}

// increasing the score when landing on a crumb
func addScore() {
	var i int = 0
	var found bool = false

	for i < 20 && !found {
		if mobs.crumbs[i].pos == user.pos {
			user.score = user.score + mobs.crumbs[i].val
			found = true
		}
		i++
	}
	fmt.Printf("You land on a crumb square at %d, You received %d point(s)\n", user.pos, mobs.crumbs[i-1].val)
	fmt.Printf("You now have %d point(s)\n", user.score)
}

//decreasing the score when landing on a monster
func decreaseScore() {
	user.score = user.score - (user.score / 10)
	fmt.Printf("You hit a monster square at %d, your point(s) are decresed to %d\n", user.pos, user.score)
}

// Updating the player's position and whether it need interaction
func Update(die int) {
	user.pos = user.pos + die
	if box[user.pos] == 1 {
		moveLadder()
	} else if box[user.pos] == 2 {
		moveSnake()
	} else if box[user.pos] == 3 {
		addScore()
	} else if box[user.pos] == 4 {
		decreaseScore()
	} else if user.pos > 100 {
		temp := user.pos - 100
		user.pos = 100 - temp
	}
	fmt.Printf("Your current position is now at %d\n", user.pos)
}

//Giving congratulations to the player for winning the game
func winMessage(play *bool) {
	fmt.Printf("\n!!!!!!!!!!## Y O U   W I N ##!!!!!!!!!!\n")
	fmt.Printf("Congratulations %s, you reached the 100th square with %d roll(s) of dice and achieved %d point(s)\n", user.name, rolls, user.score)
	fmt.Println()
	*play = false
}

//Giving gratitude to the player for participating in the game
func exitMessage() {
	fmt.Printf("Thank you very much %s, for participating in this game XD\n", user.name)
	fmt.Println("MADE BY :")
	fmt.Println("+---------------------------------+")
	fmt.Println("|    AULIA RAHMAN ARIF WAHYUDI    |")
	fmt.Println("|    IF-43-INT  ||  1301194195    |")
	fmt.Println("+---------------------------------+")
	fmt.Println("| Basic Algorithm and Programming |")
	fmt.Println("|     Informatics Engineering     |")
	fmt.Println("|        Telkom University        |")
	fmt.Println("+---------------------------------+")
}

func main() {
	var die int
	var play bool

	welcomeMessage(&play)
	for play {
		initialization()
		for user.pos < 100 {
			rollDice(&die)
			Update(die)
		}
		winMessage(&play)
	}
	exitMessage()
}
