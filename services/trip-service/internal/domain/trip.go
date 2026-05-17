package domain


type TripModel struct {
	ID primitive.ObjectID
	UserID string
	Status string
	RideFare RideFareModel
}

type TripRepository interface {
	CreateTrip(ctx context.context, trip *TripModel) (*TripModel, error)
}

type TripService interface{
	CreateTrip(ctx context.Context, fare RideFareModel) (*TripModel, error)
}