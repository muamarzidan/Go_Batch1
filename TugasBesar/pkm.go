package main

import "fmt"

func main() {
	var monster, shield int
	var HPBulba, ATKBulba int
	var HPSquirtle, ATKSquirtle int
	var HPPikachu, ATKPikachu int
	var HPGarga, ATKGarga int
	var decision, move string

	HPBulba = 100
	ATKBulba = 15
	HPSquirtle = 250
	ATKSquirtle = 30
	HPPikachu = 400
	ATKPikachu = 60
	HPGarga = 300
	ATKGarga = 40
	shield = 3

	fmt.Println("============== POKEMON BATTLE ARENA ==============")
	fmt.Println("Choose one of this monster to be your alliance: ")
	fmt.Println("1. Bulbasaur (HP: 100, ATK : 15)")
	fmt.Println("2. Squirtle  (HP: 250, ATK : 30)")
	fmt.Println("3. Pikachu   (HP: 400, ATK : 60)")
	fmt.Println("==================================================")
	fmt.Scanln(&monster)
	if monster == 1 {
		fmt.Println("============== POKEMON BATTLE ARENA ==============")
		fmt.Println("Alrightt, Lets Bulbasaur do it's magic")
		fmt.Println("Your opponent will be Garganacl, are you ready? (YES/NO)")
		fmt.Println("==================================================")
		fmt.Scanln(&decision)
		if decision == "YES" {
			fmt.Println("============== POKEMON BATTLE ARENA ==============")
			fmt.Println("============= BULBASAUR VS GARGANACL =============")
			fmt.Println("You can use Shield three times")
			for HPBulba > 0 && HPGarga > 0 {
				fmt.Println("Attack / Shield")
				fmt.Scanln(&move)
				if move == "Attack" {
					HPBulba = HPBulba - ATKGarga
					HPGarga = HPGarga - ATKBulba
					fmt.Println("Bulbasaur HP	:", HPBulba)
					fmt.Println("Garga HP	:", HPGarga)
					fmt.Println("")
				} else if move == "Shield" {
					if shield > 0 {
						shield = shield - 1
						fmt.Println("Shield remaining: ", shield)
						fmt.Println("Bulbasaur HP	:", HPBulba)
						fmt.Println("Garga HP	:", HPGarga)
						fmt.Println("")
					} else {
						fmt.Println("All Shield already used")
					}
				}
				fmt.Println("==================================================")
			}

			if HPBulba <= 0 {
				fmt.Println("		D E F E A T E D	")
				fmt.Println("==================================================")
			}
			if HPGarga <= 0 {
				fmt.Println("		V I C T O R Y")
				fmt.Println("==================================================")
			}
		} else {
			fmt.Println("SEE YOU NEXT GAME")
		}
	} else if monster == 2 {
		fmt.Println("============== POKEMON BATTLE ARENA ==============")
		fmt.Println("Alrightt, Lets Squirtle do it's magic")
		fmt.Println("Your opponent will be Garganacl, are you ready? (YES/NO)")
		fmt.Println("==================================================")
		fmt.Scanln(&decision)
		if decision == "YES" {
			fmt.Println("============== POKEMON BATTLE ARENA ==============")
			fmt.Println("============== SQUIRTLE VS GARGANACL =============")
			fmt.Println("You can use Shield three times")
			for HPSquirtle > 0 && HPGarga > 0 {
				fmt.Println("Attack / Shield")
				fmt.Scanln(&move)
				if move == "Attack" {
					HPSquirtle = HPSquirtle - ATKGarga
					HPGarga = HPGarga - ATKSquirtle
					fmt.Println("Squirtle HP	:", HPSquirtle)
					fmt.Println("Garga HP	:", HPGarga)
					fmt.Println("")
				} else if move == "Shield" {
					if shield > 0 {
						shield = shield - 1
						fmt.Println("Shield remaining: ", shield)
						fmt.Println("Squrirtle HP	:", HPSquirtle)
						fmt.Println("Garga HP	:", HPGarga)
						fmt.Println("")
					} else {
						fmt.Println("All Shield already used")
					}
				}
				fmt.Println("==================================================")
			}

			if HPSquirtle <= 0 {
				fmt.Println("		D E F E A T E D	")
				fmt.Println("==================================================")
			}
			if HPGarga <= 0 {
				fmt.Println("		V I C T O R Y")
				fmt.Println("==================================================")
			}
		} else {
			fmt.Println("SEE YOU NEXT GAME")
		}
	} else if monster == 3 {
		fmt.Println("============== POKEMON BATTLE ARENA ==============")
		fmt.Println("Alrightt, Lets Pikachu do it's magic")
		fmt.Println("Your opponent will be Garganacl, are you ready? (YES/NO)")
		fmt.Println("==================================================")
		fmt.Scanln(&decision)
		if decision == "YES" {
			fmt.Println("============== POKEMON BATTLE ARENA ==============")
			fmt.Println("============== PIKACHU VS GARGANACL ==============")
			fmt.Println("You can use Shield three times")
			for HPPikachu > 0 && HPGarga > 0 {
				fmt.Println("Attack / Shield")
				fmt.Scanln(&move)
				if move == "Attack" {
					HPPikachu = HPPikachu - ATKGarga
					HPGarga = HPGarga - ATKPikachu
					fmt.Println("Pikachu HP	:", HPPikachu)
					fmt.Println("Garga HP	:", HPGarga)
					fmt.Println("")
				} else if move == "Shield" {
					if shield > 0 {
						shield = shield - 1
						fmt.Println("Shield remaining: ", shield)
						fmt.Println("Pikachu HP	:", HPPikachu)
						fmt.Println("Garga HP	:", HPGarga)
						fmt.Println("")
					} else {
						fmt.Println("All Shield already used")
					}
				}
				fmt.Println("==================================================")
			}

			if HPPikachu <= 0 {
				fmt.Println("		D E F E A T E D	")
				fmt.Println("==================================================")
			}
			if HPGarga <= 0 {
				fmt.Println("		V I C T O R Y")
				fmt.Println("==================================================")
			}
		} else {
			fmt.Println("SEE YOU NEXT GAME")
		}
	}
}
