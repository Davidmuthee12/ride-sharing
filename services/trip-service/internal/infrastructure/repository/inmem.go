package TripRepository


type inmemRepository struct {
	trips map[string]*domain.TripModel
	RideFare map[string]*domain.RideFareModel
}