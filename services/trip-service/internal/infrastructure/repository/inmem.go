package TripRepository

import "ride-sharing/services/trip-service/internal/domain"

type inmemRepository struct {
	trips    map[string]*domain.TripModel
	RideFare map[string]*domain.RideFareModel
}