package main

import (
	"fmt"
)

func showWelcomeScreen() {
	fmt.Printf(`
	------------------------------------
		WECOME TO MONSTER SLAYER GAME
	------------------------------------
v%v
`, VERSION)
}

func showBasicActions() {
	fmt.Printf(`
	1) Attack
	2) Heal
`)
}

func showSpecialAction() {
	fmt.Printf(`
	1) Attack
	2) Heal
	3) Special Attack
`)
}

func showStatics(monsterHeal, playerHeal int) {
	fmt.Printf(`
	Monster: %v =======VS======= You: %v
`, monsterHeal, playerHeal)
}

func showGameOver() {
	fmt.Printf(`
=====================================
              YOU LOST
=====================================
`)
}
