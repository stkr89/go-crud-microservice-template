package dao

import (
	"github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/stkr89/mathsvc/common"
	"github.com/stkr89/mathsvc/config"
	"github.com/stkr89/mathsvc/models"
	"gorm.io/gorm"
)

//go:generate mocker --dst ../mock/model_dao_mock.go --pkg mock model_dao.go ModelDao
type ModelDao interface {
	Create(obj *models.Model) (*models.Model, error)
	Get(id uuid.UUID) (*models.Model, error)
	List() ([]*models.Model, error)
	Update(obj *models.Model) (*models.Model, error)
	Delete(id uuid.UUID) error
}

type ModelDaoImpl struct {
	logger log.Logger
	db     gorm.DB
}

func NewModelDaoImpl() *ModelDaoImpl {
	return &ModelDaoImpl{
		logger: common.NewLogger(),
		db:     config.NewDB(),
	}
}

func NewModelDaoImplArgs(db gorm.DB) ModelDaoImpl {
	return ModelDaoImpl{
		db: db,
	}
}

func (m ModelDaoImpl) Delete(id uuid.UUID) error {
	result := m.db.Delete(&models.Model{ID: id})
	if result.Error != nil {
		m.logger.Log("message", "unable to delete", "error", result.Error)
		return common.SomethingWentWrong
	}

	return nil
}

func (m ModelDaoImpl) Update(obj *models.Model) (*models.Model, error) {
	result := m.db.Save(obj)
	if result.Error != nil {
		m.logger.Log("message", "unable to update", "error", result.Error)
		return nil, common.SomethingWentWrong
	}

	return obj, nil
}

func (m ModelDaoImpl) List() ([]*models.Model, error) {
	var objs []*models.Model
	result := m.db.Find(&objs)
	if result.Error != nil {
		m.logger.Log("message", "unable to list", "error", result.Error)
		return nil, common.SomethingWentWrong
	}

	return objs, nil
}

func (m ModelDaoImpl) Get(id uuid.UUID) (*models.Model, error) {
	obj := models.Model{ID: id}

	result := m.db.First(&obj)
	if result.Error != nil {
		m.logger.Log("message", "unable to get", "error", result.Error)
		return nil, common.SomethingWentWrong
	}

	return &obj, nil
}

func (m ModelDaoImpl) Create(obj *models.Model) (*models.Model, error) {
	result := m.db.Create(&obj)
	if result.Error != nil {
		m.logger.Log("message", "unable to create", "error", result.Error)
		return nil, common.SomethingWentWrong
	}

	m.logger.Log("message", "created successfully", "return", obj.ID)

	return obj, nil
}