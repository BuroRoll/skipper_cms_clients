package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//Base information
	Email          string
	Phone          string `gorm:"index:unique"`
	FirstName      string
	SecondName     string
	Patronymic     string
	DateOfBirthday string
	Description    string
	Time           string
	ProfilePicture string          `gorm:"default:'default_profile_picture.jpeg'"`
	IsMentor       bool            `gorm:"default:false"`
	IsVerifyEmail  bool            `gorm:"default:false"`
	IsVerifyPhone  bool            `gorm:"default:false"`
	Communications []Communication `gorm:"ForeignKey:ParentId"`
	Rating         float64         `gorm:"default:0"`
	// Жалобы пользователя
	UserReports []Report `gorm:"ForeignKey:FromUserId"`
	// Жалобы на пользователя
	Reports []Report `gorm:"ForeignKey:ToUserId"`

	//Mentor information
	Specialization    string
	Educations        []Education        `gorm:"ForeignKey:ParentId"`
	WorkExperiences   []WorkExperience   `gorm:"ForeignKey:ParentId"`
	OtherInformation  []OtherInformation `gorm:"ForeignKey:ParentId"`
	AverageClassPrice uint               `gorm:"default:0"`

	Statistic Statistic `gorm:"-"`
}
type Communication struct {
	gorm.Model
	ParentId  uint
	Messenger []*Messenger `gorm:"many2many:messenger_communication;"`
	Login     string
}

type Messenger struct {
	gorm.Model
	Name           string
	Communications []*Communication `gorm:"many2many:messenger_communication;"`
}

type Education struct {
	gorm.Model
	ParentId    uint
	Institution string
	StartYear   int
	EndYear     int
	Degree      string
}

type WorkExperience struct {
	gorm.Model
	ParentId     uint
	Organization string
	StartYear    string
	EndYear      string
}

type OtherInformation struct {
	gorm.Model
	ParentId uint
	Data     string
}

type Statistic struct {
	LessonsCount                      float64 `json:"lessons_count"`
	StudentsCount                     uint    `json:"students_count"`
	LastMonthLessonsCount             float64 `json:"last_month_lessons_count"`
	LastThreeMonthsLessonsCount       float64 `json:"last_three_months_lessons_count"`
	UncomplitedLessons                float64 `json:"uncomplited_lessons"`
	LastMonthUnclompletedLessons      float64 `json:"last_month_unclompleted_lessons"`
	LastThreeMonthUnclompletedLessons float64 `json:"last_three_month_unclompleted_lessons"`
	FullAttendance                    float64 `json:"full_attendance"`
	LastMonthAttendance               float64 `json:"last_month_attendance"`
	LastThreeMonthAttendance          float64 `json:"last_three_month_attendance"`
}

type Report struct {
	gorm.Model
	FromUserId  uint
	ToUserId    uint
	ReportTheme string
	ReportText  string
}

type Pagination struct {
	Limit  int      `json:"limit"`
	Page   int      `json:"page"`
	Sort   string   `json:"sort"`
	Search []string `json:"search"`
}

func (user *User) UpdateStatistic(stat *Statistic) {
	user.Statistic = *stat
}
