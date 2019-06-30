package controllers

import (
	"bytes"
	"encoding/csv"
	"inventory-toko-ijah/utils/forms"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BarangMasukController is
type BarangMasukController struct{}

// AddBarangMasuk is
func (b *BarangMasukController) AddBarangMasuk(c *gin.Context) {
	var barangMasuk forms.BarangMasuk
	c.Request.ParseForm()
	barangMasuk.BarangID, _ = strconv.Atoi(c.PostForm("id_barang"))
	barangMasuk.Waktu = c.PostForm("waktu")
	barangMasuk.JumlahPesan, _ = strconv.Atoi(c.PostForm("jml_pesan"))
	barangMasuk.JumlahTerima, _ = strconv.Atoi(c.PostForm("jml_terima"))
	barangMasuk.HargaBeli, _ = strconv.Atoi(c.PostForm("harga_beli"))
	barangMasuk.TotalHarga = barangMasuk.JumlahPesan * barangMasuk.HargaBeli
	barangMasuk.Kwitansi = c.PostForm("kwitansi")
	barangMasuk.Catatan = c.PostForm("catatan")
	strErr := inventoryModel.AddBarangMasuk(barangMasuk)
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

// GenerateCsv is
func (b *BarangMasukController) GenerateCsv(c *gin.Context) {
	buf := &bytes.Buffer{}
	w := csv.NewWriter(buf)

	var headerCsv []string
	var brg forms.BarangMasukComplete
	val := reflect.ValueOf(brg)
	for i := 0; i < val.Type().NumField(); i++ {
		headerCsv = append(headerCsv, val.Type().Field(i).Tag.Get("json"))
	}
	if err := w.Write(headerCsv); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}
	listBarang := inventoryModel.GetBarangMasukComplete()
	for _, item := range listBarang {
		var rowCsv []string
		rowCsv = append(rowCsv, item.SKU)
		rowCsv = append(rowCsv, item.Nama)
		rowCsv = append(rowCsv, item.Waktu)
		rowCsv = append(rowCsv, strconv.Itoa(item.JumlahPesan))
		rowCsv = append(rowCsv, strconv.Itoa(item.JumlahTerima))
		rowCsv = append(rowCsv, strconv.Itoa(item.HargaBeli))
		rowCsv = append(rowCsv, strconv.Itoa(item.TotalHarga))
		rowCsv = append(rowCsv, item.Kwitansi)
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
