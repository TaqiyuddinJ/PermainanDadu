package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Player struct represents a player with its dice and points
type Player struct {
	id     int
	dice   []int
	points int
}

var players []Player
var activePlayers []Player
var numPlayers, numDicePerPlayer int

func main() {
	// Seed untuk generator angka acak
	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))

	// Input jumlah pemain dan jumlah dadu
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&numPlayers)

	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&numDicePerPlayer)

	// Inisialisasi pemain
	players = make([]Player, numPlayers)
	for i := range players {
		players[i] = Player{id: i + 1, dice: make([]int, numDicePerPlayer)}
	}

	// Permainan dimulai
	fmt.Printf("\nPemain = %d, Dadu = %d\n", numPlayers, numDicePerPlayer)
	fmt.Println("==================")

	// Counter giliran
	turn := 1
	activePlayers = players
	// Loop permainan hingga hanya ada satu pemain tersisa
	for len(activePlayers) > 1 {
		fmt.Printf("Giliran %d lempar dadu:\n", turn)

		// Semua pemain melempar dadunya
		for i := range activePlayers {
			fmt.Printf("Pemain #%d (%d): ", activePlayers[i].id, activePlayers[i].points)
			rollDiceForPlayer(random, &activePlayers[i])
			if len(activePlayers[i].dice) > 0 {
				fmt.Println(activePlayers[i].dice)
			}
		}
		// Oper dadu angka 1 ke pemain berikutnya
		for i := range activePlayers {
			activePlayers[i].passDiceToNextPlayer()
		}

		// Evaluasi lemparan dadu
		for i := range activePlayers {
			activePlayers[i].evaluateDice()
		}

		// Hapus pemain yang telah selesai bermain (tidak memiliki dadu)
		activePlayers = removeFinishedPlayers(activePlayers)

		// Tampilkan hasil setelah evaluasi
		displayGameStatus(activePlayers)

		// Tambahkan counter giliran
		turn++
	}

	// Menentukan pemenang
	var winner Player
	for _, player := range players {
		if player.points > winner.points {
			winner = player
		}
	}

	fmt.Printf("==================\nGame berakhir karena hanya pemain #%d yang memiliki dadu.\n", activePlayers[0].id)
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.\n", winner.id)
}

// rollDiceForPlayer melempar dadu untuk seorang pemain
func rollDiceForPlayer(random *rand.Rand, player *Player) {
	for i := range player.dice {
		player.dice[i] = rollDice(random)
	}
}

// evaluateDice mengevaluasi hasil lemparan dadu pemain
func (player *Player) evaluateDice() {
	var newDice []int
	for _, dice := range player.dice {
		switch dice {
		case 1:
			// player.passDiceToNextPlayer()
			continue
		case 6:
			player.points++ // Hitung poin untuk dadu 6
		default:
			newDice = append(newDice, dice)
		}
	}
	player.dice = newDice
	fmt.Println("player", player.id, "dice:", player.dice)
}

// passDiceToNextPlayer memberikan dadu angka 1 kepada pemain selanjutnya
func (player *Player) passDiceToNextPlayer() {
	nextPlayerIndex := player.id % len(activePlayers)
	fmt.Println("next player index to pass:", nextPlayerIndex)
	nextPlayer := &players[nextPlayerIndex]
	nextPlayer.dice = append(nextPlayer.dice, 1)
	fmt.Println("next player dice after pass:", nextPlayer.dice)
}

// removeFinishedPlayers menghapus pemain yang tidak memiliki dadu
func removeFinishedPlayers(players []Player) []Player {
	var activePlayers []Player
	for _, player := range players {
		if len(player.dice) > 0 {
			activePlayers = append(activePlayers, player)
		}
	}
	return activePlayers
}

// displayGameStatus menampilkan status setiap pemain setelah lemparan dadu
func displayGameStatus(players []Player) {
	fmt.Println("==================")
	fmt.Println("Setelah evaluasi:")
	for _, player := range players {
		fmt.Printf("Pemain #%d (%d): ", player.id, player.points)
		if len(player.dice) == 0 {
			fmt.Printf("_ (Berhenti bermain karena tidak memiliki dadu)\n")
		} else {
			for _, dice := range player.dice {
				fmt.Printf("%d,", dice)
			}
			fmt.Println()
		}
	}
	fmt.Println("==================")
}

// rollDice menghasilkan angka acak antara 1 hingga 6 untuk lemparan dadu
func rollDice(random *rand.Rand) int {
	return random.Intn(6) + 1
}
