package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	formatDate = "2006-01-02"
)

func main() {
	var (
		input        string
		startDate    time.Time
		cutiDate     time.Time
		cutiBersama  int64
		cutiPribadi  int64
		durationCuti int64
		err          error
	)

	fmt.Println("Input: ")
	fmt.Print("* Jumlah Cuti Bersama = ")
	fmt.Scan(&input)
	cutiBersama, err = strconv.ParseInt(input, 10, 64)
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
	durationCuti, err = strconv.ParseInt(input, 10, 64)
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
	if cutiStart.Before(cutiDate) {
		fmt.Println(" False")
		fmt.Println(" Alasan: Karena belum 180 hari sejak tanggal join karyawan")
		return
	}

	fmt.Println("Sedang tahap pengerjaan")
	fmt.Println("cutibersama", cutiBersama)
	fmt.Println("cutipribadi", cutiPribadi)
	fmt.Println("durationCuti", durationCuti)

}
