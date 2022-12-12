package timetable

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/leenzstra/timetable_server/common/models"
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
// @Success      200  {object} models.ResponseBase{data=[]string}
// @Router       /timetable/types/{group_name} [get]
func (h handler) GetTimetableTypes(c *fiber.Ctx) error {
	var groupTypes []*models.Group

	groupName, err := url.QueryUnescape(c.Params("group_name"))
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	err = h.DB.Where(&models.Group{GroupName: groupName}).Find(&groupTypes).Error
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	types := make([]string, 0)
	for _, g := range groupTypes {
		types = append(types, g.TableKind)
	}

	return c.JSON(utils.WrapResponse(true, "", types))
}
