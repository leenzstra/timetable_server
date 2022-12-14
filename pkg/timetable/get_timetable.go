package timetable

import (
	"net/url"
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/leenzstra/timetable_server/common/models"
	"github.com/leenzstra/timetable_server/common/responses"
	"github.com/leenzstra/timetable_server/common/utils"
)

func NewTimetableResponse(g *models.Timetable) (*responses.TimetableResponse, error) {

	subjects, err := utils.FetchGroupSubjectsFromTimetable(*g)
	if err != nil {
		return nil, err
	}

	sort.Slice(subjects, func(i, j int) bool {
		return subjects[i].Time < subjects[j].Time
	})

	return &responses.TimetableResponse{
		ID:      g.ID,
		GroupID: g.GroupID,
		Day:     g.Day,
		WeekNum: g.WeekNum,
		Table:   subjects,
	}, nil
}

// GetGroupTimetable godoc
// @Summary      Get group timetable
// @Description  Get group timetable by group name and type
// @ID get-group-timetable
// @Tags         timetable
// @Param group_name  path string true "Group Name"
// @Param timetable_type  path string true "timetable type"
// @Accept       json
// @Produce      json
// @Success      200  {object} responses.ResponseBase{data=[]responses.TimetableResponse}
// @Router       /timetable/timetables/{group_name}/{timetable_type} [get]
func (h handler) GetGroupTimetable(c *fiber.Ctx) error {
	groupName, err := url.QueryUnescape(c.Params("group_name"))
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	tableType, err := url.QueryUnescape(c.Params("timetable_type"))
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	// err = h.DB.Where(&models.Group{GroupName: groupName, TableKind: tableType}).First(&u).Error
	// if err != nil {
	// 	return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	// }

	// возможно GroupId
	timetable, err := h.DB.GetGroupTimetable(groupName, tableType)
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	tResponses := make([]*responses.TimetableResponse, 0)
	for _, t := range timetable {
		tr, err := NewTimetableResponse(t)
		if err != nil {
			return c.JSON(utils.WrapResponse(false, err.Error(), nil))
		}
		tResponses = append(tResponses, tr)
	}

	sort.Slice(tResponses, func(i, j int) bool {
		if tResponses[i].Day != tResponses[j].Day {
			return tResponses[i].Day < tResponses[j].Day
		}
		return tResponses[i].WeekNum < tResponses[j].WeekNum
	})

	return c.JSON(utils.WrapResponse(true, "", tResponses))
}
