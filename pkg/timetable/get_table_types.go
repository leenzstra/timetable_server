package timetable

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/leenzstra/timetable_server/common/responses"
	"github.com/leenzstra/timetable_server/common/utils"
)

// GetTimetableTypes godoc
// @Summary      Get timetable types
// @Description  Get timetable types by group name
// @ID get-timetable-types
// @Tags         timetable
// @Param group_name  path string true "Group Name"
// @Accept       json
// @Produce      json
// @Success      200  {object} responses.ResponseBase{data=[]responses.TimetableTypeResponse}
// @Router       /timetable/types/{group_name} [get]
func (h handler) GetTimetableTypes(c *fiber.Ctx) error {
	groupName, err := url.QueryUnescape(c.Params("group_name"))
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	groupTypes, err := h.DB.GetGroupTimetableTypes(groupName)
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	types := make([]responses.TimetableTypeResponse, 0)
	for _, g := range groupTypes {
		presentationType := utils.StringBetween(g.TableKind, "(", ")", g.TableKind)
		types = append(types, responses.TimetableTypeResponse{Name: g.TableKind, Presentation: presentationType})
	}

	return c.JSON(utils.WrapResponse(true, "", types))
}
