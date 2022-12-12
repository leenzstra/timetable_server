package teachers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/leenzstra/timetable_server/common/models"
	"github.com/leenzstra/timetable_server/common/utils"
)

type TeachersResponse struct {
	Id        int            `json:"id"`
	Name       string         `json:"name"`
	Department string `json:"department"`
	Position   string         `json:"position"`
}

func NewTeachersResponse(g *models.Teacher) *TeachersResponse {
	var dep string
	if !g.Department.Valid {
		dep = ""
	} else {
		dep = g.Department.String
	}
	return &TeachersResponse{
		Id:        g.Id,
		Name:       g.FIO,
		Department: dep,
		Position:   g.Position,
	}
}

// GetTeachers godoc
// @Summary      Get teachers list
// @Description  Get teachers list
// @ID get-teachers
// @Tags         teachers
// @Param filter  path string false "Name filter"
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.ResponseBase{data=[]TeachersResponse}
// @Router       /teachers/{filter} [get]
func (h handler) GetTeachers(c *fiber.Ctx) error {
	var g []*models.Teacher

	filter, err := url.QueryUnescape(c.Params("filter", ""))
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	if filter == "undefined" {
		filter = ""
	}

	err = h.DB.Where("fio LIKE ?",filter+"%").Find(&g).Error
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	gResponses := make([]*TeachersResponse, 0)
	for _, gr := range g {
		tr := NewTeachersResponse(gr)
		gResponses = append(gResponses, tr)
	}

	return c.JSON(utils.WrapResponse(true, "", gResponses))
}
