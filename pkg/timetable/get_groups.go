package timetable

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leenzstra/timetable_server/common/models"
	"github.com/leenzstra/timetable_server/common/utils"
)

type GroupResponse struct {
	ID        uint   `json:"id"`
	Faculty   string `json:"faculty"`
	Direction string `json:"direction"`
	GroupName string `json:"group_name"`
}

func NewGroupResponse(g *models.Group) *GroupResponse {
	return &GroupResponse{
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
// @Success      200  {object}  models.ResponseBase{data=[]GroupResponse}
// @Router       /timetable/groups/ [get]
func (h handler) GetGroups(c *fiber.Ctx) error {
	var g []*models.Group

	err := h.DB.Find(&g).Group("GroupName").Error
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	gResponses := make([]*GroupResponse, 0)
	for _, gr := range g {
		tr := NewGroupResponse(gr)
		gResponses = append(gResponses, tr)
	}

	return c.JSON(utils.WrapResponse(true, "", gResponses))
}
