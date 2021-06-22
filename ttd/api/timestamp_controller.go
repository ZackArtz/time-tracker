package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zackartz/ttd/models"
	"github.com/zackartz/ttd/prisma/db"
	"github.com/zackartz/ttd/utils"
)

var (
	c context.Context
)

func (s *Server) GetAllTimestamps(ctx *fiber.Ctx) error {
	timestamp, err := s.client.Timestamp.FindMany().Exec(c)
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
	timestamp, err := s.client.Timestamp.CreateOne(
		db.Timestamp.Project.Set(ts.Project),
		db.Timestamp.EndTime.Set(time.Time{}),
		db.Timestamp.Category.SetIfPresent(&ts.Category),
		db.Timestamp.Comment.SetIfPresent(&ts.Comment),
	).Exec(c)
	if err != nil {
		return utils.Error(ctx, http.StatusBadRequest, err)
	}
	return utils.JSON(ctx, http.StatusCreated, timestamp)
}

func (s *Server) GetTimestampsByProject(ctx *fiber.Ctx) error {
	project := ctx.Params("project")
	timestamps, err := s.client.Timestamp.FindMany(
		db.Timestamp.Project.Equals(project),
	).Exec(c)
	if err != nil {
		return utils.Error(ctx, http.StatusInternalServerError, err)
	}
	return utils.JSON(ctx, http.StatusOK, timestamps)
}

func (s *Server) GetAllActiveTimestamps(ctx *fiber.Ctx) error {
	timestamps, err := s.client.Timestamp.FindMany().Exec(c)
	if err != nil {
		return utils.Error(ctx, http.StatusUnprocessableEntity, err)
	}
	return utils.JSON(ctx, http.StatusOK, timestamps)
}

func (s *Server) GetTimestampByUUID(ctx *fiber.Ctx) error {
	id := ctx.Params("uuid")
	ts, err := s.client.Timestamp.FindUnique(
		db.Timestamp.ID.Equals(id),
	).Exec(c)
	if err != nil {
		return utils.Error(ctx, http.StatusUnprocessableEntity, err)
	}
	return utils.JSON(ctx, http.StatusOK, ts)
}

func (s *Server) EndTimestampByUUID(ctx *fiber.Ctx) error {
	id := ctx.Params("uuid")
	ts, err := s.client.Timestamp.FindUnique(
		db.Timestamp.ID.Equals(id),
	).Exec(c)
	if err != nil {
		return utils.Error(ctx, http.StatusBadRequest, err)
	}
	if !ts.Active {
		return utils.Error(ctx, http.StatusBadRequest, errors.New("timestamp has already ended"))
	}
	_, err = s.client.Timestamp.FindUnique(
		db.Timestamp.ID.Equals(id),
	).Update(
		db.Timestamp.EndTime.Set(time.Now()),
		db.Timestamp.Active.Set(false),
	).Exec(c)
	if err != nil {
		return utils.Error(ctx, http.StatusBadRequest, err)
	}
	return utils.JSON(ctx, http.StatusOK, ts)
}

func (s *Server) DeleteTimestamp(ctx *fiber.Ctx) error {
	tid := ctx.Params("id")
	ts, err := s.client.Timestamp.FindUnique(
		db.Timestamp.ID.Equals(tid),
	).Delete().Exec(c)
	if err != nil {
		return utils.Error(ctx, http.StatusInternalServerError, err)
	}
	ctx.Set("Entity", tid)
	return utils.JSON(ctx, http.StatusNoContent, ts)
}
