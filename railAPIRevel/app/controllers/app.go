package controllers

import (
	"github.com/revel/revel"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	*revel.Controller
}

type TrainResource struct {
	ID              int    `json:"id"`
	DriverName      string `json:"driver_name"`
	OperatingStatus bool   `json:"operating_status"`
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) GetTrain() revel.Result {
	// curl -X GET "http://127.0.0.1:8008/v1/trains/1"
	var train TrainResource
	id := c.Params.Route.Get("train-id")
	train.ID, _ = strconv.Atoi(id)
	train.DriverName = "Logan"
	train.OperatingStatus = true
	c.Response.Status = http.StatusOK
	return c.RenderJSON(train)
}

func (c App) CreateTrain() revel.Result {
	// curl -X POST http://127.0.0.1:8008/v1/trains -H 'cache-control: no-cache' -H 'content-type: application/json' -d '{"driver_name":"Magneto", "operating_status": true}'
	var train TrainResource
	c.Params.BindJSON(&train)
	train.ID = 2
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(train)
}

func (c App) RemoveTrain() revel.Result {
	// curl -X DELETE "http://127.0.0.1:8008/v1/trains/1"
	id := c.Params.Route.Get("train-id")
	log.Println("Successfully deleted the resource:", id)
	c.Response.Status = http.StatusOK
	return c.RenderText("")
}
