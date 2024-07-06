package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

const (
	formatDate      = "2006-01-02"
	maxCutiBeruntun = 3
)

func roundDownFloatToInt(value float64) int {
	ratio := math.Pow(10, float64(0))
	res := math.Floor(value*ratio) / ratio
	return int(res)
}

func main() {
	var (
		input        string
		startDate    time.Time
		cutiDate     time.Time
		cutiBersama  int
		durationCuti int
		err          error
	)

	fmt.Println("Input: ")
	fmt.Print("* Jumlah Cuti Bersama = ")
	fmt.Scan(&input)
	cutiBersama, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("Jumlah Cuti Bersama Hanya menerima bilangan bulat")
		return
	}

	fmt.Print("* Tanggal join karyawan = ")
	fmt.Scan(&input)
	startDate, err = time.Parse(formatDate, input)
	if err != nil {
		fmt.Println("Tanggal join karyawan hanya menerima format yyyy-mm-dd")
		return
	}

	fmt.Print("* Tanggal rencana cuti = ")
	fmt.Scan(&input)
	cutiDate, err = time.Parse(formatDate, input)
	if err != nil {
		fmt.Println("Tanggal rencana cuti hanya menerima format yyyy-mm-dd")
		return
	}
	fmt.Print("* Durasi cuti (hari) = ")
	fmt.Scan(&input)
	durationCuti, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("Jumlah Durasi cuti (hari) Hanya menerima bilangan bulat")
		return
	}

	fmt.Println("\nOutput: ")
	if cutiDate.Before(startDate) {
		fmt.Println(" False")
		fmt.Println(" Alasan: Keputusan rencana cuti harus setelah tanggal join karyawan")
		return
	}

	cutiStart := startDate.AddDate(0, 0, 180)
	if cutiStart.After(cutiDate) {
		fmt.Println(" False")
		fmt.Println(" Alasan: Karena belum 180 hari sejak tanggal join karyawan")
		return
	}

	cutiEnd := time.Date(cutiStart.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	totalKerjaBisaCuti := 0
	for d := cutiStart; d.After(cutiEnd) == false; d = d.AddDate(0, 0, 1) {
		totalKerjaBisaCuti++
	}

	maxCutiOnYear := roundDownFloatToInt(float64(totalKerjaBisaCuti) / float64(365) * float64(cutiBersama))
	if durationCuti > maxCutiOnYear {
		fmt.Println(" False")
		fmt.Printf(" Alasan: Hanya boleh mengambil %v hari cuti", maxCutiOnYear)
		return
	}

	if durationCuti > maxCutiBeruntun {
		fmt.Println(" False")
		fmt.Printf(" Alasan: Hanya boleh mengambil %v hari cuti berturutan.", maxCutiBeruntun)
		return
	}

	fmt.Println(" True")

}
