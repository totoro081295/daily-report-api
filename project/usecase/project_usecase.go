package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/totoro081295/daily-report-api/project"
	"github.com/totoro081295/daily-report-api/project/repository"
)

type projectUsecase struct {
	projectRepo repository.ProjectRepository
}

// NewProjectUsecase mount project usecase
func NewProjectUsecase(project repository.ProjectRepository) ProjectUsecase {
	return &projectUsecase{
		projectRepo: project,
	}
}

// ProjectUsecase usecase interface
type ProjectUsecase interface {
	Get(id uuid.UUID) (*project.Response, error)
	List() ([]*project.Response, error)
	Create(payload project.Payload) (*project.Response, error)
}

func (u *projectUsecase) Get(id uuid.UUID) (*project.Response, error) {
	p, err := u.projectRepo.Get(id)
	if err != nil {
		return nil, err
	}
	res := format(p)
	return &res, nil
}
func (u *projectUsecase) List() ([]*project.Response, error) {
	projects, err := u.projectRepo.List()
	if err != nil {
		return nil, err
	}
	var res = []*project.Response{}
	for _, p := range projects {
		r := format(p)
		res = append(res, &r)
	}
	return res, nil
}

func (u *projectUsecase) Create(payload project.Payload) (*project.Response, error) {
	projectID, _ := uuid.NewV4()
	var p = project.Project{
		ID:          projectID,
		Name:        payload.Name,
		Description: payload.Description,
	}
	createdProject, err := u.projectRepo.Create(&p)
	if err != nil {
		return nil, err
	}
	res := format(createdProject)
	return &res, nil
}

func format(p *project.Project) project.Response {
	var res = project.Response{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
	}
	return res
}
