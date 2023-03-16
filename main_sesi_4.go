package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	No        string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	dataBiodata := []Biodata{
		{
			No:        "1",
			Nama:      "Arman",
			Alamat:    "Jln Podomoro",
			Pekerjaan: "Pengangguran",
			Alasan:    "Ingin Menambah Ilmu",
		},
		{
			No:        "2",
			Nama:      "Afif",
			Alamat:    "Jln Kembayan",
			Pekerjaan: "Pengangguran",
			Alasan:    "Ingin Menambah Ilmu",
		},
	}
	osArgsInput(dataBiodata)
}

func osArgsInput(dataBiodata []Biodata) {
	if len(os.Args) == 2 {
		commandAngka := os.Args[1]
		for _, loop := range dataBiodata {
			if loop.No == commandAngka {
				fmt.Println("Nama :", loop.Nama)
				fmt.Println("Alamat :", loop.Alamat)
				fmt.Println("Pekerjaan :", loop.Pekerjaan)
				fmt.Println("Alasan :", loop.Alasan)
			}
		}
	} else {
		fmt.Println("Tidak ada argumen secara spesifik, hanya bisa satu argumen")
	}
}
