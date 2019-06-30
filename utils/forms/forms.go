package forms

//Barang is
type Barang struct {
	BarangID int    `gorm:"primary_key;column:id_barang" json:"id"`
	Sku      string `gorm:"column:sku_barang" json:"sku"`
	Nama     string `gorm:"column:nama_barang" json:"nama"`
	Jumlah   int    `gorm:"column:jumlah_barang" json:"jumlah"`
}

//BarangMasuk is
type BarangMasuk struct {
	BarangMasukID int    `gorm:"primary_key;column:id_barang_masuk"`
	BarangID      int    `gorm:"column:id_barang" json:"id_barang"`
	Waktu         string `gorm:"column:waktu" json:"waktu"`
	JumlahPesan   int    `gorm:"column:jml_pesan" json:"jml_pesan"`
	JumlahTerima  int    `gorm:"column:jml_terima" json:"jml_terima"`
	HargaBeli     int    `gorm:"column:harga_beli" json:"harga_beli"`
	TotalHarga    int    `gorm:"column:total_harga"`
	Kwitansi      string `gorm:"column:no_kwitansi" json:"kwitansi"`
	Catatan       string `gorm:"column:catatan" json:"catatan"`
}

//BarangMasukComplete is
type BarangMasukComplete struct {
	SKU          string `gorm:"column:sku_barang" json:"sku_barang"`
	Nama         string `gorm:"column:nama_barang" json:"nama_barang"`
	Waktu        string `gorm:"column:waktu" json:"waktu"`
	JumlahPesan  int    `gorm:"column:jml_pesan" json:"jml_pesan"`
	JumlahTerima int    `gorm:"column:jml_terima" json:"jml_terima"`
	HargaBeli    int    `gorm:"column:harga_beli" json:"harga_beli"`
	TotalHarga   int    `gorm:"column:total_harga" json:"total_harga"`
	Kwitansi     string `gorm:"column:no_kwitansi" json:"kwitansi"`
	Catatan      string `gorm:"column:catatan" json:"catatan"`
}

//BarangKeluar is
type BarangKeluar struct {
	BarangKeluarID int    `gorm:"primary_key;column:id_barang_keluar"`
	BarangID       int    `gorm:"column:id_barang" json:"id_barang"`
	KodePesanan    string `gorm:"column:kode_pesanan" json:"kode_pesanan"`
	Waktu          string `gorm:"column:waktu" json:"waktu"`
	JumlahKeluar   int    `gorm:"column:jml_keluar" json:"jml_keluar"`
	HargaJual      int    `gorm:"column:harga_jual" json:"harga_jual"`
	TotalHarga     int    `gorm:"column:total_harga"`
	Catatan        string `gorm:"column:catatan" json:"catatan"`
}

//BarangKeluarComplete is
type BarangKeluarComplete struct {
	KodePesanan  string `gorm:"column:kode_pesanan" json:"kode_pesanan"`
	Waktu        string `gorm:"column:waktu" json:"waktu"`
	SKU          string `gorm:"column:sku_barang" json:"sku_barang"`
	Nama         string `gorm:"column:nama_barang" json:"nama_barang"`
	JumlahKeluar int    `gorm:"column:jml_keluar" json:"jml_keluar"`
	HargaJual    int    `gorm:"column:harga_jual" json:"harga_jual"`
	TotalHarga   int    `gorm:"column:total_harga" json:"total_harga"`
	Catatan      string `gorm:"column:catatan" json:"catatan"`
}

// ItemLaporanBarang is
type ItemLaporanBarang struct {
	SKU       string `json:"sku"`
	Nama      string `json:"nama"`
	Jumlah    int    `json:"jumlah"`
	Rata2Beli int    `json:"avg_harga_beli"`
	Total     int    `json:"total"`
}

// LaporanBarang is
type LaporanBarang struct {
	TglCetak   string              `json:"tgl_cetak"`
	JmlSku     int                 `json:"jml_sku"`
	JmlTotal   int                 `json:"jml_barang"`
	TotalNilai int                 `json:"total_nilai"`
	Item       []ItemLaporanBarang `json:"barang"`
}

// ItemPesanan is
type ItemPesanan struct {
	SKU       string `json:"sku"`
	Nama      string `json:"nama"`
	Jumlah    int    `json:"jumlah"`
	HargaJual int    `json:"harga_jual"`
	Total     int    `json:"total"`
	HargaBeli int    `json:"harga_beli"`
	Laba      int    `json:"laba"`
}

// Pesanan is
type Pesanan struct {
	KodePesanan string        `json:"id_pesanan"`
	Waktu       string        `json:"pesanan"`
	Barang      []ItemPesanan `json:"item"`
}

// LaporanBarangKeluar is
type LaporanBarangKeluar struct {
	TglCetak       string    `json:"tgl_cetak"`
	Omzet          int       `json:"omzet"`
	LabaKotor      int       `json:"laba_kotor"`
	TotalPenjualan int       `json:"total_penjualan"`
	TotalBarang    int       `json:"total_barang"`
	Pesanan        []Pesanan `json:"pesanan"`
}
