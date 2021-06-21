package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/zackartz/ttd/ent/timestamp"
	"github.com/zackartz/ttd/models"
	"github.com/zackartz/ttd/utils"
)

var (
	c context.Context
)

func (s *Server) GetAllTimestamps(ctx *fiber.Ctx) error {
	timestamp, err := s.client.Timestamp.Query().All(c)
	if err != nil {
		return utils.Error(ctx, http.StatusInternalServerError, err)
	}
	return utils.JSON(ctx, http.StatusOK, timestamp)
}

func (s *Server) CreateTimestamp(ctx *fiber.Ctx) error {
	ts := &models.Timestamp{}
	if err := ctx.BodyParser(ts); err != nil {
		return utils.Error(ctx, http.StatusUnprocessableEntity, err)
	}
	timestamp, err := s.client.Timestamp.Create().
		SetNillableCategory(&ts.Category).
		SetProject(ts.Project).
		SetEndTime(time.Time{}).
		SetNillableComment(&ts.Comment).
		Save(c)
	if err != nil {
		return utils.Error(ctx, http.StatusBadRequest, err)
	}
	return utils.JSON(ctx, http.StatusCreated, timestamp)
}

func (s *Server) GetTimestampsByProject(ctx *fiber.Ctx) error {
	project := ctx.Params("project")
	timestamps, err := s.client.Timestamp.Query().
		Where(timestamp.Project(project)).All(c)
	if err != nil {
		return utils.Error(ctx, http.StatusInternalServerError, err)
	}
	return utils.JSON(ctx, http.StatusOK, timestamps)
}

func (s *Server) GetAllActiveTimestamps(ctx *fiber.Ctx) error {
	timestamps, err := s.client.Timestamp.Query().
		Where(timestamp.Active(true)).
		All(c)
	if err != nil {
		return utils.Error(ctx, http.StatusUnprocessableEntity, err)
	}
	return utils.JSON(ctx, http.StatusOK, timestamps)
}

func (s *Server) GetTimestampByUUID(ctx *fiber.Ctx) error {
	id := ctx.Params("uuid")
	uid, err := uuid.Parse(id)
	if err != nil {
		return utils.Error(ctx, http.StatusBadRequest, errors.New("malformed ID"))
	}
	ts, err := s.client.Timestamp.Get(c, uid)
	if err != nil {
		return utils.Error(ctx, http.StatusUnprocessableEntity, err)
	}
	return utils.JSON(ctx, http.StatusOK, ts)
}

func (s *Server) EndTimestampByUUID(ctx *fiber.Ctx) error {
	id := ctx.Params("uuid")
	uid, err := uuid.Parse(id)
	if err != nil {
		return utils.Error(ctx, http.StatusBadRequest, err)
	}
	ts, err := s.client.Timestamp.Update().
		Where(timestamp.ID(uid)).
		SetEndTime(time.Now()).
		SetActive(false).
		Save(c)
	if err != nil {
		return utils.Error(ctx, http.StatusBadRequest, err)
	}
	return utils.JSON(ctx, http.StatusOK, ts)
}

func (s *Server) DeleteTimestamp(ctx *fiber.Ctx) error {
	tid := ctx.Params("id")
	uid, err := uuid.Parse(tid)
	if err != nil {
		return utils.Error(ctx, http.StatusBadRequest, err)
	}
	ts, err := s.client.Timestamp.Delete().Where(timestamp.ID(uid)).Exec(c)
	if err != nil {
		return utils.Error(ctx, http.StatusInternalServerError, err)
	}
	ctx.Set("Entity", tid)
	return utils.JSON(ctx, http.StatusNoContent, ts)
}

func RemoveIndex(s []models.Timestamp, index int) []models.Timestamp {
	return append(s[:index], s[index+1:]...)
}
