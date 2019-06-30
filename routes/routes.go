package routes

import (
	"inventory-toko-ijah/controllers"

	"github.com/gin-gonic/gin"
)

//DefaultRoute is
type DefaultRoute struct{}

//Init is
func (d *DefaultRoute) Init(router *gin.Engine) {
	barangController := new(controllers.BarangController)
	barangMasukController := new(controllers.BarangMasukController)
	barangKeluarController := new(controllers.BarangKeluarController)
	v1 := router.Group("/v1")
	{
		v1.GET("/barang", barangController.GetAllBarang)
		v1.POST("/barang/add", barangController.AddBarang)
		v1.GET("/barang/:id_barang", barangController.GetBarangbyID)
		v1.POST("/barang/update/:id_barang", barangController.UpdateBarang)
		v1.POST("/barangmasuk/add", barangMasukController.AddBarangMasuk)
		v1.POST("/barangkeluar/add", barangKeluarController.AddBarangKeluar)
		v1.GET("/laporan/barang", barangController.LaporanBarang)
		v1.GET("/laporan/penjualan", barangKeluarController.LaporanBarangKeluar)
		v1.GET("/csv/barang", barangController.GenerateCsv)
		v1.GET("/csv/barangmasuk", barangMasukController.GenerateCsv)
		v1.GET("/csv/barangkeluar", barangKeluarController.GenerateCsv)
	}
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"status":  404,
			"message": "not found",
		})
	})
}
