package controllers

import (
	"bytes"
	"encoding/csv"
	"inventory-toko-ijah/models"
	"inventory-toko-ijah/utils/forms"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// BarangController is
type BarangController struct{}

var inventoryModel = new(models.InventoryModel)

// GetAllBarang is
func (b *BarangController) GetAllBarang(c *gin.Context) {
	listBarang := inventoryModel.GetAllBarang()
	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data": gin.H{
			"barang": listBarang,
		},
	})
}

// AddBarang is
func (b *BarangController) AddBarang(c *gin.Context) {
	var barang forms.Barang
	c.Request.ParseForm()
	barang.Sku = c.PostForm("sku")
	barang.Nama = c.PostForm("nama")
	jml, _ := strconv.Atoi(c.PostForm("jumlah"))
	barang.Jumlah = jml
	inventoryModel.AddBarang(barang)
	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
	})
}

// GetBarangbyID is
func (b *BarangController) GetBarangbyID(c *gin.Context) {
	barangID, _ := strconv.Atoi(c.Params.ByName("id_barang"))
	barang := inventoryModel.GetBarangbyID(barangID)
	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data":    barang,
	})
}

// UpdateBarang is
func (b *BarangController) UpdateBarang(c *gin.Context) {
	barangID, _ := strconv.Atoi(c.Params.ByName("id_barang"))
	var barang forms.Barang
	c.Request.ParseForm()
	barang.Sku = c.PostForm("sku")
	barang.Nama = c.PostForm("nama")
	jml, _ := strconv.Atoi(c.PostForm("jumlah"))
	barang.Jumlah = jml
	inventoryModel.UpdateBarang(barangID, barang)
	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
	})
}

// LaporanBarang is
func (b *BarangController) LaporanBarang(c *gin.Context) {
	var laporan forms.LaporanBarang
	listBarang := inventoryModel.GetAllBarang()
	jmlSku := len(listBarang)
	totalJml := 0
	totalNilai := 0
	for _, barang := range listBarang {
		var itemBarang forms.ItemLaporanBarang
		itemBarang.SKU = barang.Sku
		itemBarang.Nama = barang.Nama
		itemBarang.Jumlah = barang.Jumlah
		tempJml := 0
		brgMasuk := inventoryModel.GetBarangMasukbyBarang(barang.BarangID)
		avgBeli := 0
		if len(brgMasuk) > 0 {
			for _, item := range brgMasuk {
				tempJml = tempJml + item.HargaBeli
			}
			avgBeli = tempJml / len(brgMasuk)
		}
		itemBarang.Rata2Beli = avgBeli
		itemBarang.Total = itemBarang.Jumlah * itemBarang.Rata2Beli
		laporan.Item = append(laporan.Item, itemBarang)
		totalJml = totalJml + barang.Jumlah
		totalNilai = totalNilai + itemBarang.Total
	}
	laporan.TglCetak = time.Now().Format("02 January 2006")
	laporan.JmlSku = jmlSku
	laporan.JmlTotal = totalJml
	laporan.TotalNilai = totalNilai
	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data":    laporan,
	})
}

// GenerateCsv is
func (b *BarangController) GenerateCsv(c *gin.Context) {
	buf := &bytes.Buffer{}
	w := csv.NewWriter(buf)

	var headerCsv []string
	var brg forms.Barang
	val := reflect.ValueOf(brg)
	for i := 0; i < val.Type().NumField(); i++ {
		headerCsv = append(headerCsv, val.Type().Field(i).Tag.Get("json"))
	}
	if err := w.Write(headerCsv); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}
	listBarang := inventoryModel.GetAllBarang()
	for _, item := range listBarang {
		var rowCsv []string
		rowCsv = append(rowCsv, strconv.Itoa(item.BarangID))
		rowCsv = append(rowCsv, item.Sku)
		rowCsv = append(rowCsv, item.Nama)
		rowCsv = append(rowCsv, strconv.Itoa(item.Jumlah))
		if err := w.Write(rowCsv); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=barang.csv")
	c.Data(http.StatusOK, "text/csv", buf.Bytes())
}
