package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/itlightning/dateparse"
	"github.com/jackc/pgx/v5"
	db "goAPI/database/sqlc"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type TasksController struct {
	db *db.Queries
}

func NewTasksController(db *db.Queries) *TasksController {
	return &TasksController{db}
}

type StartTaskParams struct {
	Name string `binding:"required,min=3" db:"name" form:"name" json:"name"`
}

// StartTask godoc
//
//	@Summary		Start task
//	@Description	start task by people id
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			people	path		string	true	"People id"
//	@Param			name	query		string	true	"Task name"
//	@Success		200		{object}	db.Task
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/people/{id}/task/start [post]
func (tc TasksController) StartTask(ctx *gin.Context) {
	urlParams := PeopleUrlParams{}
	slog.Debug("Parsing people id from request...")
	if err := ctx.ShouldBindUri(&urlParams); err != nil {
		slog.Debug("Invalid people peopleId in request: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	slog.Debug("Parsing people peopleId from params...")
	peopleId, err := uuid.Parse(urlParams.Id)
	if err != nil {
		slog.Debug("Invalid people peopleId in params: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startTask := StartTaskParams{}
	slog.Debug("Parsing task info from request...")
	if err := ctx.ShouldBind(&startTask); err != nil {
		slog.Debug("Invalid task info in request: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := &db.StartTaskParams{
		PeopleID: peopleId,
		Name:     startTask.Name,
	}

	slog.Debug("Saving task to database...")
	task, err := tc.db.StartTask(ctx, *args)
	if err != nil {
		slog.Error("Cannot save task to database: ", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// EndTask godoc
//
//	@Summary		End task
//	@Description	end task by people id
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			people	path		string	true	"People id"
//	@Success		200		{object}	db.Task
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/people/{id}/task/end [post]
func (tc TasksController) EndTask(ctx *gin.Context) {
	urlParams := PeopleUrlParams{}
	slog.Debug("Parsing people id from request...")
	if err := ctx.ShouldBindUri(&urlParams); err != nil {
		slog.Debug("Invalid people peopleId in request: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	peopleId, err := uuid.Parse(urlParams.Id)
	if err != nil {
		slog.Debug("Invalid people peopleId in request: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	slog.Debug("Ending task in database...")
	task, err := tc.db.EndTask(ctx, peopleId)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		slog.Error("Cannot end task in database: ", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if errors.Is(err, pgx.ErrNoRows) {
		slog.Debug("Task from people not found: ", err)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

type PeopleUrlParams struct {
	Id string `binding:"required,uuid" db:"id" uri:"id" form:"id" json:"id"`
}

type GetTasksByPeopleParams struct {
	PeriodStart *string `binding:"omitempty" db:"period_start" form:"periodStart" json:"periodStart"`
	PeriodEnd   *string `binding:"omitempty" db:"period_end"   form:"periodEnd"   json:"periodEnd"`
}

// GetTasksByPeople godoc
//
//	@Summary		Get people tasks
//	@Description	get tasks by people id and period of time
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			people		path		string	true	"People id"
//	@Param			periodStart	query		string	false	"Period start"
//	@Param			periodEnd	query		string	false	"Period end"
//	@Success		200			{object}	[]db.Task
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/people/{id}/tasks [get]
func (tc TasksController) GetTasksByPeople(ctx *gin.Context) {
	urlParams := PeopleUrlParams{}
	slog.Debug("Parsing people id from url...")
	if err := ctx.ShouldBindUri(&urlParams); err != nil {
		slog.Debug("Invalid people id in url: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := uuid.Parse(urlParams.Id)
	if err != nil {
		slog.Debug("Invalid people id in url: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	getTasksByPeople := GetTasksByPeopleParams{}
	slog.Debug("Parsing tasks period from request...")
	if err := ctx.ShouldBindJSON(&getTasksByPeople); err != nil && !errors.Is(err, io.EOF) {
		slog.Debug("Invalid period in request: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if errors.Is(err, io.EOF) {
		slog.Debug("No period in request: ", err)
	}

	var periodStart, periodEnd *time.Time
	if getTasksByPeople.PeriodStart != nil {
		parsedTime, err := dateparse.ParseAny(*getTasksByPeople.PeriodStart)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		periodStart = &parsedTime
	}

	if getTasksByPeople.PeriodEnd != nil {
		parsedTime, err := dateparse.ParseAny(*getTasksByPeople.PeriodEnd)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		periodEnd = &parsedTime
	}

	args := &db.GetTasksByPeopleParams{
		PeopleID:    id,
		PeriodStart: periodStart,
		PeriodEnd:   periodEnd,
	}

	slog.Debug("Getting tasks with filters from database...", args)

	tasks, err := tc.db.GetTasksByPeople(ctx, *args)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		slog.Error("Cannot get tasks with filters: ", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if errors.Is(err, pgx.ErrNoRows) || len(tasks) == 0 {
		slog.Debug("Tasks not found")
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}
