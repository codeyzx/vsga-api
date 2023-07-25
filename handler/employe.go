package handler

import (
	"log"
	"vsga-api/database"
	"vsga-api/model/entity"
	"vsga-api/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// EmployeHandlerGetAll handles the GET request to retrieve all employees.
// @Summary Get all employees
// @Description Retrieve all employees from the database.
// @Tags Employees
// @Produce json
// @Success 200 {array} entity.Employe
// @Failure 500
// @Router /employe [get]
func EmployeHandlerGetAll(ctx *fiber.Ctx) error {
	var Employes []entity.Employe

	result := database.DB.Debug().Find(&Employes)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(Employes)
}

// EmployeHandlerGetById handles the GET request to retrieve an employee by ID.
// @Summary Get employee by ID
// @Description Retrieve an employee by its unique ID from the database.
// @Tags Employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} entity.Employe
// @Failure 404
// @Router /employe/{id} [get]
func EmployeHandlerGetById(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")

	var employe entity.Employe
	err := database.DB.First(&employe, "id = ?", ID).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Employe not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    employe,
	})
}

// EmployeHandlerCreate handles the POST request to create a new employee.
// @Summary Create an employee
// @Description Create a new employee and store it in the database.
// @Tags Employees
// @Accept json
// @Produce json
// @Param employe body request.EmployeRequest true "Employee details"
// @Success 200 {object} entity.Employe
// @Failure 400
// @Failure 500
// @Router /employe [post]
func EmployeHandlerCreate(ctx *fiber.Ctx) error {
	Employe := new(request.EmployeRequest)
	if err := ctx.BodyParser(Employe); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(Employe)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newEmploye := entity.Employe{
		Name:     Employe.Name,
		Position: Employe.Position,
		Salary:   Employe.Salary,
	}

	errCreateEmploye := database.DB.Create(&newEmploye).Error
	if errCreateEmploye != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"messaage": "success",
		"data":     newEmploye,
	})
}

// EmployeHandlerUpdate handles the PUT request to update an existing employee.
// @Summary Update an employee
// @Description Update an existing employee in the database based on its unique ID.
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employe body request.EmployeRequest true "Employee details"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /employe/{id} [put]
func EmployeHandlerUpdate(ctx *fiber.Ctx) error {
	EmployeRequest := new(request.EmployeRequest)
	if err := ctx.BodyParser(EmployeRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	errUpdate := database.DB.Where(
		"id = ?", ctx.Params("id"),
	).Updates(
		entity.Employe{
			Name:     EmployeRequest.Name,
			Position: EmployeRequest.Position,
			Salary:   EmployeRequest.Salary,
		},
	).Error

	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

// EmployeHandlerDelete handles the DELETE request to delete an existing employee.
// @Summary Delete an employee
// @Description Delete an existing employee from the database based on its unique ID.
// @Tags Employees
// @Param id path int true "Employee ID"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /employe/{id} [delete]
func EmployeHandlerDelete(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	var Employe entity.Employe

	err := database.DB.Debug().First(&Employe, "id=?", ID).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Employe not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&Employe).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Employe was deleted",
	})
}
