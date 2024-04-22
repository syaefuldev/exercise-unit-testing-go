package main

import "testing"

func TestPembayaranBarang(t *testing.T) {
	tests := []struct {
		name             string
		hargaTotal       float64
		metodePembayaran string
		dicicil          bool
		wantErr          bool
	}{
		{
			name:             "Harga total nol",
			hargaTotal:       0,
			metodePembayaran: "cod",
			dicicil:          false,
			wantErr:          true,
		},
		{
			name:             "Metode pembayaran tidak dikenali",
			hargaTotal:       100000,
			metodePembayaran: "paypal",
			dicicil:          false,
			wantErr:          true,
		},
		{
			name:             "Pembayaran credit tanpa dicicil",
			hargaTotal:       600000,
			metodePembayaran: "credit",
			dicicil:          false,
			wantErr:          true,
		},
		{
			name:             "Pembayaran cicil dengan metode selain credit",
			hargaTotal:       600000,
			metodePembayaran: "transfer",
			dicicil:          true,
			wantErr:          true,
		},
		{
			name:             "Pembayaran cicil dengan harga total kurang dari 500000",
			hargaTotal:       400000,
			metodePembayaran: "credit",
			dicicil:          true,
			wantErr:          true,
		},
		{
			name:             "Pembayaran berhasil",
			hargaTotal:       700000,
			metodePembayaran: "credit",
			dicicil:          true,
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PembayaranBarang(tt.hargaTotal, tt.metodePembayaran, tt.dicicil); (err != nil) != tt.wantErr {
				t.Errorf("PembayaranBarang() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
