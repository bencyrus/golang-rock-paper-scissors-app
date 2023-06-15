package rps

import (
	"math/rand"
	"time"
)

const (
	Rock     = 0
	Paper    = 1
	Scissors = 2
)

type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computerChoice"`
	RoundResult    string `json:"roundResult"`
	PlayerScore    int    `json:"playerScore"`
	ComputerScore  int    `json:"computerScore"`
}

var winMessages = []string{
	"Good job!",
	"Nice one!",
	"You're a winner!",
}

var loseMessages = []string{
	"Better luck next time!",
	"Hard luck!",
	"You lost!",
}

var drawMessages = []string{
	"It's a draw!",
	"Great minds think alike!",
	"Nobody wins, but you can try again!",
}

var playerScore = 0

var computerScore = 0

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	messageInt := rand.Intn(3)
	message := ""

	switch computerValue {
	case Rock:
		computerChoice = "ROCK"
	case Paper:
		computerChoice = "PAPER"
	case Scissors:
		computerChoice = "SCISSORS"
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a draw!"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		message = winMessages[messageInt]
		playerScore++
	} else {
		roundResult = "Computer Wins!"
		message = loseMessages[messageInt]
		computerScore++
	}

	result := Round{
		Message:        message,
		ComputerChoice: computerChoice,
		RoundResult:    roundResult,
		PlayerScore:    playerScore,
		ComputerScore:  computerScore,
	}

	return result
}
