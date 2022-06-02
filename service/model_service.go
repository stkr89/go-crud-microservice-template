package service

import (
	"github.com/go-kit/kit/log"
	"github.com/stkr89/modelsvc/common"
	"github.com/stkr89/modelsvc/dao"
	"github.com/stkr89/modelsvc/models"
	"github.com/stkr89/modelsvc/types"
)

// ModelService interface
type ModelService interface {
	Create(request *types.CreateRequest) (*types.CreateResponse, error)
	Get(request *types.GetRequest) (*types.GetResponse, error)
	List(request *types.ListRequest) (*types.ListResponse, error)
	Update(request *types.UpdateRequest) (*types.UpdateResponse, error)
	Delete(request *types.DeleteRequest) error
}

type ModelServiceImpl struct {
	logger   log.Logger
	modelDao dao.ModelDao
}

func NewModelServiceImpl() *ModelServiceImpl {
	return &ModelServiceImpl{
		logger:   common.NewLogger(),
		modelDao: dao.NewModelDaoImpl(),
	}
}

func NewModelServiceImplArgs(logger log.Logger, mathDao dao.ModelDao) ModelService {
	return &ModelServiceImpl{
		logger:   logger,
		modelDao: mathDao,
	}
}

func (m ModelServiceImpl) Delete(request *types.DeleteRequest) error {
	err := m.modelDao.Delete(request.ID)
	if err != nil {
		m.logger.Log("error", err)
		return err
	}

	return nil
}

func (m ModelServiceImpl) Update(request *types.UpdateRequest) (*types.UpdateResponse, error) {
	_, err := m.modelDao.Update(&models.Model{})
	if err != nil {
		m.logger.Log("error", err)
		return nil, err
	}

	return &types.UpdateResponse{}, nil
}

func (m ModelServiceImpl) List(request *types.ListRequest) (*types.ListResponse, error) {
	_, err := m.modelDao.List()
	if err != nil {
		m.logger.Log("error", err)
		return nil, err
	}

	return &types.ListResponse{}, err
}

func (m ModelServiceImpl) Get(request *types.GetRequest) (*types.GetResponse, error) {
	_, err := m.modelDao.Get(request.ID)
	if err != nil {
		m.logger.Log("error", err)
		return nil, err
	}

	return &types.GetResponse{}, nil
}

func (m ModelServiceImpl) Create(request *types.CreateRequest) (*types.CreateResponse, error) {
	_, err := m.modelDao.Create(&models.Model{})
	if err != nil {
		m.logger.Log("error", err)
		return nil, err
	}

	return &types.CreateResponse{}, nil
}
