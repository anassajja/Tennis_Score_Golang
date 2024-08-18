package main

import "fmt"

func affichage_score(points int) (score string) {
	switch points {
	case 0:
		score = "0"
	case 1:
		score = "15"
	case 2:
		score = "30"
	case 3:
		score = "40"
	case 4:
		score = "A"
	default:
		return "erreur"
	}
	return score
}

func simulation_jeu() int {
	var joueur, points1, points2, jeu int
	// or jeu == 0
	for jeu != 1 {
		fmt.Print("Quel joueur gagne le point ? ")
		fmt.Scanln(&joueur)
		switch joueur {
		case 1:
			points1++
		case 2:
			points2++
		default:
			fmt.Println("Erreur, on a seulement 2 joueurs !")
		}
		switch {
		// cas normal !
		case ((points1 >= 0 && points1 <= 3) && (points2 >= 0 && points2 <= 3)) || (points1 == 4 && points2 == 3) || (points1 == 3 && points2 == 4):
			fmt.Println("le score du jeu :")
			fmt.Println(affichage_score(points1), "-", affichage_score(points2))
		// cas egalite d'avantage
		case (points1 == 4 && points2 == 4):
			fmt.Println("le score du jeu :")
			fmt.Println(affichage_score(3), "-", affichage_score(3))
			points1 = 3
			points2 = 3
		//Victoire jeu cas
		case (points1 == 4 && (points2 >= 0 && points2 <= 2)) || (points1 == 5 && points2 == 3):
			fmt.Println("le score du jeu :")
			fmt.Println(affichage_score(points1), "-", affichage_score(points2))
			joueur = 1
			jeu = 1
		case (points2 == 4 && (points1 >= 0 && points1 <= 2)) || (points2 == 5 && points2 == 3):
			fmt.Println("le score du jeu :")
			fmt.Println(affichage_score(points1), "-", affichage_score(points2))
			joueur = 2
			jeu = 1
		}
	}
	fmt.Print("Le gagnant de ce jeu est le joueur numero : ")
	return joueur
}

func simulation_tiebreak() int {
	var set, joueur, points1, points2 int
	// or set == 0
	for set != 1 {
		fmt.Print("Quel joueur gagne le point ? ")
		fmt.Scanln(&joueur)
		switch joueur {
		case 1:
			points1++

		case 2:
			points2++

		default:
			fmt.Println("Erreur, on a seulement deux joueurs !")
		}
		switch {
		//Victoire tiebreak cas
		case (points1 == 7 && (points2 >= 0 && points2 <= 5)) || (points1 >= 7 && points2 == points1-2):
			joueur = 1
			set = 1
			fmt.Println("le score du jeu decisif :")
			fmt.Println(points1, "-", points2)
		case ((points1 >= 0 && points1 <= 5) && points2 == 7) || (points2 >= 7 && points1 == points2-2):
			joueur = 2
			set = 1
			fmt.Println("le score du jeu decisif :")
			fmt.Println(points1, "-", points2)
		//cas normal !
		case (points1 >= 0 && points2 >= 0):
			fmt.Println("le score du jeu decisif :")
			fmt.Println(points1, "-", points2)
		}
	}
	fmt.Print("Le gagnant du set est le joueur numero : ")
	return joueur
}

func simulation_set1() int {
	var set, joueur, jeup1, jeup2 int
	// or set == 0
	for set != 1 {
		switch {
		case simulation_jeu() == 1:
			joueur = 1
			jeup1++
			fmt.Println(joueur)
		case simulation_jeu() == 2:
			joueur = 2
			jeup2++
			fmt.Println(joueur)
		}
		switch {
		// Victoire set cas
		case (jeup1 == 6 && (jeup2 >= 0 && jeup2 <= 4)) || (jeup1 >= 6 && jeup2 == jeup1-2):
			joueur = 1
			set = 1
			fmt.Println("le score du set :")
			fmt.Println(jeup1, "-", jeup2)
		case (jeup2 == 6 && (jeup1 >= 0 && jeup1 <= 4)) || (jeup2 >= 6 && jeup1 == jeup2-2):
			joueur = 2
			set = 1
			fmt.Println("le score du set :")
			fmt.Println(jeup1, "-", jeup2)
			//Normale cas !
		case (jeup1 >= 0 && jeup2 >= 0):
			fmt.Println("le score du set :")
			fmt.Println(jeup1, "-", jeup2)
		}
		if jeup1 == 6 && jeup2 == 6 {
			simulation_tiebreak()
			set = 1
		}
	}
	fmt.Print("Le gagnant du set est le joueur numero : ")
	return joueur
}

func simulation_set2() int {
	var set, joueur, jeup1, jeup2 int
	// or set == 0
	for set != 1 {
		switch {
		case simulation_jeu() == 1:
			joueur = 1
			jeup1++
			fmt.Println(joueur)
		case simulation_jeu() == 2:
			joueur = 2
			jeup2++
			fmt.Println(joueur)
		}
		switch {
		// Victoire set cas
		case (jeup1 == 6 && (jeup2 >= 0 && jeup2 <= 4)) || (jeup1 >= 6 && jeup2 == jeup1-2):
			joueur = 1
			set = 1
			fmt.Println("le score du set :")
			fmt.Println(jeup1, "-", jeup2)
		case (jeup2 == 6 && (jeup1 >= 0 && jeup1 <= 4)) || (jeup2 >= 6 && jeup1 == jeup2-2):
			joueur = 2
			set = 1
			fmt.Println("le score du set :")
			fmt.Println(jeup1, "-", jeup2)
		//Normale cas !
		case (jeup1 >= 0 && jeup2 >= 0):
			fmt.Println("le score du set :")
			fmt.Println(jeup1, "-", jeup2)
			//Special case
		case (jeup1 == 6 && jeup2 == 6):
			fmt.Println("le score du set :")
			fmt.Println(jeup1, "-", jeup2)
		case (jeup1 > 6 && jeup2 == jeup1-2):
			joueur = 1
			set = 1
			fmt.Println("le score du set :")
			fmt.Println(jeup1, "-", jeup2)
		case (jeup2 > 6 && jeup1 == jeup2-2):
			joueur = 2
			set = 1
			fmt.Println("le score du set :")
			fmt.Println(jeup1, "-", jeup2)
		}
	}
	fmt.Print("Le gagnant du set est le joueur numero : ")
	return joueur
}

func simulation_match() int {
	var match, joueur, set1, set2 int
	// or match == 0
	for match != 1 {
		switch {
		case simulation_set1() == 1:
			joueur = 1
			set1++
			fmt.Println(joueur)

		case simulation_set1() == 2:
			joueur = 2
			set2++
			fmt.Println(joueur)
		}
		switch {
		// Victoire match cas
		case (set1 == 2 && (set2 >= 0 && set2 < 2)):
			joueur = 1
			match = 1
			fmt.Println("le score du match :")
			fmt.Println(set1, "-", set2)
		case (set2 == 2 && (set1 >= 0 && set1 < 2)):
			joueur = 2
			match = 1
			fmt.Println("le score du match :")
			fmt.Println(set1, "-", set2)
		// Cas normal !
		case (set1 >= 0 && set2 >= 0):
			fmt.Println("le score du match :")
			fmt.Println(set1, "-", set2)
		}
	}
	fmt.Print("Le gagnant du match est le joueur numero : ")
	return joueur
}

func simulation_match1() int {
	var match, joueur, set1, set2 int
	for match != 1 {
		switch {
		case simulation_set2() == 1:
			joueur = 1
			set1++
			fmt.Println(joueur)
		case simulation_set2() == 2:
			joueur = 2
			set2++
			fmt.Println(joueur)
		}
		switch {
		// Victoire match cas
		case (set1 == 3 && (set2 >= 0 && set2 <= 2)):
			joueur = 1
			match = 1
			fmt.Println("le score du match :")
			fmt.Println(set1, "-", set2)
		case (set2 == 3 && (set1 >= 0 && set1 <= 2)):
			joueur = 2
			match = 1
			fmt.Println("le score du match :")
			fmt.Println(set1, "-", set2)
		// Cas normal !
		case (set1 >= 0 && set2 >= 0):
			fmt.Println("le score du match :")
			fmt.Println(set1, "-", set2)
		}
	}
	fmt.Print("Le gagnant du match est le joueur numero : ")
	return joueur
}

func simulation_match2() int {
	var match, joueur, set1, set2 int
	for match != 1 {
		switch {
		case simulation_set2() == 1:
			joueur = 1
			set1++
			fmt.Println(joueur)
		case simulation_set2() == 2:
			joueur = 2
			set2++
			fmt.Println(joueur)
		}
		switch {
		// Victoire match cas
		case (set1 == 2 && (set2 >= 0 && set2 < 2)):
			joueur = 1
			match = 1
			fmt.Println("le score du match :")
			fmt.Println(set1, "-", set2)
		case (set2 == 2 && (set1 >= 0 && set1 < 2)):
			joueur = 2
			match = 1
			fmt.Println("le score du match :")
			fmt.Println(set1, "-", set2)
		// Cas normal !
		case (set1 >= 0 && set2 >= 0):
			fmt.Println("le score du match :")
			fmt.Println(set1, "-", set2)
		}
	}
	fmt.Print("Le gagnant du match est le joueur numero : ")
	return joueur
}

func main() {
	var n, points int
	var choix1, choix2 int
	fmt.Println("Menu fonctions :")
	fmt.Println("- affichage_score fonction : 1")
	fmt.Println("- simulation_jeu fonction : 2")
	fmt.Println("- simulation_tiebreak fonction : 3")
	fmt.Println("- simulation_set fonction : 4")
	fmt.Println("- simulation_match fonction : 5")
	fmt.Print("Choisir la fonction : ")
	fmt.Scan(&n)
	switch {
	case n == 1:
		n = 1
		fmt.Print("Enter a number of points : ")
		fmt.Scan(&points)
		fmt.Println(affichage_score(points))
	case n == 2:
		n = 2
		fmt.Scanln()
		fmt.Println(simulation_jeu())
	case n == 3:
		n = 3
		fmt.Scanln()
		fmt.Println(simulation_tiebreak())
	case n == 4:
		n = 4
		fmt.Scanln()
		fmt.Println(simulation_set1())
	case n == 5:
		n = 5
		fmt.Scanln()
		fmt.Print("Veuillez choisir s'il s'agit d'un match Grand Chelem (N.B: 1 pour oui et 0 pour non ) : ")
		fmt.Scan(&choix1)
		switch {
		case choix1 == 1:
			fmt.Print("Veuillez choisir s'il s'agit d'un match des hommes ou bien des femmes (N.B: 1 pour homme et 2 pour femme ) : ")
			fmt.Scan(&choix2)
			switch {
			case choix2 == 1:
				n = 5
				fmt.Scanln()
				fmt.Println(simulation_match1())
			case choix2 == 2:
				n = 5
				fmt.Scanln()
				fmt.Println(simulation_match2())
			default:
				fmt.Println("Erreur !")
			}
		case choix1 == 0:
			n = 5
			fmt.Scanln()
			fmt.Println(simulation_match())
		default:
			fmt.Println("Erreur!")
		}
	default:
		fmt.Println("Erreur, on a seulement 5 fonctions !")
	}
}
