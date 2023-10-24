package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	NumPlayers = 4
	NumPawns   = 4
	MaxSquare  = 40
	Start      = -1
	Finish     = MaxSquare + 1
	Obstacle   = -2
)

type Board struct {
	positions [NumPlayers][NumPawns]int
	squares   [MaxSquare]int
	mutex     sync.Mutex
	gameOver  bool
}

type Player struct {
	id      int
	board   *Board
	channel chan int
}

func (p *Player) hasWon() bool {
	for i := 0; i < NumPawns; i++ {
		if p.board.positions[p.id][i] != Finish {
			return false
		}
	}
	return true
}

func (p *Player) play(wg *sync.WaitGroup) {
	defer wg.Done()

	for !p.board.gameOver {
		// Lanzar el dado
		dice := rand.Intn(6) + 1

		p.board.mutex.Lock()

		// Elegir un peón para mover
		pawn := rand.Intn(NumPawns)

		// Si el peón está en el inicio y el dado es 6, mover el peón
		if p.board.positions[p.id][pawn] == Start && dice == 6 {
			p.board.positions[p.id][pawn] = 0
		} else if p.board.positions[p.id][pawn] != Start && p.board.positions[p.id][pawn] != Finish {
			p.board.positions[p.id][pawn] += dice

			// Verificar si el peón ha llegado a la meta
			if p.board.positions[p.id][pawn] > MaxSquare {
				p.board.positions[p.id][pawn] = Finish
			}
		}

		// Imprimir la posición actual del peón
		fmt.Printf("Jugador %d, Peón %d, Posición %d\n", p.id, pawn, p.board.positions[p.id][pawn])

		// Verificar si el jugador ha ganado
		if p.hasWon() {
			fmt.Printf("¡El Jugador %d ha ganado!\n", p.id)
			p.board.gameOver = true
		}

		p.board.mutex.Unlock()

		// Esperar un tiempo antes del siguiente turno
		time.Sleep(1 * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	board := &Board{}

	// Inicializar las posiciones de las fichas en el inicio y algunos obstáculos en el tablero
	for i := 0; i < NumPlayers; i++ {
		for j := 0; j < NumPawns; j++ {
			board.positions[i][j] = Start
		}
	}
	for i := 0; i < MaxSquare; i += 10 {
		board.squares[i] = Obstacle
	}

	// Inicializar jugadores y canales
	players := make([]Player, NumPlayers)
	for i := range players {
		players[i] = Player{
			id:      i,
			board:   board,
			channel: make(chan int),
		}
		wg.Add(1)
		go players[i].play(&wg)
	}

	wg.Wait()
}