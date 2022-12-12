package teachers

import (
	// "log"
	"net/url"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/leenzstra/timetable_server/common/models"
	"github.com/leenzstra/timetable_server/common/utils"
)

type TeacherEvalResponse struct {
	Id       int      `json:"id"`
	Mark     float32  `json:"mark"`
	Count    int      `json:"count"`
	Comments []string `json:"comments"`
}

func NewTeacherEvalResponse(evals []*models.TeacherEvaluation, id int) *TeacherEvalResponse {
	comments := []string{}
	var markSum int = 0
	var mark float32 = 0.0

	for _, e := range evals {
		if e.Comment.Valid {
			comments = append(comments, e.Comment.String)
		}
		markSum += e.Mark
		// log.Println(e.Mark)
	}

	if len(evals) == 0 {
		mark = 0
	} else {
		mark = float32(markSum) / float32(len(evals))
	}

	return &TeacherEvalResponse{
		Id:       id,
		Mark:     mark,
		Comments: comments,
		Count: len(evals),
	}
}

// GetTeacherEval godoc
// @Summary      Get teacher evaluation
// @Description  Get teacher evaluation
// @ID get-teacher-eval
// @Tags         teachers
// @Param id  path int true "Teacher ID"
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.ResponseBase{data=TeacherEvalResponse}
// @Router       /teachers/eval/{id} [get]
func (h handler) GetTeacherEval(c *fiber.Ctx) error {
	var evals []*models.TeacherEvaluation

	id, err := url.QueryUnescape(c.Params("id", ""))
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	if id == "undefined" {
		id = ""
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	err = h.DB.Joins("JOIN teachers ON teachers.id = teacher_evaluations.teacher_id").Where("teachers.id = ?", idInt).Find(&evals).Error
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	evalResp := NewTeacherEvalResponse(evals, idInt)
	// log.Println(evalResp)

	return c.JSON(utils.WrapResponse(true, "", evalResp))
}
