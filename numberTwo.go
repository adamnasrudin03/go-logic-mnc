package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	lembar100Ribu = 100000
	lembar50Ribu  = 50000
	lembar20Ribu  = 20000
	lembar10Ribu  = 10000
	lembar5Ribu   = 5000
	lembar2Ribu   = 2000
	lembar1Ribu   = 1000
	koin500Rupiah = 500
	koin200Rupiah = 200
	koin100Rupiah = 100
	lembar        = "lembar"
	koin          = "koin"
)

var (
	pecahan = []int64{
		lembar100Ribu,
		lembar50Ribu,
		lembar20Ribu,
		lembar10Ribu,
		lembar5Ribu,
		lembar2Ribu,
		lembar1Ribu,
		koin500Rupiah,
		koin200Rupiah,
		koin100Rupiah,
	}
	satuanPecahan = map[int64]string{
		lembar100Ribu: lembar,
		lembar50Ribu:  lembar,
		lembar20Ribu:  lembar,
		lembar10Ribu:  lembar,
		lembar5Ribu:   lembar,
		lembar2Ribu:   lembar,
		lembar1Ribu:   lembar,
		koin500Rupiah: koin,
		koin200Rupiah: koin,
		koin100Rupiah: koin,
	}
)

func main() {
	var (
		input        string
		totalBelanja int64
		totalBayar   int64
		err          error
	)

	fmt.Println("Input:")
	fmt.Print("Total belanja seorang customer: Rp ")
	fmt.Scan(&input)
	totalBelanja, err = strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println("Hanya menerima bilangan bulat")
		return
	}

	fmt.Print("Pembeli membayar: Rp ")
	fmt.Scan(&input)
	totalBayar, err = strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println("Hanya menerima bilangan bulat")
		return
	}

	kembalian := totalBayar - totalBelanja
	if kembalian < 0 {
		fmt.Println("False, kurang bayar")
		return
	}

	splitKembalian := strings.Split(fmt.Sprintf("%v", kembalian), "")
	iEnd := len(splitKembalian) - 1
	for i, _ := range splitKembalian {
		if i == iEnd || i == iEnd-1 {
			splitKembalian[i] = "0"
		}
	}

	roundKembalian, err := strconv.ParseInt(strings.Join(splitKembalian, ""), 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Println("\nOutput:")
	fmt.Println("Kembalian yang harus diberikan kasir:", kembalian)
	fmt.Println("dibulatkan menjadi", roundKembalian)
	fmt.Println("\nPecahan uang:")

	sisa := roundKembalian
	for _, v := range pecahan {
		temp := sisa / v
		if temp > 0 {
			fmt.Printf("%v %v %v\n", temp, satuanPecahan[v], v)
			sisa = sisa - (temp * v)
		}
	}
}
