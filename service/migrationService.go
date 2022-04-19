package service

import "banking/domain"

type DefaultMigrationService struct {
	repo domain.MigrationRepository
}

type MigrationService interface {
	Prepare() bool
}

// -------------------------------------
func (m DefaultMigrationService) Prepare() bool {
	return m.repo.Prepare()
}

func NewDefaultMigrationService(r domain.MigrationRepository) DefaultMigrationService {
	return DefaultMigrationService{
		repo: r,
	}
}
