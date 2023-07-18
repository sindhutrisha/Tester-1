package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sindhutrisha/Tester-1/tester-1/pkg/rest/server/models"
	"github.com/sindhutrisha/Tester-1/tester-1/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Tester1Controller struct {
	tester1Service *services.Tester1Service
}

func NewTester1Controller() (*Tester1Controller, error) {
	tester1Service, err := services.NewTester1Service()
	if err != nil {
		return nil, err
	}
	return &Tester1Controller{
		tester1Service: tester1Service,
	}, nil
}

func (tester1Controller *Tester1Controller) CreateTester1(context *gin.Context) {
	// validate input
	var input models.Tester1
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger tester1 creation
	if _, err := tester1Controller.tester1Service.CreateTester1(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Tester1 created successfully"})
}

func (tester1Controller *Tester1Controller) UpdateTester1(context *gin.Context) {
	// validate input
	var input models.Tester1
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger tester1 update
	if _, err := tester1Controller.tester1Service.UpdateTester1(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Tester1 updated successfully"})
}

func (tester1Controller *Tester1Controller) FetchTester1(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger tester1 fetching
	tester1, err := tester1Controller.tester1Service.GetTester1(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, tester1)
}

func (tester1Controller *Tester1Controller) DeleteTester1(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger tester1 deletion
	if err := tester1Controller.tester1Service.DeleteTester1(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Tester1 deleted successfully",
	})
}

func (tester1Controller *Tester1Controller) ListTester1s(context *gin.Context) {
	// trigger all tester1s fetching
	tester1s, err := tester1Controller.tester1Service.ListTester1s()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, tester1s)
}

func (*Tester1Controller) PatchTester1(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*Tester1Controller) OptionsTester1(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*Tester1Controller) HeadTester1(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
