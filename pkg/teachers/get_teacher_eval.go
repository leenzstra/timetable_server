package teachers

import (
	// "log"
	"net/url"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/leenzstra/timetable_server/common/models"
	"github.com/leenzstra/timetable_server/common/responses"
	"github.com/leenzstra/timetable_server/common/utils"
)

func NewTeacherEvalResponse(evals []*models.TeacherEvaluation, id int) *responses.TeacherEvalResponse {
	evaluations := []responses.Evaluation{}
	var markSum int = 0
	var avgMark float32 = 0.0

	for _, e := range evals {
		if e.Comment.Valid {
			evaluations = append(evaluations, responses.Evaluation{Comment: e.Comment.String, Mark: float32(e.Mark)})
		}
		markSum += e.Mark
	}

	if len(evals) == 0 {
		avgMark = 0
	} else {
		avgMark = float32(markSum) / float32(len(evals))
	}

	return &responses.TeacherEvalResponse{
		Id:       id,
		AverageMark:     avgMark,
		Evaluations: evaluations,
		Count:    len(evals),
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
// @Success      200  {object}  responses.ResponseBase{data=responses.TeacherEvalResponse}
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

	evals, err = h.DB.GetTeacherEvaluations(idInt)
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	evalResp := NewTeacherEvalResponse(evals, idInt)

	return c.JSON(utils.WrapResponse(true, "", evalResp))
}
