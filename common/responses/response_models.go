package responses

type ResponseBase struct {
	Result  bool        `json:"result"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type TimetableResponse struct {
	ID      uint            `json:"id"`
	GroupID uint            `json:"group_id"`
	Day     uint          `json:"day"`
	WeekNum uint            `json:"week_num"`
	Table   []*GroupSubject `json:"table"`
}

type GroupSubject struct {
	Time        uint   `json:"time"`
	SubjectName string `json:"subject_name"`
	SubjectType string `json:"subject_type"`
	Teacher     string `json:"teacher"`
	Location    string `json:"location"`
}

type TeacherEvalResponse struct {
	Id          int          `json:"id"`
	AverageMark float32      `json:"average_mark"`
	Count       int          `json:"count"`
	Evaluations []Evaluation `json:"evaluations"`
}

type Evaluation struct {
	Comment string  `json:"comment"`
	Mark    float32 `json:"mark"`
}

type TeachersResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Position   string `json:"position"`
}

type GroupResponse struct {
	ID        uint   `json:"id"`
	Faculty   string `json:"faculty"`
	Direction string `json:"direction"`
	GroupName string `json:"group_name"`
}

type SessionResponse struct {
	ID       uint       `json:"id"`
	GroupID  uint       `json:"group_id"`
	Addition string     `json:"addition"`
	Table    []*SessionSubject `json:"table"`
}

type SessionSubject struct {
	Date        string `json:"date"`
	SubjectName string `json:"subject_name"`
	SubjectType string `json:"subject_type"`
	Teacher     string `json:"teacher"`
	Location    string `json:"location"`
}

type TeacherMarkBody struct {
	SID     int    `json:"sid"`
	Mark    int    `json:"mark"`
	Comment string `json:"comment"`
}

