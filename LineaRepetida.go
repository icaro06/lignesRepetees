/*Recherche de lignes répétées sans tenir compte des commentaires  A.Villanueva*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*Renvoie true si l'élément existe dans la liste.*/
func exists(elem string, lista []string) bool {
	if elem == "" || strings.Contains(elem, "/*") {
		return false
	}
	for _, valor := range lista {
		if valor == elem {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println("Recherche de lignes répétées sans tenir compte des commentaires")
	if len(os.Args) < 2 {

		fmt.Println("Erreur args  ex:", os.Args[0], " fichier.txt")
		os.Exit(1)
	}
	filename := os.Args[1]

	// ouvrir le fichier
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	lines := []string{} //Créer une liste avec les lignes lues
	line := 0           //ligne actuelle
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line++
		text := scanner.Text()
		text = strings.ReplaceAll(text, " ", "")
		text = strings.ReplaceAll(text, "\t", "")
		if exists(text, lines) {
			fmt.Printf("numéro de ligne %d \t %s \n", line, text)
		}
		lines = append(lines, text)
	}
}
