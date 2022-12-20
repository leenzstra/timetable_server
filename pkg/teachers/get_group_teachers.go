package teachers

import (
	"log"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/leenzstra/timetable_server/common/models"
	"github.com/leenzstra/timetable_server/common/responses"
	"github.com/leenzstra/timetable_server/common/utils"
)


func fetchTeachersNames(subjects []*responses.GroupSubject) []string {
	names := make(map[string]bool)
	for _, subj := range subjects {
		trimmed := strings.Trim(subj.Teacher, " ")
		if trimmed != "" {
			names[trimmed] = true
		}
	}

	keys := make([]string, len(names))
	for k := range names {
        keys = append(keys, k)
    }
	return keys
}

// TODO: можно и получше сделать
func selectTeachersFromGroup(teachers []*models.Teacher, names []string) []*models.Teacher {
	results := make([]*models.Teacher, 0)
	for _, teacher := range teachers {
		for _, name := range names {
			nameTrimmed := strings.Trim(strings.ReplaceAll(strings.ToLower(name), ".", " "), " ")
			nameSplitted := strings.Split(nameTrimmed, " ")
			targetNameTrimmed := strings.Trim(strings.ToLower(teacher.FIO), " ")
			targetNameSplitted := strings.Split(targetNameTrimmed, " ")

			if len(nameSplitted) != len(targetNameSplitted) || len(nameSplitted) == 0 {
				continue
			}

			// фамилия
			namesEquals := nameSplitted[0] == targetNameSplitted[0]

			//остальное
			for i := 1; i < len(nameSplitted); i++ {
				namesEquals = namesEquals && ([]rune(nameSplitted[i])[0] == ([]rune(targetNameSplitted[i])[0]))
			}


			if namesEquals {
				log.Println(nameSplitted, targetNameSplitted)
				results = append(results, teacher)
			}

		}
	}
	return results
}

func selectType(groups []*models.Group) string {
	for i := len(groups)-1; i >= 0; i-- {
		if !strings.Contains(groups[i].TableKind, "сессия") {
			return groups[i].TableKind;
		}
	}
	return ""
}

// GetGroupTeachers godoc
// @Summary      Get group teachers list
// @Description  Get group teachers list
// @ID get-group-teachers
// @Tags         teachers
// @Param group_name  path string true "Name group_name"
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.ResponseBase{data=[]responses.TeachersResponse}
// @Router       /teachers/group/{group_name} [get]
func (h handler) GetGroupTeachers(c *fiber.Ctx) error {

	groupName, err := url.QueryUnescape(c.Params("group_name", ""))
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	if groupName == "undefined" {
		return c.JSON(utils.WrapResponse(false, "No group name", nil))
	}

	groupTimetableTypes, err := h.DB.GetGroupTimetableTypes(groupName)
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	if len(groupTimetableTypes) == 0 {
		return c.JSON(utils.WrapResponse(false, "Timetable types not found", nil))
	}

	tableType := selectType(groupTimetableTypes);

	groupTimetable, err := h.DB.GetGroupTimetable(groupName, tableType)
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	subjects := make([]*responses.GroupSubject, 0)
	for _, timetable := range groupTimetable {
		s, err := utils.FetchGroupSubjectsFromTimetable(*timetable)
		if err != nil {
			return c.JSON(utils.WrapResponse(false, err.Error(), nil))
		}
		subjects = append(subjects, s...)
	}

	teachers, err := h.DB.GetTeachers("")
	if err != nil {
		return c.JSON(utils.WrapResponse(false, err.Error(), nil))
	}

	names := fetchTeachersNames(subjects)
	
	groupTeachers := selectTeachersFromGroup(teachers, names)
	response := make([]*responses.TeachersResponse, 0)
	for _, gr := range groupTeachers {
		tr := NewTeachersResponse(gr)
		response = append(response, tr)
	}

	return c.JSON(utils.WrapResponse(true, "", response))
}

