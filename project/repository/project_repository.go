package repository

import (
	"log"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/totoro081295/daily-report-api/project"
	"github.com/totoro081295/daily-report-api/status"
)

type projectRepository struct {
	Conn *gorm.DB
}

// NewProjectRepository mount project repository
func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{
		Conn: db,
	}
}

// ProjectRepository repository interface
type ProjectRepository interface {
	Get(id uuid.UUID) (*project.Project, error)
	List() ([]*project.Project, error)
	ListByIDs(ids []uuid.UUID) ([]*project.Project, error)
	Create(p *project.Project) (*project.Project, error)
	Update(p *project.Project) error
	Delete(id uuid.UUID) error
}

func (m *projectRepository) Get(id uuid.UUID) (*project.Project, error) {
	var p project.Project
	err := m.Conn.Model(&p).Where("id = ?", id).Find(&p).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Wrap(status.ErrNotFound, err.Error())
	} else if err != nil {
		return nil, err
	}
	return &p, nil
}

func (m *projectRepository) List() ([]*project.Project, error) {
	var p []*project.Project
	err := m.Conn.Model(&p).Find(&p).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (m *projectRepository) ListByIDs(ids []uuid.UUID) ([]*project.Project, error) {
	var p []*project.Project
	err := m.Conn.Model(&p).Where("id in (?)", ids).Find(&p).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (m *projectRepository) Create(p *project.Project) (*project.Project, error) {
	err := m.Conn.Create(&p).Error
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return p, nil
}

func (m *projectRepository) Update(p *project.Project) error {
	err := m.Conn.Model(p).Update(p).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *projectRepository) Delete(id uuid.UUID) error {
	var p project.Project
	err := m.Conn.Model(&p).Where("id = ?", id).Delete(&p).Error
	if err != nil {
		return errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return nil
}
