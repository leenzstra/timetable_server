package teachers

import (
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/leenzstra/timetable_server/common/models"
	"github.com/leenzstra/timetable_server/common/utils"
)

const MaxMark = 100
const MinMark = 0

type TeacherMarkBody struct {
	SID     int    `json:"sid"`
	Mark    int    `json:"mark"`
	Comment string `json:"comment"`
}

// SetTeacherMark godoc
// @Summary      Set teachers mark
// @Description  Set teachers mark
// @ID set-teacher-mark
// @Tags         teachers
// @Param payload  body TeacherMarkBody true "Mark payload"
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.ResponseBase{data=TeacherMarkBody}
// @Router       /teachers/set_mark/ [post]
func (h handler) SetMark(c *fiber.Ctx) error {

	payload := TeacherMarkBody{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if payload.Mark < MinMark && payload.Mark > MaxMark {
		return errors.New("mark < minMark or mark > maxMark")
	}

	teacher := models.Teacher{}
	err := h.DB.Where("id = ?", payload.SID).Find(&teacher).Error
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	if err := h.DB.Create(&models.TeacherEvaluation{TeacherId: int(teacher.Id), Mark: payload.Mark, Comment: sql.NullString{String: payload.Comment, Valid: true}}).Error; err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	return c.JSON(utils.WrapResponse(true, "", payload))
}
