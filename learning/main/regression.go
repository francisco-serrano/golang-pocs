package main

import (
	"encoding/csv"
	"fmt"
	"gonum.org/v1/gonum/stat"
	"log"
	"os"
	"strconv"
)

type xy struct {
	x []float64
	y []float64
}

func main() {
	filename := "path"

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	size := len(records)

	data := xy{
		x: make([]float64, size),
		y: make([]float64, size),
	}

	for i, v := range records {
		if len(v) != 2 {
			fmt.Println("expected two elements")
			continue
		}

		if s, err := strconv.ParseFloat(v[0], 64); err == nil {
			data.y[i] = s
		}

		if s, err := strconv.ParseFloat(v[1], 64); err == nil {
			data.x[i] = s
		}
	}

	b, a := stat.LinearRegression(data.x, data.y, nil, false)

	fmt.Printf("%.4v x + %.4v\n", a, b)
	fmt.Printf("a = %.4v b = %.4v\n", a, b)
}
