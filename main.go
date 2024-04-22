package main

import (
	"errors"
	"fmt"
)

func PembayaranBarang(hargaTotal float64, metodePembayaran string, dicicil bool) error {

	if hargaTotal <= 0 {
		return errors.New("harga tidak bisa nol")
	}

	switch metodePembayaran {
	case "cod", "transfer", "debit", "credit", "gerai":

		if metodePembayaran == "credit" && !dicicil {
			return errors.New("credit harus dicicil")
		}

		if dicicil && (metodePembayaran != "credit" || hargaTotal <= 500000) {
			return errors.New("cicilan tidak memenuhi syarat")
		}
	default:
		return errors.New("metode tidak dikenali")
	}

	return nil
}

func main() {

	harga := 600000.0
	metode := "credit"
	dicicil := true

	err := PembayaranBarang(harga, metode, dicicil)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Pembayaran berhasil")
	}
}
