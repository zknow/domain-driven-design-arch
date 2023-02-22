package delivery

import (
	"arch/domain"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ExampleHandler ...
type ExampleHandler struct {
	exampleUsecase domain.ExampleUsecase
}

// SetExampleHandler ...
func SetExampleHandler(e *gin.Engine, exampleUsecase domain.ExampleUsecase) {
	handler := &ExampleHandler{
		exampleUsecase: exampleUsecase,
	}
	g := e.Group("v1")
	g.GET("/example/:id", handler.GetDataByID)
	g.POST("/example", handler.PostToCreate)
}

// GetDataByID ...
func (d *ExampleHandler) GetDataByID(c *gin.Context) {
	id := c.Param("id")

	data, err := d.exampleUsecase.GetByID(c, id)
	if err != nil {
		logrus.Error(err)
		c.JSON(500, &gin.H{
			"Code":    3000,
			"Message": "Internal error. Query data error",
		})
		return
	}

	c.JSON(200, data)
}

// PostToCreate ...
func (d *ExampleHandler) PostToCreate(c *gin.Context) {
	var body domain.Example
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.JSON(500, &gin.H{
			"Code":    3000,
			"Message": "Internal error. Parsing failed",
		})
		return
	}
	data := domain.Example{
		Name: body.Name,
	}
	if err := d.exampleUsecase.Store(c, &data); err != nil {
		logrus.Error(err)
		c.JSON(500, &gin.H{
			"Code":    3000,
			"Message": "Internal error. Store failed",
		})
		return
	}

	c.JSON(200, data)
}
