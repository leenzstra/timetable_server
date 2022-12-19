package teachers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/leenzstra/timetable_server/common/models"
	"github.com/leenzstra/timetable_server/common/responses"
	"github.com/leenzstra/timetable_server/common/utils"
)

func NewTeachersResponse(g *models.Teacher) *responses.TeachersResponse {
	var dep string
	if !g.Department.Valid {
		dep = ""
	} else {
		dep = g.Department.String
	}
	return &responses.TeachersResponse{
		Id:         g.Id,
		Name:       g.FIO,
		Department: dep,
		Position:   g.Position,
		ImageUrl: g.ImageUrl,
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
// @Success      200  {object}  responses.ResponseBase{data=[]responses.TeachersResponse}
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

	g, err = h.DB.GetTeachers(filter)
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	gResponses := make([]*responses.TeachersResponse, 0)
	for _, gr := range g {
		tr := NewTeachersResponse(gr)
		gResponses = append(gResponses, tr)
	}

	return c.JSON(utils.WrapResponse(true, "", gResponses))
}
