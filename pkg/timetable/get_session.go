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

type SessionResponse struct {
	ID       uint       `json:"id"`
	GroupID  uint       `json:"group_id"`
	Addition string     `json:"addition"`
	Table    []*Subject `json:"table"`
}

type Subject struct {
	Time        string `json:"time"`
	SubjectName string `json:"subject_name"`
	SubjectType string `json:"subject_type"`
	Teacher     string `json:"teacher"`
	Location    string `json:"location"`
}

func NewSessionResponse(g *models.Session) *SessionResponse {
	bytes, _ := g.Exams.MarshalJSON()
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
	return &SessionResponse{
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
// @Success      200  {object} models.ResponseBase{data=[]SessionResponse}
// @Router       /timetable/sessions/{group_name} [get]
func (h handler) GetGroupSession(c *fiber.Ctx) error {
	var g []*models.Session
	var u *models.Group

	groupName, err := url.QueryUnescape(c.Params("group_name"))
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	err = h.DB.Where(&models.Group{GroupName: groupName}).First(&u).Error
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	err = h.DB.Model(&models.Group{ID: u.ID}).Association("Sessions").Find(&g)
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	sResponses := make([]*SessionResponse, 0)
	for _, s := range g {
		sr := NewSessionResponse(s)
		sResponses = append(sResponses, sr)
	}

	return c.JSON(utils.WrapResponse(true, "", sResponses))
}
