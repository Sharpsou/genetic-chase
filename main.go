package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

const (
	gridSize    = 10 // taille de la grille
	numElements = 2  // nombre d'éléments à créer
	radarSize   = 18
)

// structure pour un élément de la grille
type element struct {
	x, y           int
	char           rune
	speedX, speedY int
	radar          [radarSize]float64
	//neurone        *Neurone
}

/*type Neurone struct {
	nn *deep.NeuralNetwork
}*/

// grille de jeu
var grid [gridSize][gridSize]rune

func (e *element) updateRadar(elements []*element) {
	for i, _ := range e.radar {
		e.radar[i] = 0
	}

	for _, other := range elements {
		if e == other {
			continue
		}

		dx := other.x - e.x
		dy := other.y - e.y

		angle := math.Atan2(float64(dy), float64(dx))
		if angle < 0 {
			angle += 2 * math.Pi
		}
		sector := int(angle / (math.Pi / 9))

		distance := math.Sqrt(float64(dx*dx + dy*dy))
		if distance < e.radar[sector] || e.radar[sector] == 0 {
			e.radar[sector] = distance
		}
	}
}

// mouvement d'un élément
func (e *element) move() {
	// effacement de la position précédente
	grid[e.y][e.x] = ' '

	// mise à jour de la position
	e.x += e.speedX
	e.y += e.speedY

	// vérifie si l'élément touche un bord et change de direction
	if e.x <= 0 || e.x >= gridSize-1 {
		e.speedX = -e.speedX
	}
	if e.y <= 0 || e.y >= gridSize-1 {
		e.speedY = -e.speedY
	}

	// vérifie si l'élément n'est pas hors du tableau
	if e.x < 0 {
		e.x = 0
	}
	if e.y < 0 {
		e.y = 0
	}
	if e.x > gridSize-1 {
		e.x = gridSize - 1
	}
	if e.y > gridSize-1 {
		e.y = gridSize - 1
	}

	// mise à jour de la nouvelle position de l'élément dans la grille
	grid[e.y][e.x] = e.char
}

// Initialisation des éléments
func initElements() []*element {
	var elements []*element
	rand.Seed(time.Now().UnixNano())

	// création des éléments
	for i := 0; i < numElements; i++ {
		e := &element{
			x:      rand.Intn(gridSize),
			y:      rand.Intn(gridSize),
			char:   rune(rand.Intn(26) + 'A'),
			speedX: rand.Intn(3) - 1,
			speedY: rand.Intn(3) - 1,
		}
		fmt.Println(e)
		if e.speedX == 0 && e.speedY == 0 {
			e.speedX = 1
		}
		elements = append(elements, e)
		grid[e.y][e.x] = e.char
	}
	return elements
}

// Affichage de la grille
func printGrid() {
	os.Stdout.Write([]byte("\033[H\033[2J"))
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c  ", cell)
		}
		fmt.Println()
	}
}

func main() {
	elements := initElements()

	// boucle de jeu
	for {
		// mouvement des éléments
		for _, e := range elements {
			e.move()
		}
		for _, e := range elements {
			e.updateRadar(elements)
			//fmt.Println(e)
		}
		// Affichage de la grille
		printGrid()

		time.Sleep(time.Second / 20)
	}
}
