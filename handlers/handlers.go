package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"projeto-transportadora/db/repository"
	"projeto-transportadora/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllClient(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusFound, gin.H{
			"message": "invalidId",
		})
	}
	res, err := repository.GetAllClient(int64(resp))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error GetAllClient",
		})
		return
	}
	if res == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "pushy user",
		})
	} else {
		ctx.JSON(http.StatusOK, res)
	}

}

func GetConsultClientID(ctx *gin.Context) {
	res, err := repository.GetClient()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid Id request",
		})
		return
	}
	if res == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "pushy user",
		})
	} else {
		ctx.JSON(http.StatusOK, res)
	}

}

func GetInsertCLientID(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to insert into the database",
		})
		return
	}
	var client models.Info
	if err := json.Unmarshal([]byte(body), &client); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse JSON",
		})
		return
	}

	res, err := repository.InsertClientList(client)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to insert batabase",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Sucess": res,
	})

}
func DeleteClientId(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error GetDeleteClientId",
		})
		return
	}
	res, err := repository.DeleteClientId(int64(resp))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "error delete client",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Sucess Client delete": res,
	})
}

func UpdateCleint(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to insert into the database",
		})
		return
	}
	var client models.Cliente
	if err := json.Unmarshal(body, &client); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse JSON",
		})
		return
	}

	res, err := repository.UpdateCleintId(client.Id, models.Cliente{Id: client.Id,
		Data: models.Info{
			Name:    client.Data.Name,
			Address: client.Data.Address,
			Phone:   client.Data.Phone}})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "already registered user phone number",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Sucess": res,
	})
}
