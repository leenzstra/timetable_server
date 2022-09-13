package constant

import "regexp"

var EmptySubject = "None"

var SubjectPattern = regexp.MustCompile(`(.*?)(\(.+\))([^\d]+)(.+)`)
var TeacherNamePattern = regexp.MustCompile(`[^\. ]+ .\.[\s]*.\.`)