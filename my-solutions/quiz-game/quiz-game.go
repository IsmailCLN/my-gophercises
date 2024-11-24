package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filePath := "problems.csv"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Dosya açılamadı:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Dosya okunamadı:", err)
		return
	}

	for i := 0; i < len(rows); i++ {
		fmt.Printf("Soru %d: %s = ?\n", i+1, rows[i][0])

		var userAnswer int
		_, err := fmt.Scanf("%d", &userAnswer)
		fmt.Scanln()
		if err != nil {
			fmt.Printf("Cevapta hata: %s\n", err)
			break
		}

		answer, err := strconv.Atoi(rows[i][1])
		if err != nil {
			fmt.Printf("Cevap formatında hata: %s", err)
			break
		}

		if userAnswer == answer {
			fmt.Println("Doğru cevap!")
		} else {
			fmt.Printf("Yanlış cevap! Doğru cevap: %d, Puanınız: %d\n", answer, i)
			break
		}
	}
}
