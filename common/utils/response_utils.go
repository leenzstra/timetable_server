package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/leenzstra/timetable_server/common/constant"
	"github.com/leenzstra/timetable_server/common/models"
	"github.com/leenzstra/timetable_server/common/responses"
)

func WrapResponse(result bool, message string, data interface{}) *responses.ResponseBase {
	return &responses.ResponseBase{
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

func FetchSubjectInfo(rawSubject string) (*responses.GroupSubject, error) {
	subject := &responses.GroupSubject{}
	if rawSubject != constant.EmptySubject {
		matches := constant.SubjectPattern.FindAllStringSubmatch(rawSubject, -1)

		if len(matches) == 0 {
			return subject, errors.New("no subject info found")
		}
		if len(matches[0]) < 4 {
			return subject, errors.New("no exact subject info found")
		}

		subject.SubjectName = strings.Trim(matches[0][1], " ")
		subject.SubjectType = strings.Trim(matches[0][2], " ()")
		subject.Teacher = strings.Trim(constant.TeacherNamePattern.FindAllStringSubmatch(matches[0][3], -1)[0][0], " ")
		subject.Location = strings.Trim(matches[0][4], " ")
	}
	return subject, nil
}

 func FetchGroupSubjectsFromTimetable(timetable models.Timetable) ([]*responses.GroupSubject, error) {
	bytes, err := timetable.TableJson.MarshalJSON()
	if err != nil {
		return nil, err
	}

	jsonData := make(map[string]string)
	err = json.Unmarshal(bytes, &jsonData)
	if err != nil {
		return nil, err
	}

	subjects := make([]*responses.GroupSubject, 0)

	for time, rawSubject := range jsonData {
		timeUnix, _ := StringTimeToUnix(time)
		subject, _ := FetchSubjectInfo(rawSubject)
		subject.Time = timeUnix

		subjects = append(subjects, subject)
	}

	return subjects, nil
 }

 func StringBetween(str, left, right, def string) string {
	start := strings.Index(str, left)
	end := strings.Index(str, right)
	if start != -1 && end != -1 {
		return str[start+1:end]
	} else {
		return def
	}
 }
