package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/totoro081295/daily-report-api/category"
	"github.com/totoro081295/daily-report-api/project"
	model "github.com/totoro081295/daily-report-api/task"
)

func (t *taskUsecase) List(accountID uuid.UUID) ([]*model.Response, error) {
	data := taskData{}
	var err error
	data.tasks, err = t.taskRepo.ListByAccountID(accountID)
	if err != nil {
		return nil, err
	}
	taskIDs := data.tasks.GetIDs()
	data.taskDates, err = t.taskDateRepo.ListByTaskIDs(taskIDs)
	if err != nil {
		return nil, err
	}
	categoryIDs := data.tasks.GetCategoryIDs()
	data.categories, err = t.categoryRepo.ListByIDs(categoryIDs)
	if err != nil {
		return nil, err
	}
	projectIDs := data.tasks.GetProjectIDs()
	data.projects, err = t.projectRepo.ListByIDs(projectIDs)
	if err != nil {
		return nil, err
	}

	res := []*model.Response{}
	for _, date := range data.taskDates {
		var taskDateRes = model.Response{
			ID:         date.TaskID,
			TargetDate: date.TargetDate,
		}
		res = append(res, &taskDateRes)
	}

	for _, r := range res {
		for _, task := range data.tasks {
			if r.ID == task.ID {
				r.Name = task.Name
				for _, c := range data.categories {
					if task.CategoryID == c.ID {
						r.Category = category.Response{
							ID:          c.ID,
							Name:        c.Name,
							Color:       c.Color,
							Description: c.Description,
						}
					}
				}
				for _, p := range data.projects {
					if task.ProjectID == p.ID {
						r.Project = project.Response{
							ID:          p.ID,
							Name:        p.Name,
							Description: p.Description,
						}
					}
				}
			}
		}
	}
	return res, nil
}
