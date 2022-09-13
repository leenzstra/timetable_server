package utils

import "github.com/leenzstra/timetable_server/common/models"

func WrapResponse(result bool, message string, data interface{}) *models.ResponseBase {
	return &models.ResponseBase{
		Result: result,
		Message: message,
		Data: data,
	}
}