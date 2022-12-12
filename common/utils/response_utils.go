package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/leenzstra/timetable_server/common/models"
)

func WrapResponse(result bool, message string, data interface{}) *models.ResponseBase {
	return &models.ResponseBase{
		Result:  result,
		Message: message,
		Data:    data,
	}
}

func StringTimeToUnix(timeString string) (uint, error) {
	hm := strings.Split(timeString, ":")
	if len(hm) != 2 {
		return 0, fmt.Errorf("wrong time format: %s", timeString)
	}

	hours, err := strconv.Atoi(hm[0])
	if err != nil {
		return 0, fmt.Errorf("wrong hours format: %s", hm[0])
	}

	mins, err := strconv.Atoi(hm[1])
	if err != nil {
		return 0, fmt.Errorf("wrong minutes format: %s", hm[1])
	}

	unixTime := hours*int(time.Hour.Seconds()) + mins*int(time.Minute.Seconds())

	return uint(unixTime), nil

}
