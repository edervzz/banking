package domain

type MigrationRepository interface {
	Prepare() bool
}
