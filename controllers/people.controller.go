package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/jackc/pgtype"
	"github.com/jackc/pgx/v5"
	db "goAPI/database/sqlc"
	"log/slog"
	"net/http"
	"strings"
)

type PeopleController struct {
	db *db.Queries
}

func NewPeopleController(db *db.Queries) *PeopleController {
	return &PeopleController{db}
}

type GetByPassportRequestParams struct {
	PassportSerie  string `binding:"required,numeric,len=4" db:"passport_serie"  form:"passportSerie"  json:"passportSerie"`
	PassportNumber string `binding:"required,numeric,len=6" db:"passport_number" form:"passportNumber" json:"passportNumber"`
}

// GetByPassport godoc
//
//	@Summary		Show a people
//	@Description	get people by passport details
//	@Tags			people
//	@Accept			json
//	@Produce		json
//	@Param			passportSerie	query		string	true	"Passport serie"
//	@Param			passportNumber	query		string	true	"Passport number"
//	@Success		200				{object}	db.Person
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/info [get]
func (pc PeopleController) GetByPassport(ctx *gin.Context) {
	passport := GetByPassportRequestParams{}
	slog.Debug("Parsing passport from request...")
	if err := ctx.ShouldBind(&passport); err != nil {
		slog.Debug("Invalid passport in request: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	args := &db.GetByPassportParams{
		PassportSerie:  passport.PassportSerie,
		PassportNumber: passport.PassportNumber,
	}

	slog.Debug("Getting people by passport from database...")
	people, err := pc.db.GetByPassport(ctx, *args)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		slog.Error("Cannot get people by passport:", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if errors.Is(err, pgx.ErrNoRows) {
		slog.Debug("People not found")
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, people)
}

type GetMultipleParams struct {
	ID             *string `db:"id"              form:"id"             json:"id"             binding:"omitempty,uuid"          format:"uuid" example:"00000000-0000-0000-0000-000000000000"`
	Name           *string `db:"name"            form:"name"           json:"name"           binding:"omitempty,min=3"                       example:"Иван"`
	Surname        *string `db:"surname"         form:"surname"        json:"surname"        binding:"omitempty,min=4"                       example:"Иванов"`
	Patronymic     *string `db:"patronymic"      form:"patronymic"     json:"patronymic"     binding:"omitempty,min=3"                       example:"Иванович"`
	Address        *string `db:"address"         form:"address"        json:"address"        binding:"omitempty,min=10"                      example:"3-й Автозаводский проезд, вл13, Москва, 115280"`
	PassportSerie  *string `db:"passport_serie"  form:"passportSerie"  json:"passportSerie"  binding:"omitempty,numeric,len=4"               example:"1234"`
	PassportNumber *string `db:"passport_number" form:"passportNumber" json:"passportNumber" binding:"omitempty,numeric,len=6"               example:"567756"`
	Offset         *int64  `db:"offset"          form:"offset"         json:"offset"         binding:"omitempty,min=0"                       example:"0"`
	Limit          *int64  `db:"limit"           form:"limit"          json:"limit"          binding:"omitempty,min=1"                       example:"10"`
}

// GetMultiple godoc
//
//	@Summary		Show a multiple full people
//	@Description	get people by multiple filters
//	@Tags			people
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	false	"Id"
//	@Param			surname			query		string	false	"Surname"
//	@Param			patronymic		query		string	false	"Patronymic"
//	@Param			address			query		string	false	"Address"
//	@Param			passportSerie	query		string	false	"Passport serie"
//	@Param			passportNumber	query		string	false	"Passport number"
//	@Success		200				{object}	[]db.Person
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/people [get]
func (pc PeopleController) GetMultiple(ctx *gin.Context) {
	filters := GetMultipleParams{}
	slog.Debug("Parsing people filters from json...")
	if err := ctx.ShouldBind(&filters); err != nil {
		slog.Debug("Invalid people filters in request: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id *uuid.UUID = nil
	if filters.ID != nil {
		parsedID, err := uuid.Parse(*filters.ID)
		if err != nil {
			slog.Debug("Invalid people id in url: ", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id = &parsedID
	}

	args := &db.GetMultipleParams{
		ID:             id,
		Surname:        filters.Surname,
		Patronymic:     filters.Patronymic,
		Address:        filters.Address,
		PassportSerie:  filters.PassportSerie,
		PassportNumber: filters.PassportNumber,
		Offset:         filters.Offset,
		Limit:          filters.Limit,
	}

	slog.Debug("Getting people from database...")
	people, err := pc.db.GetMultiple(ctx, *args)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		slog.Error("Cannot get people with filters:", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if errors.Is(err, pgx.ErrNoRows) {
		slog.Debug("People not found")
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, people)
}

type DeleteParams struct {
	ID             *string `db:"id"              form:"id"             json:"id"             binding:"omitempty,uuid"`
	Surname        *string `db:"surname"         form:"surname"        json:"surname"        binding:"omitempty,min=3"`
	Patronymic     *string `db:"patronymic"      form:"patronymic"     json:"patronymic"     binding:"omitempty,min=3"`
	Address        *string `db:"address"         form:"address"        json:"address"        binding:"omitempty,min=10"`
	PassportSerie  *string `db:"passport_serie"  form:"passportSerie"  json:"passportSerie"  binding:"omitempty,numeric,len=4"`
	PassportNumber *string `db:"passport_number" form:"passportNumber" json:"passportNumber" binding:"omitempty,numeric,len=6"`
}

// Delete godoc
//
//	@Summary		Delete people
//	@Description	delete people by multiple filters
//	@Tags			people
//	@Accept			json
//	@Produce		json
//	@Param			id				path	string	false	"Id"
//	@Param			surname			query	string	false	"Surname"
//	@Param			patronymic		query	string	false	"Patronymic"
//	@Param			address			query	string	false	"Address"
//	@Param			passportSerie	query	string	false	"Passport serie"
//	@Param			passportNumber	query	string	false	"Passport number"
//	@Success		200
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/people [delete]
func (pc PeopleController) Delete(ctx *gin.Context) {
	filters := DeleteParams{}
	slog.Debug("Parsing people delete filters from request...")
	if err := ctx.ShouldBind(&filters); err != nil {
		slog.Debug("Invalid people delete filters in request: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if filters.ID == nil && filters.Surname == nil && filters.Patronymic == nil &&
		filters.Address == nil &&
		filters.PassportSerie == nil &&
		filters.PassportNumber == nil {
		slog.Debug("No filters specified, aborting")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "no filters specified"})
		return
	}

	var id *uuid.UUID = nil
	if filters.ID != nil {
		parsedID, err := uuid.Parse(*filters.ID)
		if err != nil {
			slog.Debug("Invalid people id in url: ", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id = &parsedID
	}

	args := &db.DeleteParams{
		ID:             id,
		Surname:        filters.Surname,
		Patronymic:     filters.Patronymic,
		Address:        filters.Address,
		PassportSerie:  filters.PassportSerie,
		PassportNumber: filters.PassportNumber,
	}
	slog.Debug("Deleting people from database...")
	err := pc.db.Delete(ctx, *args)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		slog.Error("Cannot get people with filters: ", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if errors.Is(err, pgx.ErrNoRows) {
		slog.Debug("People not found")
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusOK)
}

type EditUrlParams struct {
	Id string `db:"id" form:"id" uri:"id" json:"id" binding:"required,uuid"`
}

type EditParams struct {
	Name           *string `db:"name"            form:"name"           json:"name"           binding:"omitempty,min=3"         minLength:"3"  example:"Иван"`
	Surname        *string `db:"surname"         form:"surname"        json:"surname"        binding:"omitempty,min=3"         minLength:"3"  example:"Иванов"`
	Patronymic     *string `db:"patronymic"      form:"patronymic"     json:"patronymic"     binding:"omitempty,min=3"         minLength:"3"  example:"Иванович"`
	Address        *string `db:"address"         form:"address"        json:"address"        binding:"omitempty,min=10"        minLength:"10" example:"3-й Автозаводский проезд, вл13, Москва, 115280"`
	PassportSerie  *string `db:"passport_serie"  form:"passportSerie"  json:"passportSerie"  binding:"omitempty,numeric,len=4" minLength:"4"  example:"1234"                                           maxLength:"4"`
	PassportNumber *string `db:"passport_number" form:"passportNumber" json:"passportNumber" binding:"omitempty,numeric,len=6" minLength:"6"  example:"123456"                                         maxLength:"6"`
}

// Edit godoc
//
//	@Summary		Edit people
//	@Description	edit people by id
//	@Tags			people
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string					true	"Id"
//	@Param			people	body		controllers.EditParams	true	"People data to edit"
//	@Success		200		{object}	db.Person
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/people/{id} [patch]
func (pc PeopleController) Edit(ctx *gin.Context) {
	editPeopleUrl := EditUrlParams{}
	slog.Debug("Parsing people id from url...")
	if err := ctx.BindUri(&editPeopleUrl); err != nil {
		slog.Debug("Invalid people id in url: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(editPeopleUrl.Id)
	if err != nil {
		slog.Debug("Invalid people id in url: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	editPeople := EditParams{}
	slog.Debug("Parsing people data to edit from request...")
	if err := ctx.ShouldBind(&editPeople); err != nil {
		slog.Debug("Invalid people data to edit in request: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if editPeople.Surname == nil && editPeople.Patronymic == nil && editPeople.Address == nil &&
		editPeople.PassportSerie == nil && editPeople.PassportNumber == nil {
		slog.Debug("No data to edit specified, aborting")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "no data to edit specified"})
		return
	}
	args := &db.EditParams{
		ID:             id,
		Name:           editPeople.Name,
		Surname:        editPeople.Surname,
		Patronymic:     editPeople.Patronymic,
		Address:        editPeople.Address,
		PassportSerie:  editPeople.PassportSerie,
		PassportNumber: editPeople.PassportNumber,
	}
	slog.Debug("Editing people in database...")
	people, err := pc.db.Edit(ctx, *args)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		slog.Error("Cannot edit people: ", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if errors.Is(err, pgx.ErrNoRows) {
		slog.Debug("People to edit not found")
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, people)
}

type CreateParams struct {
	PassportNumber string `binding:"required,passportNumber" db:"passport_number" form:"passportNumber" json:"passportNumber"`
}

// Create godoc
//
//	@Summary		Create people
//	@Description	create people by passport number
//	@Tags			people
//	@Accept			json
//	@Produce		json
//	@Param			passport	query		controllers.CreateParams	true	"Passport data"
//	@Success		200			{object}	db.Person
//	@Failure		400
//	@Failure		500
//	@Router			/people [post]
func (pc PeopleController) Create(ctx *gin.Context) {
	passport := CreateParams{}
	slog.Debug("Parsing people from request...")
	if err := ctx.ShouldBindJSON(&passport); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	passportData := strings.Fields(passport.PassportNumber)
	args := &db.CreateParams{
		PassportSerie:  passportData[0],
		PassportNumber: passportData[1],
	}
	slog.Debug("Creating people in database...")
	people, err := pc.db.Create(ctx, *args)

	if err != nil {
		slog.Error("Cannot create people:", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, people)
}
