package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/totoro081295/daily-report-api/category"
	categoryRepo "github.com/totoro081295/daily-report-api/category/repository"
	"github.com/totoro081295/daily-report-api/project"
	projectRepo "github.com/totoro081295/daily-report-api/project/repository"
	"github.com/totoro081295/daily-report-api/task"
	model "github.com/totoro081295/daily-report-api/task"
	"github.com/totoro081295/daily-report-api/task/repository"
	"github.com/totoro081295/daily-report-api/taskdate"
	taskDateRepo "github.com/totoro081295/daily-report-api/taskdate/repository"
)

type taskUsecase struct {
	taskRepo     repository.TaskRepository
	taskDateRepo taskDateRepo.TaskDateRepository
	projectRepo  projectRepo.ProjectRepository
	categoryRepo categoryRepo.CategoryRepository
}

// NewTaskUsecase mount task usecase
func NewTaskUsecase(
	task repository.TaskRepository,
	taskDate taskDateRepo.TaskDateRepository,
	project projectRepo.ProjectRepository,
	category categoryRepo.CategoryRepository,
) TaskUsecase {
	return &taskUsecase{
		taskRepo:     task,
		taskDateRepo: taskDate,
		projectRepo:  project,
		categoryRepo: category,
	}
}

// TaskUsecase usecase interface
type TaskUsecase interface {
	List(accountID uuid.UUID) ([]*model.Response, error)
}

type taskData struct {
	tasks      task.Collection
	taskDates  []*taskdate.TaskDate
	categories []*category.Category
	projects   []*project.Project
}
