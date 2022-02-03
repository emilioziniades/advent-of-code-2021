package day21

import (
	"fmt"
	"regexp"
	"strconv"
)

type player struct {
	id       int
	position int
	score    int
	winner   bool
}

func (p *player) move(n int) {
	p.position = (p.position-1+n)%10 + 1 // positions from 1 - 10
	p.score += p.position
}

type dice struct {
	rolls   int
	current int
}

func (d *dice) roll() int {
	d.rolls += 1
	d.current = d.current%100 + 1
	return d.current
}

func Play(input []string) int {
	dice := dice{}
	players := setupPlayers(input)

game:
	for {
		for i := 1; i <= len(players); i++ {
			n1, n2, n3 := dice.roll(), dice.roll(), dice.roll()
			players[i].move(n1 + n2 + n3)
			//			fmt.Printf("player %d rolled %d + %d + %d = %d to position %d for a total score of %d\n", i, n1, n2, n3, n1+n2+n3, players[i].position, players[i].score)

			if players[i].score >= 1000 {
				players[i].winner = true
				break game
			}
		}
	}

	for _, player := range players {
		if !player.winner {
			return player.score * dice.rolls
		}
	}
	return 0
}

func PlayDirac(input []string) int {
	players := setupPlayers(input)
	// dereference pointer so that copies can be passed to recursive function
	player1 := *players[1]
	player2 := *players[2]
	winCount := make(map[int]int)

	var recursivePlay func(player, player)
	recursivePlay = func(p1, p2 player) {
		// 3 * 3 * 3 = 27 possible universes for each player turn
		for n1 := 1; n1 <= 3; n1++ {
			for n2 := 1; n2 <= 3; n2++ {
				for n3 := 1; n3 <= 3; n3++ {
					//					fmt.Printf("%p\n", &p1)
					fmt.Println(p1)
					// player 1 moves
					p1.move(n1 + n2 + n3)
					if p1.score >= 21 {
						winCount[1] += 1
						return
					}
					for n4 := 1; n4 <= 3; n4++ {
						for n5 := 1; n5 <= 3; n5++ {
							for n6 := 1; n6 <= 3; n6++ {
								fmt.Println(p2)
								// player 2 moves
								p2.move(n4 + n5 + n6)
								if p2.score >= 21 {
									winCount[2] += 1
									return
								} else {
									recursivePlay(p1, p2)
								}
							}
						}
					}
				}
			}
		}
	}

	recursivePlay(player1, player2)

	fmt.Println(winCount)
	return 0

	//game:
	//	for {
	//	}
}

func setupPlayers(input []string) map[int]*player {
	re := regexp.MustCompile(`Player\s(\d)\sstarting\sposition:\s(\d)`)
	players := make(map[int]*player)
	for _, line := range input {
		match := re.FindStringSubmatch(line)
		id, _ := strconv.Atoi(match[1])
		startPos, _ := strconv.Atoi(match[2])
		players[id] = &player{id: id, position: startPos, score: 0}
	}
	return players
}
