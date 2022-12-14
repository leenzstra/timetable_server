package timetable

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leenzstra/timetable_server/common/models"
	"github.com/leenzstra/timetable_server/common/responses"
	"github.com/leenzstra/timetable_server/common/utils"
)

func NewGroupResponse(g *models.Group) *responses.GroupResponse {
	return &responses.GroupResponse{
		ID:        g.ID,
		Faculty:   g.Faculty,
		Direction: g.Direction,
		GroupName: g.GroupName,
	}
}

// GetGroups godoc
// @Summary      Get groups list
// @Description  Get all groups list
// @ID get-groups
// @Tags         timetable
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.ResponseBase{data=[]responses.GroupResponse}
// @Router       /timetable/groups/ [get]
func (h handler) GetGroups(c *fiber.Ctx) error {
	groups, err := h.DB.GetGroups()
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	gResponses := make([]*responses.GroupResponse, 0)
	for _, gr := range groups {
		tr := NewGroupResponse(gr)
		gResponses = append(gResponses, tr)
	}

	return c.JSON(utils.WrapResponse(true, "", gResponses))
}
