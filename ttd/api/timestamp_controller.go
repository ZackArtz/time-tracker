package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"github.com/zackartz/ttd/models"
	"github.com/zackartz/ttd/utils"
)

var activeTimestamps []models.Timestamp

func (s *Server) GetAllTimestamps(ctx *fiber.Ctx) {
	ts := models.Timestamp{}
	timestamps, err := ts.GetAllTimestamps(s.DB)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, err)
		return
	}
	utils.JSON(ctx, http.StatusOK, timestamps)
}

func (s *Server) CreateTimestamp(ctx *fiber.Ctx) {
	ts := &models.Timestamp{}
	if err := ctx.BodyParser(ts); err != nil {
		utils.Error(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	err := ts.Validate()
	if err != nil {
		utils.Error(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	guid := uuid.New()
	ts.StartTime = time.Now()
	ts.UUID = guid.String()
	activeTimestamps = append(activeTimestamps, *ts)
	utils.JSON(ctx, http.StatusCreated, ts)
}

func (s *Server) GetTimestampsByProject(ctx *fiber.Ctx) {
	ts := models.Timestamp{}
	project := ctx.Params("project")
	timestamps, err := ts.GetAllTimestampsByProject(s.DB, project)
	if err != nil {
		utils.Error(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	utils.JSON(ctx, http.StatusOK, timestamps)
}

func (s *Server) GetAllActiveTimestamps(ctx *fiber.Ctx) {
	utils.JSON(ctx, http.StatusOK, activeTimestamps)
}

func (s *Server) GetTimestampByUUID(ctx *fiber.Ctx) {
	uuid := ctx.Params("uuid")
	for _, ts := range activeTimestamps {
		if ts.UUID == uuid {
			utils.JSON(ctx, http.StatusOK, ts)
			return
		}
	}
	ts := &models.Timestamp{}
	ts, err := ts.GetTimestampByID(s.DB, uuid)
	if err != nil {
		utils.Error(ctx, http.StatusNotFound, errors.New("could not find timestamp"))
		return
	}
	utils.JSON(ctx, http.StatusOK, ts)
}

func (s *Server) EndTimestampByUUID(ctx *fiber.Ctx) {
	uuid := ctx.Params("uuid")
	for i, ts := range activeTimestamps {
		if ts.UUID == uuid {
			ts.EndTime = time.Now()
			ts, err := ts.Create(s.DB)
			if err != nil {
				utils.Error(ctx, http.StatusBadRequest, err)
				return
			}
			activeTimestamps = RemoveIndex(activeTimestamps, i)
			utils.JSON(ctx, http.StatusOK, ts)
		}
	}

}

func (s *Server) DeleteTimestamp(ctx *fiber.Ctx) {
	ts := new(models.Timestamp)

	tid := ctx.Params("id")
	_, err := ts.Delete(s.DB, tid)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.Set("Entity", fmt.Sprintf("%s", tid))
	utils.JSON(ctx, http.StatusNoContent, "")
}

func RemoveIndex(s []models.Timestamp, index int) []models.Timestamp {
	return append(s[:index], s[index+1:]...)
}
