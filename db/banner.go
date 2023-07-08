package db

type Banner struct {
	ID                             int64             `json:"id"`
	Term                           string            `json:"term"`
	TermDesc                       string            `json:"termDesc"`
	CourseReferenceNumber          string            `json:"courseReferenceNumber"`
	PartOfTerm                     string            `json:"partOfTerm"`
	CourseNumber                   string            `json:"courseNumber"`
	Subject                        string            `json:"subject"`
	SubjectDescription             string            `json:"subjectDescription"`
	SequenceNumber                 string            `json:"sequenceNumber"`
	CampusDescription              string            `json:"campusDescription"`
	ScheduleTypeDescription        string            `json:"scheduleTypeDescription"`
	CourseTitle                    string            `json:"courseTitle"`
	CreditHours                    float64           `json:"creditHours"`
	MaximumEnrollment              int64             `json:"maximumEnrollment"`
	Enrollment                     int64             `json:"enrollment"`
	SeatsAvailable                 int64             `json:"seatsAvailable"`
	WaitCapacity                   int64             `json:"waitCapacity"`
	WaitCount                      int64             `json:"waitCount"`
	WaitAvailable                  int64             `json:"waitAvailable"`
	CrossList                      interface{}       `json:"crossList"`
	CrossListCapacity              interface{}       `json:"crossListCapacity"`
	CrossListCount                 interface{}       `json:"crossListCount"`
	CrossListAvailable             interface{}       `json:"crossListAvailable"`
	CreditHourHigh                 interface{}       `json:"creditHourHigh"`
	CreditHourLow                  float64           `json:"creditHourLow"`
	CreditHourIndicator            interface{}       `json:"creditHourIndicator"`
	OpenSection                    bool              `json:"openSection"`
	LinkIdentifier                 interface{}       `json:"linkIdentifier"`
	IsSectionLinked                bool              `json:"isSectionLinked"`
	SubjectCourse                  string            `json:"subjectCourse"`
	Faculty                        []Faculty         `json:"faculty"`
	MeetingsFaculty                []MeetingsFaculty `json:"meetingsFaculty"`
	ReservedSeatSummary            interface{}       `json:"reservedSeatSummary"`
	SectionAttributes              interface{}       `json:"sectionAttributes"`
	InstructionalMethod            string            `json:"instructionalMethod"`
	InstructionalMethodDescription string            `json:"instructionalMethodDescription"`
}

type Faculty struct {
	BannerID              string      `json:"bannerId"`
	Category              interface{} `json:"category"`
	Class                 string      `json:"class"`
	CourseReferenceNumber string      `json:"courseReferenceNumber"`
	DisplayName           string      `json:"displayName"`
	EmailAddress          string      `json:"emailAddress"`
	PrimaryIndicator      bool        `json:"primaryIndicator"`
	Term                  string      `json:"term"`
}

type MeetingsFaculty struct {
	Category              string        `json:"category"`
	Class                 string        `json:"class"`
	CourseReferenceNumber string        `json:"courseReferenceNumber"`
	Faculty               []interface{} `json:"faculty"`
	MeetingTime           MeetingTime   `json:"meetingTime"`
	Term                  string        `json:"term"`
}

type MeetingTime struct {
	BeginTime              interface{} `json:"beginTime"`
	Building               interface{} `json:"building"`
	BuildingDescription    interface{} `json:"buildingDescription"`
	Campus                 interface{} `json:"campus"`
	CampusDescription      interface{} `json:"campusDescription"`
	Category               string      `json:"category"`
	Class                  string      `json:"class"`
	CourseReferenceNumber  string      `json:"courseReferenceNumber"`
	CreditHourSession      float64     `json:"creditHourSession"`
	EndDate                string      `json:"endDate"`
	EndTime                interface{} `json:"endTime"`
	Friday                 bool        `json:"friday"`
	HoursWeek              int64       `json:"hoursWeek"`
	MeetingScheduleType    string      `json:"meetingScheduleType"`
	MeetingType            string      `json:"meetingType"`
	MeetingTypeDescription string      `json:"meetingTypeDescription"`
	Monday                 bool        `json:"monday"`
	Room                   interface{} `json:"room"`
	Saturday               bool        `json:"saturday"`
	StartDate              string      `json:"startDate"`
	Sunday                 bool        `json:"sunday"`
	Term                   string      `json:"term"`
	Thursday               bool        `json:"thursday"`
	Tuesday                bool        `json:"tuesday"`
	Wednesday              bool        `json:"wednesday"`
}
