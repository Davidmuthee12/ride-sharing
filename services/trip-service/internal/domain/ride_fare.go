package domain

 type RideFareModel struct {
	ID primitive.ObjectID
	UserID string
	PackageSlug string  //ex: van, Luxury, sedan
	TotalPriceInCents float64
 }