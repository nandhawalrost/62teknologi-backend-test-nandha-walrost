package controllers

import (
	"enamdua/initializers"
	"enamdua/models"

	"github.com/gin-gonic/gin"
)

func BusinessesCreate(c *gin.Context) {
	var body struct {
		Location                 string
		Latitude                 uint
		Longitude                uint
		Term                     string
		Radius                   uint
		Categories               string
		Locale                   string
		Price                    uint
		Open_now                 bool
		Open_at                  uint
		Attributes               string
		Sort_by                  string
		Device_platform          string
		Reservation_date         string
		Reservation_time         string
		Reservation_covers       uint
		Matches_party_size_param bool
		Limit                    uint
		Offset                   uint
	}

	c.Bind(&body)

	business := models.Business{
		Location:                 body.Location,
		Latitude:                 body.Latitude,
		Longitude:                body.Longitude,
		Term:                     body.Term,
		Radius:                   body.Radius,
		Categories:               body.Categories,
		Locale:                   body.Locale,
		Price:                    body.Price,
		Open_now:                 body.Open_now,
		Open_at:                  body.Open_at,
		Attributes:               body.Attributes,
		Sort_by:                  body.Sort_by,
		Device_platform:          body.Device_platform,
		Reservation_date:         body.Reservation_date,
		Reservation_time:         body.Reservation_time,
		Reservation_covers:       body.Reservation_covers,
		Matches_party_size_param: body.Matches_party_size_param,
		Limit:                    body.Limit,
		Offset:                   body.Offset,
	}

	result := initializers.DB.Create(&business) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"message": "created!",
	})
}

func BusinessesUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Location                 string
		Latitude                 uint
		Longitude                uint
		Term                     string
		Radius                   uint
		Categories               string
		Locale                   string
		Price                    uint
		Open_now                 bool
		Open_at                  uint
		Attributes               string
		Sort_by                  string
		Device_platform          string
		Reservation_date         string
		Reservation_time         string
		Reservation_covers       uint
		Matches_party_size_param bool
		Limit                    uint
		Offset                   uint
	}

	c.Bind(&body)

	var business models.Business
	initializers.DB.First(&business, id)

	initializers.DB.Model(&business).Updates(models.Business{
		Location:                 body.Location,
		Latitude:                 body.Latitude,
		Longitude:                body.Longitude,
		Term:                     body.Term,
		Radius:                   body.Radius,
		Categories:               body.Categories,
		Locale:                   body.Locale,
		Price:                    body.Price,
		Open_now:                 body.Open_now,
		Open_at:                  body.Open_at,
		Attributes:               body.Attributes,
		Sort_by:                  body.Sort_by,
		Device_platform:          body.Device_platform,
		Reservation_date:         body.Reservation_date,
		Reservation_time:         body.Reservation_time,
		Reservation_covers:       body.Reservation_covers,
		Matches_party_size_param: body.Matches_party_size_param,
		Limit:                    body.Limit,
		Offset:                   body.Offset,
	})

	if body.Open_now == false {
		initializers.DB.Exec("UPDATE businesses SET open_now = ? WHERE id = ?", 0, id)
	} else {
		initializers.DB.Exec("UPDATE businesses SET open_now = ? WHERE id = ?", 1, id)
	}

	if body.Matches_party_size_param == false {
		initializers.DB.Exec("UPDATE businesses SET matches_party_size_param = ? WHERE id = ?", 0, id)
	} else {
		initializers.DB.Exec("UPDATE businesses SET matches_party_size_param = ? WHERE id = ?", 1, id)
	}

	c.JSON(200, gin.H{
		"message": "updated!",
	})
}

func BusinessesDelete(c *gin.Context) {
	id := c.Param("id")

	var businesses []models.Business
	initializers.DB.Delete(&businesses, id)

	c.JSON(200, gin.H{
		"message": "deleted!",
	})
}

func BusinessesGet(c *gin.Context) {
	term := c.Param("term")
	location := c.Param("location")
	latitude := c.Param("latitude")
	longitude := c.Param("longitude")
	radius := c.Param("radius")

	var Businesses []models.Business
	// initializers.DB.First(&Businesses, id)
	// initializers.DB.Where("term LIKE ?", "%"+term+"%").Or("location LIKE ?", "%"+location+"%").Find(&Businesses)
	initializers.DB.Where("term LIKE ? AND location LIKE ? AND latitude LIKE ? AND longitude LIKE ? AND radius LIKE ? ", "%"+term+"%", "%"+location+"%", "%"+latitude+"%", "%"+longitude+"%", "%"+radius+"%").Find(&Businesses)
	c.JSON(200, gin.H{
		"businesses": Businesses,
	})
}
