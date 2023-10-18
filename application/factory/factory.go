package factory

import (
	"github.com/guhstanley/fc-go-codepix/application/usecase"
	"github.com/guhstanley/fc-go-codepix/infrastructure/repository"
	"github.com/jinzhu/gorm"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactioRepository := repository.TransactionRepositoryDb{Db: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactioRepository,
		PixRepository:         pixRepository,
	}

	return transactionUseCase
}
