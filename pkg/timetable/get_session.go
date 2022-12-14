package timetable

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/leenzstra/timetable_server/common/constant"
	"github.com/leenzstra/timetable_server/common/models"
	"github.com/leenzstra/timetable_server/common/responses"
	"github.com/leenzstra/timetable_server/common/utils"
)


func NewSessionResponse(g *models.Session) *responses.SessionResponse {
	bytes, _ := g.Exams.MarshalJSON()
	m := make(map[string]string)
	json.Unmarshal(bytes, &m)

	subjects := make([]*responses.SessionSubject, 0)

	for date, v := range m {
		subject := &responses.SessionSubject{}
		if v != constant.EmptySubject {
			matches := constant.SubjectPattern.FindAllStringSubmatch(v, -1)
			subject.Date = date
			subject.SubjectName = strings.Trim(matches[0][1], " ")
			subject.SubjectType = strings.Trim(matches[0][2], " ()")
			subject.Teacher = strings.Trim(constant.TeacherNamePattern.FindAllStringSubmatch(matches[0][3], -1)[0][0], " ")
			subject.Location = strings.Trim(matches[0][4], " ")
		} else {
			subject.Date = date
		}
		subjects = append(subjects, subject)

	}
	return &responses.SessionResponse{
		ID:      g.ID,
		GroupID: g.GroupID,
		Table:   subjects,
	}
}

// GetGroupSession godoc
// @Summary      Get session timetable
// @Description  Get session timetable by group name
// @ID get-group-session
// @Param group_name  path string true "Group Name"
// @Tags         timetable
// @Accept       json
// @Produce      json
// @Success      200  {object} responses.ResponseBase{data=[]responses.SessionResponse}
// @Router       /timetable/sessions/{group_name} [get]
func (h handler) GetGroupSession(c *fiber.Ctx) error {
	groupName, err := url.QueryUnescape(c.Params("group_name"))
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	// group, err := h.DB.GetGroupByName(groupName)
	// if err != nil {
	// 	return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	// }

	session, err := h.DB.GetGroupSession(groupName)
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	response := NewSessionResponse(session)

	return c.JSON(utils.WrapResponse(true, "", response))
}
