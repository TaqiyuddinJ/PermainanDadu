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

func main() {
	// Seed untuk generator angka acak
	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))

	var numPlayers, numDice int

	// Input jumlah pemain dan jumlah dadu
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&numPlayers)

	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&numDice)

	// Inisialisasi pemain
	players := make([]Player, numPlayers)
	for i := range players {
		players[i] = Player{id: i + 1, dice: make([]int, numDice)}
	}

	// Permainan dimulai
	fmt.Println("\nPermainan dimulai!")

	// Counter giliran
	turn := 1

	// Loop permainan hingga hanya ada satu pemain tersisa
	for len(players) > 1 {
		fmt.Printf("==================\nGiliran %d lempar dadu:\n", turn)

		// Semua pemain melempar dadunya
		for i := range players {
			fmt.Printf("Pemain #%d (%d): \n", players[i].id, players[i].points)
			for j := range players[i].dice {
				players[i].dice[j] = rollDice(random)
			}
		}

		// Evaluasi lemparan dadu
		for i := range players {
			for j, dice := range players[i].dice {
				switch dice {
				case 1:
					giveDiceToNextPlayer(players, i, j)
				case 6:
					players[i].points++
				case 2, 3:
					players[i].points += 2
				}
			}
		}

		// Hapus pemain yang telah selesai bermain (tidak memiliki dadu)
		players = removeFinishedPlayers(players)
		if len(players) == 1 {
			break
		}

		// Tampilkan hasil setiap lemparan dadu
		displayGameStatus(players)

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

	fmt.Printf("==================\nGame berakhir karena hanya pemain #%d yang memiliki dadu.", winner.id)
	fmt.Printf("Game dimenangkan oleh pemain #%d dengan total poin %d.\n", winner.id, winner.points)
}

// rollDice menghasilkan angka acak antara 1 hingga 6 untuk lemparan dadu
func rollDice(random *rand.Rand) int {
	return random.Intn(6) + 1
}

// giveDiceToNextPlayer memberikan dadu angka 1 kepada pemain selanjutnya
func giveDiceToNextPlayer(players []Player, currentPlayerIndex, diceIndex int) {
	nextPlayerIndex := (currentPlayerIndex + 1) % len(players)
	players[nextPlayerIndex].dice = append(players[nextPlayerIndex].dice, 1)
	players[currentPlayerIndex].dice[diceIndex] = 0 // Remove dice with 1
}

// removeFinishedPlayers menghapus pemain yang tidak memiliki dadu
func removeFinishedPlayers(players []Player) []Player {
	var activePlayers []Player
	for _, player := range players {
		if len(player.dice) > 0 {
			activePlayers = append(activePlayers, player)
		} else {
			fmt.Printf("Pemain #%d (0): _ (Berhenti bermain karena tidak memiliki dadu)\n", player.id)
		}
	}
	return activePlayers
}

// displayGameStatus menampilkan status setiap pemain setelah lemparan dadu
func displayGameStatus(players []Player) {
	fmt.Println("Setelah evaluasi:")
	for _, player := range players {
		fmt.Printf("Pemain #%d (%d): \n", player.id, player.points)
		for _, dice := range player.dice {
			if dice > 0 && dice < 7 {
				fmt.Printf("%d, ", dice)
			}
		}
		fmt.Println()
	}
}
