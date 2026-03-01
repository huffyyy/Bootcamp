package createemployee

import (
	"context"
	"errors"
	"net/http"
	"pmo/internal/pkg/response"
	"pmo/internal/pkg/storage"
	"pmo/services/hr-service/features/employee/shared/dto"
	"pmo/services/hr-service/features/employee/shared/repository"
	"pmo/services/hr-service/features/employee/shared/validators"

	"pmo/services/hr-service/internal/domain/models"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateEmployeeHandler struct {
	repo      repository.EmployeeRepository
	validator *validators.EmployeeValidator
	storage   *storage.StorageService
}

func NewCreateEmployeeHandler(repo repository.EmployeeRepository, v *validators.EmployeeValidator, storage *storage.StorageService) *CreateEmployeeHandler {
	return &CreateEmployeeHandler{repo: repo, validator: v, storage: storage}
}
func (h *CreateEmployeeHandler) Handle(c *gin.Context) {

	// Parse multipart form (max 10MB total)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse[any]("Failed to parse form: "+err.Error()))
		return
	}

	// Parse form values
	var req dto.CreateEmployeeRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse[any]("Invalid request format: "+err.Error()))
		return
	}

	// Get photos from form
	form, _ := c.MultipartForm()
	req.Photos = form.File["photos"]

	// Validate
	if errs := h.validator.ValidateCreateEmployee(&req); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, response.ValidationError[any](errs))
		return
	}
	// Parse hire date
	hireDate, err := time.Parse("2006-01-02", req.HireDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse[any]("Invalid hire date format"))
		return
	}
	// Create command
	cmd := NewCreateEmployeeCommand(
		req.FirstName,
		req.LastName,
		req.Email,
		req.PhoneNumber,
		hireDate,
		req.JobID,
		req.Salary,
		req.ManagerID,
		req.DepartmentID,
		req.Photos,
	)

	// Execute in transaction
	ctx := context.Background()
	employee := cmd.ToModel()
	// Save photos jika ada
	if len(req.Photos) > 0 {
		photos := make([]models.EmployeePhoto, 0, len(req.Photos))
		for i, file := range req.Photos {
			// Save file to storage
			fileURL, fileSize, err := h.storage.SaveFile(file, "employees/photos")
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.ErrorResponse[any]("Failed to save photo: "+err.Error()))
				return
			}

			// Create photo record
			fileName := file.Filename
			fileType := file.Header.Get("Content-Type")
			isPrimary := i == 0 // Photo pertama sebagai primary
			photo := models.EmployeePhoto{
				FileName:  &fileName,
				FileSize:  (&fileSize),
				FileType:  &fileType,
				FileURL:   &fileURL,
				IsPrimary: &isPrimary,
			}
			photos = append(photos, photo)
		}
		employee.Photos = photos
	}

	// Save to database
	if err := h.repo.Create(ctx, employee); err != nil {

		// Rollback files if database fails
		if len(employee.Photos) > 0 {
			for _, photo := range employee.Photos {
				if photo.FileURL != nil {
					h.storage.DeleteFile(*photo.FileURL)
				}
			}
		}
		if errors.Is(err, repository.ErrDuplicateEmail) {
			c.JSON(http.StatusConflict, response.ErrorResponse[any]("Email already exists"))
			return
		}
		c.JSON(http.StatusInternalServerError, response.ErrorResponse[any]("Failed to create employee: "+err.Error()))
		return
	}

	// Return response
	resp := dto.ToEmployeeResponse(employee)
	c.JSON(http.StatusCreated, response.SuccessResponse(resp, "Employee created successfully"))
}
