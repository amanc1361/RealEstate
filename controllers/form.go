package controllers

import (
	"net/http"
	"realstate/models"
	"realstate/repository"
	"realstate/util"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2/bson"
)

type FormController interface {
	CreateForm(ctx *fiber.Ctx) error
	GetForms(cts *fiber.Ctx) error
	GetForm(ctx *fiber.Ctx) error
	DeleteForm(ctx *fiber.Ctx) error
	UpdateForm(ctx *fiber.Ctx) error
}

type formController struct {
	form repository.FormRepository
}

func NewFormController(formrepo repository.FormRepository) FormController {
	return &formController{formrepo}
}

// Get Froms ... Get a new Froms
// @Summary  Get Forms
// @Description Get Forms
// @Tags Froms
// @Success 200 {array} models.Froms
// @Failure 404 {object} object
// @Router /forms/ [get]
func (r *formController) GetForms(ctx *fiber.Ctx) error {
	forms, err := r.form.GetForms()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusOK).JSON(util.NewRresult(forms))
}

func (r *formController) CreateForm(ctx *fiber.Ctx) error {
	var form models.Form
	err := ctx.BodyParser(&form)

	if err != nil {
		return ctx.Status(http.StatusBadGateway).JSON(util.NewJError(err))
	}

	form.Updateid()
	form.Id = bson.NewObjectId()

	err = r.form.SaveForm(&form)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
	}
	return ctx.Status(http.StatusCreated).JSON(util.NewRresult(form))
}

// Get From ... Get a new Froms
// @Summary  Get Form
// @Description Get Form
// @Tags Froms
// @Success 200 {object} models.From
// @Failure 404 {object} object
// @Router /forms/ [get]
func (r *formController) GetForm(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	form, err := r.form.GetForm(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	return ctx.Status(http.StatusOK).JSON(util.NewRresult(form))

}

// Delete From ... Delete a Form
// @Summary  Delete Form
// @Description Delete Form
// @Tags Froms
// @Success 200 {object} models.From
// @Failure 404 {object} object
// @Router /forms/ [Delete]
func (r *formController) DeleteForm(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := r.form.DeleteForm(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrNotFound))
	}
	return ctx.Status(http.StatusOK).JSON(util.NewRresult(util.SuccessDelete))
}

// Delete From ... Delete a Form
// @Summary  Delete Form
// @Description Delete Form
// @Tags Froms
// @Success 200 {object} models.From
// @Failure 404 {object} object
// @Router /forms/ [Delete]
func (r *formController) UpdateForm(ctx *fiber.Ctx) error {
	var form models.Form
	id := ctx.Params("id")
	err := ctx.BodyParser(&form)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrInvalidCredentials))
	}
	err = r.form.UpdateForm(id, &form)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrInvalidCredentials))
	}
	return ctx.Status(http.StatusOK).JSON(util.NewRresult(util.SuccessUpdate))
}
