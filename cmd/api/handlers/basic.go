package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/reinaldocomputer/basicgoproject/internal/basic"
	"github.com/reinaldocomputer/basicgoproject/internal/platform/mockDB"
	"net/http"
	"strconv"
)

func InsertBasic(c *gin.Context) {
	basicData := []basic.Request{}
	if err := c.ShouldBindJSON(&basicData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, data := range basicData {
		newBasicData := basic.NewBasic(data)
		if err := newBasicData.Insert(); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, `
		{ 
			"status": "OK",
			"message" "Data inserted with success.",
		}
	`)
}

func GetBasicByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var basicData basic.Basic
	basicData.Id = idInt
	gotData, err := basicData.GetByID()
	switch err {
	case mockDB.DataNotFoundByIDError:
		c.JSON(http.StatusBadRequest, err.Error())
		return
	case nil:
	default:
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gotData)
}

func DeleteBasicByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var basicData basic.Basic
	basicData.Id = idInt
	err = basicData.DeleteByID()
	switch err {
	case mockDB.DataNotFoundByIDError:
		c.JSON(http.StatusBadRequest, err.Error())
		return
	case nil:
	default:
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, `
		{ 
			"status": "OK",
			"message" "Data was deleted successfully",
		}
	`)
}

func UpdateBasicByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var basicData basic.Basic
	if err = c.ShouldBindJSON(&basicData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	basicData.Id = idInt
	got, err := basicData.UpdateByID()
	switch err {
	case mockDB.DataNotFoundByIDError:
		c.JSON(http.StatusBadRequest, err.Error())
		return
	case nil:
	default:
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, got)
}

func GetBasicAll(c *gin.Context) {
	data, err := basic.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}
