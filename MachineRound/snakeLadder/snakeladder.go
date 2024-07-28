package snakeLadder

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Game struct {
	Snakes  map[int]int
	Ladder  map[int]int
	Players map[string]int
	Winner  string
}

func newGame(snakes, ladder map[int]int, players []string) *Game {
	playerPosition := make(map[string]int, len(players))
	for _, player := range players {
		playerPosition[player] = 0
	}
	return &Game{Snakes: snakes, Ladder: ladder, Players: playerPosition, Winner: ""}
}
func (g *Game) rollDice() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(6) + 1
}
func (g *Game) movePlayer(player string) error {
	roll := g.rollDice()
	currentPosition := g.Players[player]
	newPosition := currentPosition + roll
	if g.Winner != "" {
		return nil
	}

	defer func() {
		fmt.Println(player + " rolled a " + strconv.Itoa(roll) + " and moved from " + strconv.Itoa(currentPosition) + " to " + strconv.Itoa(newPosition))
	}()

	if newPosition > 100 {
		newPosition = currentPosition
		return nil
	}

	if tail, ok := g.Snakes[newPosition]; ok {
		newPosition = tail
	} else if end, ok := g.Ladder[newPosition]; ok {
		newPosition = end
	}
	g.Players[player] = newPosition

	if newPosition == 100 {
		g.Winner = player
	}
	return nil

}
func (g *Game) playGame() {
	for g.Winner == "" {
		for player, _ := range g.Players {
			_ = g.movePlayer(player)
			if g.Winner != "" {
				fmt.Println(player + " wins the game")
				return
			}
		}
	}
}

func parseInput() (map[int]int, map[int]int, []string) {
	reader := bufio.NewReader(os.Stdin)

	var s, l, p int
	fmt.Scanln(&s)

	snakes := make(map[int]int, s)
	for i := 0; i < s; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		head, _ := strconv.Atoi(parts[0])
		tail, _ := strconv.Atoi(parts[1])
		snakes[head] = tail
	}

	fmt.Scanln(&l)
	ladders := make(map[int]int, l)
	for i := 0; i < l; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		ladders[start] = end
	}

	fmt.Scanln(&p)
	players := make([]string, p)
	for i := 0; i < p; i++ {
		line, _ := reader.ReadString('\n')
		players[i] = strings.TrimSpace(line)
	}
	return snakes, ladders, players
}

func Init() {
	snakes, ladder, players := parseInput()
	game := newGame(snakes, ladder, players)
	game.playGame()
}
