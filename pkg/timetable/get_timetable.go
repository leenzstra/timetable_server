package timetable

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/leenzstra/timetable_server/common/constant"
	"github.com/leenzstra/timetable_server/common/models"
	"github.com/leenzstra/timetable_server/common/utils"
)

type TimetableResponse struct {
	ID      uint       `json:"id"`
	GroupID uint       `json:"group_id"`
	Day     string     `json:"day"`
	WeekNum uint       `json:"week_num"`
	Table   []*Subject `json:"table"`
}

func NewTimetableResponse(g *models.Timetable) *TimetableResponse {
	bytes, _ := g.TableJson.MarshalJSON()
	m := make(map[string]string)
	json.Unmarshal(bytes, &m)

	subjects := make([]*Subject, 0)

	for time, v := range m {
		subject := &Subject{}
		if v != constant.EmptySubject {
			matches := constant.SubjectPattern.FindAllStringSubmatch(v, -1)
			subject.Time = time
			subject.SubjectName = strings.Trim(matches[0][1], " ")
			subject.SubjectType = strings.Trim(matches[0][2], " ()")
			subject.Teacher = strings.Trim(constant.TeacherNamePattern.FindAllStringSubmatch(matches[0][3], -1)[0][0], " ")
			subject.Location = strings.Trim(matches[0][4], " ")
		} else {
			subject.Time = time
		}
		subjects = append(subjects, subject)

	}
	return &TimetableResponse{
		ID:      g.ID,
		GroupID: g.GroupID,
		Day:     g.Day,
		WeekNum: g.WeekNum,
		Table:   subjects,
	}
}

// GetGroupTimetable godoc
// @Summary      Get group timetable
// @Description  Get group timetable by group name
// @ID get-group-timetable
// @Tags         timetable
// @Param group_name  path string true "Group Name"
// @Accept       json
// @Produce      json
// @Success      200  {object} models.ResponseBase{data=[]TimetableResponse}
// @Router       /timetable/timetables/{group_name} [get] 
func (h handler) GetGroupTimetable(c *fiber.Ctx) error {
	var g []*models.Timetable
	var u *models.Group

	groupName, err := url.QueryUnescape(c.Params("group_name"))
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	err = h.DB.Where(&models.Group{GroupName: groupName}).First(&u).Error
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	err = h.DB.Model(&models.Group{ID: u.ID}).Association("Timetables").Find(&g)
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	tResponses := make([]*TimetableResponse, 0)
	for _, t := range g {
		tr := NewTimetableResponse(t)
		tResponses = append(tResponses, tr)
	}

	return c.JSON(utils.WrapResponse(true, "", tResponses))
}
