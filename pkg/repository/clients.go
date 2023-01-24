package repository

import (
	"Skipper_cms_clients/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ClientsPostgres struct {
	db *gorm.DB
}

func NewClientsPostgres(db *gorm.DB) *ClientsPostgres {
	return &ClientsPostgres{db: db}
}

func (c ClientsPostgres) GetClients(pagination **models.Pagination) ([]models.User, error) {
	var clients []models.User
	offset := ((*pagination).Page - 1) * (*pagination).Limit
	queryBuider := c.db.Limit((*pagination).Limit).Offset(offset)

	var result *gorm.DB
	if (len((*pagination).Search)) > 0 {
		result = queryBuider.
			Preload(clause.Associations).
			Where("phone IN (?) OR first_name IN (?) OR second_name IN (?) OR patronymic IN (?)", (*pagination).Search, (*pagination).Search, (*pagination).Search, (*pagination).Search).
			Find(&clients)
	} else {
		result = queryBuider.
			Preload(clause.Associations).
			Find(&clients)
	}
	//c.db.Preload(clause.Associations).Find(&clients)
	if result.Error != nil {
		return nil, result.Error
	}

	return clients, nil
}

func (c ClientsPostgres) GetClient(userId uint) (models.User, error) {
	var client models.User
	c.db.Preload(clause.Associations).Find(&client, userId)
	return client, nil
}

type completedLessons struct {
	UserId       uint    `json:"user_id"`
	CountLessons float64 `json:"count_lessons"`
}

type completedStudents struct {
	UserId                 uint   `json:"user_id"`
	Status                 string `json:"status"`
	CountCompletedStudents uint   `json:"count_completed_students"`
}

func (u ClientsPostgres) GetMentorCountStudents(userId uint) uint {
	var stat completedStudents
	u.db.Select("user_id, status, COUNT(DISTINCT menti_id) AS count_completed_students").
		Table("user_classes").
		Where("user_id = ? and status = ?", userId, "completed").
		Group("user_id, status").
		Find(&stat)
	return stat.CountCompletedStudents
}

func (u ClientsPostgres) GetMentorCountLessons(userId uint, time string, isComplete bool) float64 {
	var stat completedLessons
	queryBuider := u.db.Select("count(bt.is_end) as count_lessons, count(is_success) as isok, count(bt.booking_class_id) as bcid, count(uc.id) as ids, uc.user_id, count(bt.time) as ctime").
		Table("booking_times bt").
		Joins("inner join user_classes uc on uc.id = bt.booking_class_id").
		Group("uc.user_id")
	switch time {
	case "full":
		queryBuider.Where("user_id = ? and is_end = true", userId)
	case "last_month":
		queryBuider.Where("user_id = ? and is_end = true and is_success = ? and TO_DATE(left(bt.time, -2),'YYYY/MM/DD') > CURRENT_DATE - INTERVAL '30 days'", userId, isComplete)
	case "last_three_month":
		queryBuider.Where("user_id = ? and is_end = true and is_success = ? and TO_DATE(left(bt.time, -2),'YYYY/MM/DD') > CURRENT_DATE - INTERVAL '90 days'", userId, isComplete)
	default:
		queryBuider.Where("user_id = ? and is_end = true and is_success = ?", userId, isComplete)
	}
	queryBuider.Find(&stat)
	return stat.CountLessons
}

func (u ClientsPostgres) GetMentiCountLessons(userId uint, time string, isComplete bool) float64 {
	var stat completedLessons
	queryBuider := u.db.Select("count(bt.is_end) as count_lessons, count(is_success) as isok, count(bt.booking_class_id) as bcid, count(uc.id) as ids, uc.menti_id as user_id, count(bt.time) as ctime").
		Table("booking_times bt").
		Joins("inner join user_classes uc on uc.id = bt.booking_class_id").
		Group("uc.menti_id")
	switch time {
	case "full":
		queryBuider.Where("uc.menti_id = ? and is_end = true", userId)
	case "last_month":
		queryBuider.Where("uc.menti_id = ? and is_end = true and is_success = ? and TO_DATE(left(bt.time, -2),'YYYY/MM/DD') > CURRENT_DATE - INTERVAL '30 days'", userId, isComplete)
	case "last_three_month":
		queryBuider.Where("uc.menti_id = ? and is_end = true and is_success = ? and TO_DATE(left(bt.time, -2),'YYYY/MM/DD') > CURRENT_DATE - INTERVAL '90 days'", userId, isComplete)
	default:
		queryBuider.Where("uc.menti_id = ? and is_end = true and is_success = ?", userId, isComplete)
	}
	queryBuider.Find(&stat)
	return stat.CountLessons
}
