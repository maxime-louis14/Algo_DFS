package main

import "fmt"

func main() {
	// Créer une carte de 30x30 initialisée avec des valeurs 0
	carte := make([][]int, 30)
	for i := range carte {
		carte[i] = make([]int, 30)
	}

	// Définir les variables
	exit := 2
	nid := 3
	mur := 1 // Valeur pour représenter un mur

	// Placer une bordure de murs autour de la carte
	for i := 0; i < 30; i++ {
		carte[0][i] = mur  // Mur en haut
		carte[29][i] = mur // Mur en bas
		carte[i][0] = mur  // Mur à gauche
		carte[i][29] = mur // Mur à droite
	}

	// Définir les positions (par exemple)
	exitX, exitY := 16, 5 // Position pour foob
	nidX, nidY := 7, 26   // Position pour nid

	// Placer les variables sur la carte
	carte[exitX][exitY] = exit
	carte[nidX][nidY] = nid

	// Position de départ
	startX, startY := 7, 26

	// Fonction de vérification de parcours
	if canReachExit(carte, startX, startY, exitX, exitY) {
		fmt.Println("Chemin trouvé jusqu'à l'`exit`!")
	} else {
		fmt.Println("Pas de chemin possible vers l'`exit`.")
	}

	// Afficher la carte après la recherche pour voir les positions visitées
	printCarte(carte)
}

// Fonction de parcours pour vérifier l'accès à l'`exit`
func canReachExit(carte [][]int, x, y, exitX, exitY int) bool {
	// Si la position est hors limites ou un mur, retourner `false`
	if x < 0 || y < 0 || x >= len(carte) || y >= len(carte[0]) || carte[x][y] == 1 {
		return false
	}

	// Si on a atteint l'`exit`, retourner `true`
	if x == exitX && y == exitY {
		return true
	}

	// Marquer la position actuelle comme visitée
	carte[x][y] = 1

	// Déplacements : droite, gauche, bas, haut
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// Essayer chaque direction
	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]
		if canReachExit(carte, newX, newY, exitX, exitY) {
			return true
		}
	}

	return false
}

// Fonction pour afficher la carte dans la console
func printCarte(carte [][]int) {
	for _, ligne := range carte {
		for _, caseVal := range ligne {
			fmt.Print(caseVal, " ")
		}
		fmt.Println() // Nouvelle ligne après chaque ligne de la carte
	}
}
