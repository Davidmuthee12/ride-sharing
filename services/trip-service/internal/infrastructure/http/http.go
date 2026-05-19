package http

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"
)

type HttpHandler struct {
	Service domain.TripService
}

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

type previewTripResponse struct {
	Route     *types.Route `json:"route"`
	RideFares []any        `json:"rideFares"`
}

func (s *HttpHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	ctx := r.Context()

	t, err := s.Service.GetRoute(ctx, &reqBody.Pickup, &reqBody.Destination)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to fetch route", http.StatusBadGateway)
		return
	}
	if t == nil || len(t.Routes) == 0 {
		http.Error(w, "route not found", http.StatusBadGateway)
		return
	}

	route := t.Routes[0]
	coordinates := make([]*types.Coordinate, 0, len(route.Geometry.Coordinates))
	for _, coordinate := range route.Geometry.Coordinates {
		if len(coordinate) < 2 {
			continue
		}

		coordinates = append(coordinates, &types.Coordinate{
			Longitude: coordinate[0],
			Latitude:  coordinate[1],
		})
	}

	response := previewTripResponse{
		Route: &types.Route{
			Distance: route.Distance,
			Duration: route.Duration,
			Geometry: []*types.Geometry{
				{Coordinates: coordinates},
			},
		},
		RideFares: []any{},
	}

	writeJSON(w, http.StatusOK, response)
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
