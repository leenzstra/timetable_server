package timetable

import (
	"encoding/json"
	"log"
	"net/url"
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/leenzstra/timetable_server/common/constant"
	"github.com/leenzstra/timetable_server/common/models"
	"github.com/leenzstra/timetable_server/common/utils"
)

type TimetableResponse struct {
	ID      uint            `json:"id"`
	GroupID uint            `json:"group_id"`
	Day     uint          `json:"day"`
	WeekNum uint            `json:"week_num"`
	Table   []*GroupSubject `json:"table"`
}

type GroupSubject struct {
	Time        uint   `json:"time"`
	SubjectName string `json:"subject_name"`
	SubjectType string `json:"subject_type"`
	Teacher     string `json:"teacher"`
	Location    string `json:"location"`
}

func NewTimetableResponse(g *models.Timetable) *TimetableResponse {
	bytes, _ := g.TableJson.MarshalJSON()
	m := make(map[string]string)
	json.Unmarshal(bytes, &m)

	subjects := make([]*GroupSubject, 0)

	for time, v := range m {
		subject := &GroupSubject{}
		timeUnix, err := utils.StringTimeToUnix(time)
		if err != nil {
			log.Print(err)
		}
		if v != constant.EmptySubject {
			matches := constant.SubjectPattern.FindAllStringSubmatch(v, -1)
			subject.Time = timeUnix
			subject.SubjectName = strings.Trim(matches[0][1], " ")
			subject.SubjectType = strings.Trim(matches[0][2], " ()")
			subject.Teacher = strings.Trim(constant.TeacherNamePattern.FindAllStringSubmatch(matches[0][3], -1)[0][0], " ")
			subject.Location = strings.Trim(matches[0][4], " ")
		} else {
			subject.Time = timeUnix
		}
		subjects = append(subjects, subject)

	}

	sort.Slice(subjects, func(i, j int) bool {
		return subjects[i].Time < subjects[j].Time
	})

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
// @Description  Get group timetable by group name and type
// @ID get-group-timetable
// @Tags         timetable
// @Param group_name  path string true "Group Name"
// @Param timetable_type  path string true "timetable type"
// @Accept       json
// @Produce      json
// @Success      200  {object} models.ResponseBase{data=[]TimetableResponse}
// @Router       /timetable/timetables/{group_name}/{timetable_type} [get]
func (h handler) GetGroupTimetable(c *fiber.Ctx) error {
	var g []*models.Timetable
	var u *models.Group

	groupName, err := url.QueryUnescape(c.Params("group_name"))
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	tableType, err := url.QueryUnescape(c.Params("timetable_type"))
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	err = h.DB.Where(&models.Group{GroupName: groupName, TableKind: tableType}).First(&u).Error
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

	sort.Slice(tResponses, func(i, j int) bool {
		if tResponses[i].Day != tResponses[j].Day {
			return tResponses[i].Day < tResponses[j].Day
		}
		return tResponses[i].WeekNum < tResponses[j].WeekNum
	})

	return c.JSON(utils.WrapResponse(true, "", tResponses))
}
