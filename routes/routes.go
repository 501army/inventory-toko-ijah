package routes

import (
	"github.com/gin-gonic/gin"
)

//DefaultRoute is
type DefaultRoute struct{}

//Init is
func (d *DefaultRoute) Init(router *gin.Engine) {
	// searchController := new(controllers.SearchController)
	// stationController := new(controllers.StationController)
	// bookController := new(controllers.BookController)
	// seatController := new(controllers.SeatController)
	// v1 := router.Group("/v1")
	// {
	// 	v1.GET("/train/station/add", stationController.AddStation)
	// 	v1.GET("/train/station", stationController.Station)
	// 	v1.GET("/train/search", searchController.Search)
	// 	v1.POST("/train/book", bookController.Booking)
	// 	v1.POST("/train/book/cancel", bookController.CancelBooking)
	// 	v1.GET("/train/seat", seatController.SeatMap)
	// 	v1.POST("/train/seat/change", seatController.ChangeSeat)
	// 	v1.GET("/train/retrieve", bookController.Retrieve)
	// }
	// v2 := router.Group("v2")
	// {
	// 	v2.GET("/train/search", searchController.Searchv2)
	// 	v2.POST("/train/book", bookController.Bookingv2)
	// 	v2.GET("/train/retrieve", bookController.Retrievev2)
	// 	v2.POST("/train/book/cancel", bookController.CancelBookingv2)
	// }
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"status":  404,
			"message": "not found",
		})
	})
}
