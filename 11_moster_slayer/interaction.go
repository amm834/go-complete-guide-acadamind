package main

import (
	"fmt"
	"math/rand"
	"time"
)

var currentRound = 1
var specialRound = SpecialRound

var monsterHeal = MonsterHeal
var playerHeal = PlayerHeal
var monsterDamage = MonsterDamage

var playerDefaultMaxDamage = PlayerAttackDeamge

func initGame() {
	showWelcomeScreen()
	showStatics(monsterHeal, playerHeal)
	playing := isPlaying()

	for playing {
		showBasicActions()
		userSelection := getUserInput("Select Action: ")
		if userSelection == "1" {
			if monsterHeal <= 0 {
				showWinner()
				break
			}
			attackTheMonster()
		} else if userSelection == "2" {
			healThePlayer()
		}
		playing = isPlaying()
	}
}

func showWinner() {
	fmt.Printf(`
	===================================
			YOU WIN THE MONSTER
	===================================
`)
	return
}

func attackTheMonster() {

	damage := getRandomBetweenRange(playerDefaultMaxDamage)

	currentRound += 1
	monsterHeal -= damage
	playerHeal -= attackThePlayer()
	if isPlaying() == false {
		showGameOver()
		return
	}

	showStatics(monsterHeal, playerHeal)
}

func healThePlayer() {
	playerHeal += getRandomBetweenRange(10)
}

func getRandomBetweenRange(max int) int {
	rand.Seed(time.Now().UnixNano())
	damage := rand.Intn(max)
	return damage
}

func attackThePlayer() int {
	damage := getRandomBetweenRange(monsterDamage)
	return damage
}

func isPlaying() bool {
	if playerHeal <= 0 {
		return false
	}
	return true
}
