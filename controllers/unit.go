package controllers

import (
	"net/http"
	"realstate/models"
	"realstate/repository"
	"realstate/util"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UnitController interface {
	CreateUnit(ctx *fiber.Ctx) error
	UpdateUnit(ctx *fiber.Ctx) error
	GetUnit(ctx *fiber.Ctx) error
	GetUnits(ctx *fiber.Ctx) error
	DeleteUnit(ctx *fiber.Ctx) error
}

type unitController struct {
	unit repository.UnitRepository
}

func NewUnitController(unitrepo repository.UnitRepository) unitController {
	return unitController{unitrepo}
}

// Create Unit ... Create a new Unit
// @Summary Create a new Unit
// @Description Create a new Unit
// @Tags Unit
// @Success 200 {object} models.Unit
// @Failure 404 {object} object
// @Router /Unit [post]
func (c *unitController) CreateUnit(ctx *fiber.Ctx) error {
	var unit models.Unit
	err := ctx.BodyParser(&unit)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	exists, err := c.unit.GetUnitByName(unit.Name)

	if err == mgo.ErrNotFound {
		if strings.TrimSpace(unit.Name) == "" {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(util.ErrEmptyName))
		}
		unit.CreatedAt = time.Now()
		unit.UpdatedAt = time.Now()
		unit.Id = bson.NewObjectId()
		err = c.unit.SaveUnit(&unit)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		} else {
			return ctx.
				Status(http.StatusCreated).
				JSON(unit)
		}

	}
	if exists != nil {
		err = util.ErrNameAlreadyExists
	}
	return ctx.
		Status(http.StatusBadRequest).
		JSON(util.NewJError(err))
}

// update Unit ... update Unit
// @Summary update Unit
// @Description update Assginmenttype
// @Tags Unit
// @Success 200 {object} models.Unit
// @Failure 404 {object} object
// @Router /Unit/ [put]
func (r *unitController) UpdateUnit(ctx *fiber.Ctx) error {
	var unit models.Unit
	err := ctx.BodyParser(&unit)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}

	dbunit, err := r.unit.GetUnitById(unit.Id.Hex())

	if err != nil {

		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(util.ErrNotFound))
	}
	exists, err := r.unit.GetUnitByName(unit.Name)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.ErrEmptyName)
	}

	if strings.TrimSpace(exists.Name) != "" {

		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNameAlreadyExists))
	}

	dbunit.UpdatedAt = time.Now()
	dbunit.Name = unit.Name
	err = r.unit.UpdateUnit(dbunit)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(unit)

}

// GetUnit ... Get Unit by id
// @Summary Get Unit by id
// @Description Get Unit by id
// @Tags Unit
// @Success 200 {object} models.Unit
// @Param id path string true "Item ID"
// @Failure 400 {object} object
// @Router /unit/id [get]
func (r *unitController) GetUnit(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	unit, err := r.unit.GetUnitById(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(unit)
}

// GetUnits ... Get All Unit
// @Summary Get All Unit
// @Description Get All Unit
// @Tags Unit
// @Success 200 {array} models.Unit
// @Failure 400 {object} object
// @Router /units [get]
func (r *unitController) GetUnits(ctx *fiber.Ctx) error {

	units, err := r.unit.GetUnitAll()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(units)
}

// DeleteUnit ... Delete Unit by id
// @Summary Delete Unit by id
// @Description Delete Unit by id
// @Tags Unit
// @Success 200 {object} models.Unit
// @Param id path string true "Item ID"
// @Failure 400 {object} object
// @Router /unit/id [delete]
func (r *unitController) DeleteUnit(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	err := r.unit.DeleteUnit(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(util.SuccessDelete)
}
