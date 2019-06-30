package controllers

import (
	"bytes"
	"encoding/csv"
	"inventory-toko-ijah/utils/forms"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// BarangKeluarController is
type BarangKeluarController struct{}

// AddBarangKeluar is
func (b *BarangKeluarController) AddBarangKeluar(c *gin.Context) {
	var barangKeluar forms.BarangKeluar
	c.Request.ParseForm()
	barangKeluar.BarangID, _ = strconv.Atoi(c.PostForm("id_barang"))
	barangKeluar.Waktu = c.PostForm("waktu")
	barangKeluar.KodePesanan = c.PostForm("kode_pesanan")
	barangKeluar.JumlahKeluar, _ = strconv.Atoi(c.PostForm("jml_keluar"))
	barangKeluar.HargaJual, _ = strconv.Atoi(c.PostForm("harga_jual"))
	barangKeluar.TotalHarga = barangKeluar.JumlahKeluar * barangKeluar.HargaJual
	barangKeluar.Catatan = c.PostForm("catatan")
	strErr := inventoryModel.AddBarangKeluar(barangKeluar)
	if strErr != "" {
		c.JSON(404, gin.H{
			"status":  404,
			"message": strErr,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "success",
		})
	}
}

// LaporanBarangKeluar is
func (b *BarangKeluarController) LaporanBarangKeluar(c *gin.Context) {
	var laporanKeluar forms.LaporanBarangKeluar
	barangKeluar := inventoryModel.GetAllBarangKeluar()
	var prevKodePesanan string
	for _, keluar := range barangKeluar {
		var pesanan forms.Pesanan
		if prevKodePesanan != keluar.KodePesanan {
			pesanan.KodePesanan = keluar.KodePesanan
			pesanan.Waktu = keluar.Waktu
			laporanKeluar.Pesanan = append(laporanKeluar.Pesanan, pesanan)
		}
		prevKodePesanan = keluar.KodePesanan
	}
	omzet := 0
	totalPenjualan := 0
	totalBarang := 0
	laba := 0
	var newPesanan []forms.Pesanan
	for _, pesanan := range laporanKeluar.Pesanan {
		totalPenjualan = totalPenjualan + 1
		itemBarangKeluar := inventoryModel.GetBarangKeluarbyPesanan(pesanan.KodePesanan)
		for _, item := range itemBarangKeluar {
			var itemPesanan forms.ItemPesanan
			barang := inventoryModel.GetBarangbyID(item.BarangID)
			itemPesanan.SKU = barang.Sku
			itemPesanan.Nama = barang.Nama
			itemPesanan.Jumlah = item.JumlahKeluar
			totalBarang = totalBarang + itemPesanan.Jumlah
			itemPesanan.HargaJual = item.HargaJual
			itemPesanan.Total = itemPesanan.Jumlah * itemPesanan.HargaJual
			omzet = omzet + itemPesanan.Total
			brgMasuk := inventoryModel.GetBarangMasukbyBarang(barang.BarangID)
			tempJml := 0
			avgBeli := 0
			if len(brgMasuk) > 0 {
				for _, item := range brgMasuk {
					tempJml = tempJml + item.HargaBeli
				}
				avgBeli = tempJml / len(brgMasuk)
			}
			itemPesanan.HargaBeli = avgBeli
			itemPesanan.Laba = itemPesanan.Total - (itemPesanan.Jumlah * itemPesanan.HargaBeli)
			laba = laba + itemPesanan.Laba
			pesanan.Barang = append(pesanan.Barang, itemPesanan)
		}
		newPesanan = append(newPesanan, pesanan)
	}
	laporanKeluar.TglCetak = time.Now().Format("02 January 2006")
	laporanKeluar.Omzet = omzet
	laporanKeluar.LabaKotor = laba
	laporanKeluar.TotalPenjualan = totalPenjualan
	laporanKeluar.TotalBarang = totalBarang
	laporanKeluar.Pesanan = newPesanan
	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data":    laporanKeluar,
	})
}

// GenerateCsv is
func (b *BarangKeluarController) GenerateCsv(c *gin.Context) {
	buf := &bytes.Buffer{}
	w := csv.NewWriter(buf)

	var headerCsv []string
	var brg forms.BarangKeluarComplete
	val := reflect.ValueOf(brg)
	for i := 0; i < val.Type().NumField(); i++ {
		headerCsv = append(headerCsv, val.Type().Field(i).Tag.Get("json"))
	}
	if err := w.Write(headerCsv); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}
	listBarang := inventoryModel.GetBarangKeluarComplete()
	for _, item := range listBarang {
		var rowCsv []string
		rowCsv = append(rowCsv, item.KodePesanan)
		rowCsv = append(rowCsv, item.Waktu)
		rowCsv = append(rowCsv, item.SKU)
		rowCsv = append(rowCsv, item.Nama)
		rowCsv = append(rowCsv, strconv.Itoa(item.JumlahKeluar))
		rowCsv = append(rowCsv, strconv.Itoa(item.HargaJual))
		rowCsv = append(rowCsv, strconv.Itoa(item.TotalHarga))
		rowCsv = append(rowCsv, item.Catatan)
		if err := w.Write(rowCsv); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=barang_masuk.csv")
	c.Data(http.StatusOK, "text/csv", buf.Bytes())
}
