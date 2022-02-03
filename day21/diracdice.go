package day21

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/emilioziniades/adventofcode2021/util"
)

type player struct {
	id       int
	position int
	score    int
	winner   bool
}

func (p *player) move(n int) {
	p.position = (p.position+n-1)%10 + 1 // positions from 1 - 10
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

type state struct {
	player1, player2 player
	player1turn      bool
}

func (s state) String() string {
	var turn int
	if s.player1turn {
		turn = 1
	} else {
		turn = 2
	}
	format := "( p%d position %d, score %d // p%d position %d, score %d // player %d turn )"
	return fmt.Sprintf(format, s.player1.id, s.player1.position, s.player1.score, s.player2.id, s.player2.position, s.player2.score, turn)
}

func Play(input []string) int {
	dice := dice{}
	players := setupPlayers(input)

game:
	for {
		for i := 1; i <= len(players); i++ {
			n1, n2, n3 := dice.roll(), dice.roll(), dice.roll()
			players[i].move(n1 + n2 + n3)
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
	universes := make(map[state]int)
	winCount := make(map[int]int)
	initialState := state{player1: *players[1], player2: *players[2], player1turn: true}
	universes[initialState] = 1

	// whilst there are still uncompleted universes
	for len(universes) > 0 {
		newuniverses := util.CopyMap(universes)
		// iterate over each state
		for s, c := range universes {
			delete(newuniverses, s)

			// end if there is a winner
			if s.player1.score >= 21 {
				winCount[1] += c
				continue
			} else if s.player2.score >= 21 {
				winCount[2] += c
				continue
			}
			// create 27 copies of current state, 1 for each possible combination of dice rolls
			for n1 := 1; n1 <= 3; n1++ {
				for n2 := 1; n2 <= 3; n2++ {
					for n3 := 1; n3 <= 3; n3++ {
						sum := n1 + n2 + n3
						currState := s
						switch currState.player1turn {
						case true:
							currState.player1.move(sum)
							currState.player1turn = false
						case false:
							currState.player2.move(sum)
							currState.player1turn = true
						}
						newuniverses[currState] += c
					}
				}
			}
		}
		universes = util.CopyMap(newuniverses)
	}

	if winCount[1] > winCount[2] {
		return winCount[1]
	} else {
		return winCount[2]
	}

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
