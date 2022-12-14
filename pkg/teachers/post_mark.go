package teachers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/leenzstra/timetable_server/common/responses"
	"github.com/leenzstra/timetable_server/common/utils"
)

const MaxMark = 5
const MinMark = 0

// SetTeacherMark godoc
// @Summary      Set teachers mark
// @Description  Set teachers mark
// @ID set-teacher-mark
// @Tags         teachers
// @Param payload  body responses.TeacherMarkBody true "Mark payload"
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.ResponseBase{data=string}
// @Router       /teachers/set_mark/ [post]
func (h handler) SetMark(c *fiber.Ctx) error {

	payload := responses.TeacherMarkBody{}
	if err := c.BodyParser(&payload); err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	if payload.Mark < MinMark && payload.Mark > MaxMark {
		return errors.New("mark < minMark or mark > maxMark")
	}

	teacher, err := h.DB.GetTeacherById(payload.SID)
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	if err := h.DB.CreateTeacherEvaluation(teacher.Id, payload.Mark, payload.Comment); err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	return c.JSON(utils.WrapResponse(true, "", ""))
}
