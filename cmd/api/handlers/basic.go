package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/reinaldocomputer/basicgoproject/internal/basic"
	"github.com/reinaldocomputer/basicgoproject/internal/platform/mockDB"
	"net/http"
	"strconv"
	"time"
)

// set the number of works to use, generally, quantity of cores.
const numWorkers = 6

func InsertBasic(c *gin.Context) {
	startTime := time.Now()
	// check request schema
	basicData := []basic.Request{}
	if err := c.ShouldBindJSON(&basicData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// create a channel to receive data to insert
	dataChan := make(chan *basic.Basic, len(basicData))

	//start the workers
	for i := 0; i < numWorkers; i++ {
		go func(worker int) {
			for data := range dataChan {
				//fmt.Println("W:", worker, " Inserting Data: ", data.Id)
				// insert Data in DB using a GoRoutine
				if err := data.Insert(); err != nil {
					c.JSON(http.StatusInternalServerError, `
						{
							"status": "OK",
							"message" "Data inserted with success.",
						}
					`)
					return
				}
			}
		}(i)
	}
	for _, data := range basicData {
		newBasicData := basic.NewBasic(data)
		// put data in channel
		dataChan <- newBasicData
	}

	close(dataChan)
	totalTime := time.Now().Sub(startTime)

	c.JSON(http.StatusOK, `
		{ 
			"status": "OK",
			"message": "Data inserted with success.",
			"totalTime":`+totalTime.String()+`
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
	c.JSON(http.StatusOK, len(data))
}
