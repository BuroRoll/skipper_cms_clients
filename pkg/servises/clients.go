package service

import (
	"Skipper_cms_clients/pkg/models"
	"Skipper_cms_clients/pkg/repository"
	"math"
)

type ClientsService struct {
	repo repository.Clients
}

func NewClientsService(repo repository.Clients) *ClientsService {
	return &ClientsService{repo: repo}
}

func (c ClientsService) GetClients(pagination *models.Pagination) ([]models.User, error) {
	clients, err := c.repo.GetClients(&pagination)
	for i := range clients {
		if clients[i].IsMentor {
			stat := c.GetMentorStatistic(clients[i].ID)
			clients[i].UpdateStatistic(&stat)
		} else {
			stat := c.GetMentiStatistic(clients[i].ID)
			clients[i].UpdateStatistic(&stat)
		}
	}
	return clients, err
}

func (c ClientsService) GetClient(userId uint) (models.User, error) {
	client, err := c.repo.GetClient(userId)
	if client.IsMentor {
		stat := c.GetMentorStatistic(client.ID)
		client.UpdateStatistic(&stat)
	} else {
		stat := c.GetMentiStatistic(client.ID)
		client.UpdateStatistic(&stat)
	}
	return client, err
}

func (u ClientsService) GetMentiStatistic(userId uint) models.Statistic {
	countLessons := u.repo.GetMentiCountLessons(userId, "full", true)

	countCompletedLessons := u.repo.GetMentiCountLessons(userId, "", true)
	countLastMonthCompletedLessons := u.repo.GetMentiCountLessons(userId, "last_month", true)
	countLastThreeMonthCompletedLessons := u.repo.GetMentiCountLessons(userId, "last_three_month", true)

	countLastMonthUnclompletedLessons := u.repo.GetMentiCountLessons(userId, "last_month", false)
	countLastThreeMonthUnclompletedLessons := u.repo.GetMentiCountLessons(userId, "last_three_month", false)
	countUncomplitedLessons := u.repo.GetMentiCountLessons(userId, "", false)

	var fullAttendance = 0.0
	var lastMonthAttendance = 0.0
	var lastThreeMonthAttendance = 0.0
	if countLessons != 0 {
		fullAttendance = math.Round(countCompletedLessons / countLessons * 100)
	}
	if sum := countLastMonthCompletedLessons + countLastMonthUnclompletedLessons; sum != 0 {
		lastMonthAttendance = math.Round(countLastMonthCompletedLessons / (sum) * 100)
	}
	if sum := countLastThreeMonthCompletedLessons + countLastThreeMonthUnclompletedLessons; sum != 0 {
		lastThreeMonthAttendance = math.Round(countLastThreeMonthCompletedLessons / (sum) * 100)
	}
	stat := models.Statistic{
		LessonsCount:                      countCompletedLessons,
		StudentsCount:                     0,
		LastMonthLessonsCount:             countLastMonthCompletedLessons,
		LastThreeMonthsLessonsCount:       countLastThreeMonthCompletedLessons,
		LastMonthUnclompletedLessons:      countLastMonthUnclompletedLessons,
		LastThreeMonthUnclompletedLessons: countLastThreeMonthUnclompletedLessons,
		UncomplitedLessons:                countUncomplitedLessons,
		FullAttendance:                    fullAttendance,
		LastMonthAttendance:               lastMonthAttendance,
		LastThreeMonthAttendance:          lastThreeMonthAttendance,
	}
	return stat
}

func (u ClientsService) GetMentorStatistic(userId uint) models.Statistic {
	countCompletedStudents := u.repo.GetMentorCountStudents(userId)
	countLessons := u.repo.GetMentorCountLessons(userId, "full", true)

	countCompletedLessons := u.repo.GetMentorCountLessons(userId, "", true)
	countLastMonthCompletedLessons := u.repo.GetMentorCountLessons(userId, "last_month", true)
	countLastThreeMonthCompletedLessons := u.repo.GetMentorCountLessons(userId, "last_three_month", true)

	countLastMonthUnclompletedLessons := u.repo.GetMentorCountLessons(userId, "last_month", false)
	countLastThreeMonthUnclompletedLessons := u.repo.GetMentorCountLessons(userId, "last_three_month", false)
	countUncomplitedLessons := u.repo.GetMentorCountLessons(userId, "", false)

	var fullAttendance = 0.0
	var lastMonthAttendance = 0.0
	var lastThreeMonthAttendance = 0.0
	if countLessons != 0 {
		fullAttendance = math.Round(countCompletedLessons / countLessons * 100)
	}
	if sum := countLastMonthCompletedLessons + countLastMonthUnclompletedLessons; sum != 0 {
		lastMonthAttendance = math.Round(countLastMonthCompletedLessons / (sum) * 100)
	}
	if sum := countLastThreeMonthCompletedLessons + countLastThreeMonthUnclompletedLessons; sum != 0 {
		lastThreeMonthAttendance = math.Round(countLastThreeMonthCompletedLessons / (sum) * 100)
	}
	stat := models.Statistic{
		LessonsCount:                      countCompletedLessons,
		StudentsCount:                     countCompletedStudents,
		LastMonthLessonsCount:             countLastMonthCompletedLessons,
		LastThreeMonthsLessonsCount:       countLastThreeMonthCompletedLessons,
		LastMonthUnclompletedLessons:      countLastMonthUnclompletedLessons,
		LastThreeMonthUnclompletedLessons: countLastThreeMonthUnclompletedLessons,
		UncomplitedLessons:                countUncomplitedLessons,
		FullAttendance:                    fullAttendance,
		LastMonthAttendance:               lastMonthAttendance,
		LastThreeMonthAttendance:          lastThreeMonthAttendance,
	}
	return stat
}
