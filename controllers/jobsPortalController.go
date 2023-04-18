package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jobsPortal/entities"
	"jobsPortal/helpers"
	"jobsPortal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JobsPortalController struct {
	DB                *gorm.DB
	JobsPortalService *services.JobsPortalService
}

type RequestUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (jc *JobsPortalController) CreateUser(ctx *gin.Context) {
	var requestUser RequestUser
	if err := ctx.ShouldBindJSON(&requestUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	user := entities.User{
		Username: requestUser.Username,
		Password: requestUser.Password,
	}

	result, err := jc.JobsPortalService.CreateUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func (jc *JobsPortalController) Login(ctx *gin.Context) {
	var requestUser RequestUser
	if err := ctx.ShouldBindJSON(&requestUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	user := entities.User{
		Username: requestUser.Username,
		Password: requestUser.Password,
	}

	result, err := jc.JobsPortalService.Login(user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	success := helpers.ComparePass([]byte(result.Password), []byte(requestUser.Password))
	if !success {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"error":   "unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(result.ID, result.Username)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})
}

func (jc *JobsPortalController) GetJobsList(ctx *gin.Context) {

	url := "http://dev3.dansmultipro.co.id/api/recruitment/positions.json/?"
	query := false
	if desc := ctx.Query("desc"); desc != "" {
		if !query {
			url += "&"
		}
		url += fmt.Sprintf("description=%s", desc)
	}

	if loc := ctx.Query("location"); loc != "" {
		if !query {
			url += "&"
		}
		url += fmt.Sprintf("location=%s", loc)
	}

	if ft := ctx.Query("full_time"); ft != "" {
		if !query {
			url += "&"
		}
		url += fmt.Sprintf("full_time=%s", ft)
	}

	if page := ctx.Query("page"); page != "" {
		if !query {
			url += "&"
		}
		url += fmt.Sprintf("page=%s", page)
	}

	res, err := http.Get(url)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	defer res.Body.Close()

	result := make([]map[string]interface{}, 0)
	err = json.Unmarshal(body, &result)
	log.Println(string(body))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)

}
func (jc *JobsPortalController) GetJobDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	url := fmt.Sprintf("http://dev3.dansmultipro.co.id/api/recruitment/positions/%s", id)

	res, err := http.Get(url)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	defer res.Body.Close()

	result := make(map[string]interface{}, 0)
	err = json.Unmarshal(body, &result)
	log.Println(string(body))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
