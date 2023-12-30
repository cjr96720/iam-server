package domain

type IAMRepositoryInterface interface {
	CreateApplication(app Application) error
	GetApplications() ([]Application, error)
}

type IAMServiceInterface interface {
	CreateApplication(appName string) (Application, error)
	GetApplications() ([]Application, error)
}
