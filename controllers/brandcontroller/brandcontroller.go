package brandcontroller

import (
	"go-web-native/entities"
	"go-web-native/models/brandmodel"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	brands := brandmodel.GetAll()
	c.HTML(http.StatusOK, "home-brand.html", gin.H{
		"brands": brands,
	})
}

func AddForm(c *gin.Context) {
	c.HTML(http.StatusOK, "create-brand.html", nil)
}

func Add(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.HTML(http.StatusBadRequest, "create-brand.html", gin.H{
			"error": "Brand name is required",
		})
		return
	}

	brand := entities.Brand{
		Name:      name,
		CreatedAt: time.Now(),
	}

	if err := brandmodel.AddBrand(brand); err != nil {
		c.HTML(http.StatusInternalServerError, "create-brand.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/brands")
}

func EditForm(c *gin.Context) {
	id := c.Param("id")
	brand := brandmodel.Detail(id)
	c.HTML(http.StatusOK, "edit-brand.html", gin.H{
		"brand": brand,
	})
}

func Edit(c *gin.Context) {
	id := c.Param("id")
	name := c.PostForm("name")

	if name == "" {
		c.HTML(http.StatusBadRequest, "edit-brand.html", gin.H{
			"error": "Brand name is required",
		})
		return
	}

	brand := entities.Brand{
		Name:      name,
		UpdatedAt: time.Now(),
	}

	if ok := brandmodel.Update(id, brand); !ok {
		c.HTML(http.StatusInternalServerError, "edit-brand.html", gin.H{
			"error": "Failed to update brand",
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/brands")
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	if err := brandmodel.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, "/brands")
}
