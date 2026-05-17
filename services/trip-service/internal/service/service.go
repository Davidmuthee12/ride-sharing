 package service

 type service struct {
	repo domain.TripRepository
 }

 func NewService(repo domain.TripRepository) *service {
	return &service {
		repo: repo,
	}
 }