package v1

import (
	"net/http"

	"github.com/MuhammadyusufAdhamov/student/api/models"
	"github.com/MuhammadyusufAdhamov/student/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Router /student [post]
// @Summary Create a student
// @Description Create a student
// @Tags student
// @Accept json
// @Produce json
// @Param student body models.CreateStudent true "Student"
// @Success 201 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		req *models.CreateStudent
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = h.storage.Student().Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, models.ResponseOK{
		Message: "Successfully created",
	})
}

// @Router /student [get]
// @Summary Get students
// @Description Get students
// @Tags student
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param sezrch query string false "Search"
// @Success 200 {object} models.GetAllStudentsResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllStudents(c *gin.Context) {
	req, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.storage.Student().GetAll(&repo.GetAllStudentsParams{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getStudentsResponse(result))
}

func getStudentsResponse(data *repo.GetAllStudentsResult) *models.GetAllStudentsResponse {
	response := models.GetAllStudentsResponse{
		Students: make([]*models.Student, 0),
		Count: data.Count,
	}

	for _, user := range data.Students {
		u := parseStudentsModel(user)
		response.Students = append(response.Students, &u)
	}

	return &response
}

func parseStudentsModel(student *repo.Student) models.Student {
	return models.Student{
		ID:              student.ID,
		FirstName:       student.FirstName,
		LastName:        student.LastName,
		PhoneNumber:     &student.PhoneNumber,
		Email:           student.Email,
		Username:        student.Username,
		CreatedAt:       student.CreatedAt,
	}
}