package models

import (
	"inventory-toko-ijah/utils/db"
	"inventory-toko-ijah/utils/forms"
)

//InventoryModel is
type InventoryModel struct{}

// GetAllBarang is
func (i *InventoryModel) GetAllBarang() []forms.Barang {
	var barang []forms.Barang
	db.GetDB().Table("barang").Find(&barang)
	return barang
}

// AddBarang is
func (i *InventoryModel) AddBarang(barang forms.Barang) {
	db.GetDB().Table("barang").Create(&barang)
}

// GetBarangbyID is
func (i *InventoryModel) GetBarangbyID(barangID int) forms.Barang {
	var barang forms.Barang
	db.GetDB().
		Table("barang").
		Where("id_barang = ?", barangID).
		First(&barang)
	return barang
}

// UpdateBarang is
func (i *InventoryModel) UpdateBarang(barangID int, data forms.Barang) {
	var barang forms.Barang
	db.GetDB().Table("barang").Where("id_barang = ?", barangID).First(&barang)
	db.GetDB().Table("barang").Model(&barang).Updates(forms.Barang{Sku: data.Sku, Nama: data.Nama, Jumlah: data.Jumlah})
}

// AddBarangMasuk is
func (i *InventoryModel) AddBarangMasuk(data forms.BarangMasuk) string {
	var barang forms.Barang
	var strErr string
	db.GetDB().
		Table("barang").
		Where("id_barang = ?", data.BarangID).
		First(&barang)
	if barang.BarangID == 0 {
		strErr = "barang tidak ditemukan"
	} else {
		db.GetDB().Table("barang_masuk").Create(&data)
		newJml := barang.Jumlah + data.JumlahTerima
		db.GetDB().Table("barang").Where("id_barang = ?", data.BarangID).First(&barang)
		db.GetDB().Table("barang").Model(&barang).Updates(forms.Barang{Jumlah: newJml})
	}
	return strErr
}

// AddBarangKeluar is
func (i *InventoryModel) AddBarangKeluar(data forms.BarangKeluar) string {
	var barang forms.Barang
	var strErr string
	db.GetDB().
		Table("barang").
		Where("id_barang = ?", data.BarangID).
		First(&barang)
	if barang.BarangID == 0 {
		strErr = "barang tidak ditemukan"
	} else {
		db.GetDB().Table("barang_keluar").Create(&data)
		newJml := barang.Jumlah - data.JumlahKeluar
		db.GetDB().Table("barang").Where("id_barang = ?", data.BarangID).First(&barang)
		db.GetDB().Table("barang").Model(&barang).Updates(forms.Barang{Jumlah: newJml})
	}
	return strErr
}

// GetBarangMasukbyBarang is
func (i *InventoryModel) GetBarangMasukbyBarang(barangID int) []forms.BarangMasuk {
	var barang []forms.BarangMasuk
	db.GetDB().
		Table("barang_masuk").
		Where("id_barang = ?", barangID).
		Find(&barang)
	return barang
}

// GetBarangMasukComplete is
func (i *InventoryModel) GetBarangMasukComplete() []forms.BarangMasukComplete {
	var barang []forms.BarangMasukComplete
	db.GetDB().Table("barang_masuk").
		Select("barang.sku_barang,barang.nama_barang,barang_masuk.waktu,jml_pesan,jml_terima,harga_beli,total_harga,no_kwitansi,catatan").
		Joins("JOIN barang ON barang.id_barang=barang_masuk.id_barang").
		Scan(&barang)
	return barang
}

// GetAllBarangKeluar is
func (i *InventoryModel) GetAllBarangKeluar() []forms.BarangKeluar {
	var barang []forms.BarangKeluar
	db.GetDB().Table("barang_keluar").Find(&barang)
	return barang
}

// GetBarangKeluarComplete is
func (i *InventoryModel) GetBarangKeluarComplete() []forms.BarangKeluarComplete {
	var barang []forms.BarangKeluarComplete
	db.GetDB().Table("barang_keluar").
		Select("kode_pesanan,waktu,barang.sku_barang,barang.nama_barang,jml_keluar,harga_jual,total_harga,catatan").
		Joins("JOIN barang ON barang.id_barang=barang_keluar.id_barang").
		Scan(&barang)
	return barang
}

// GetBarangKeluarbyPesanan is
func (i *InventoryModel) GetBarangKeluarbyPesanan(kodePesanan string) []forms.BarangKeluar {
	var barang []forms.BarangKeluar
	db.GetDB().
		Table("barang_keluar").
		Where("kode_pesanan = ?", kodePesanan).
		First(&barang)
	return barang
}
