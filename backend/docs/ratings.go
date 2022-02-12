package docs

import (
	"github.com/elhmn/camerdevs/pkg/models/v1beta"
)

// swagger:route GET /ratings ratings idOfRatingWithoutID
// Ratings returns the list of ratings
// responses:
//   200: ratingsResponse
//   400: badRequestResponse
//   404: notFoundResponse

// swagger:route GET /ratings/{salaryID} ratings idOfRating
// Ratings returns the list of ratings
// responses:
//   200: ratingsResponse
//   400: badRequestResponse
//   404: notFoundResponse

// swagger:route GET /average-rating ratings
// Ratings returns the list of ratings
// responses:
//   200: averageRatingResponse
//   400: badRequestResponse
//   404: notFoundResponse

// This text will appear as description of your response body.
// swagger:response ratingsResponse
type RatingsResponseWrapper struct {
	// in:body
	Body v1beta.RatingResponse
}

// This text will appear as description of your response body.
// swagger:response averageRatingResponse
type AverageRatingResponseWrapper struct {
	// in:body
	Body v1beta.AverageRating
}

// swagger:response badRequestResponse
type RatingsBadRequestResponseWrapper struct {
	// in:body
	Body struct {
		// Example: bad request
		Error string `json:"error"`
	}
}

// swagger:response notFoundResponse
type RatingsNotFoundResponseWrapper struct {
	// in:body
	Body struct {
		// Example: could not found
		Error string `json:"error"`
	}
}

// swagger:parameters idOfRating
type RatingParam struct {
	//in:path
	//example: 1
	ID string `json:"id"`
}

// swagger:parameters idOfRatingWithoutID
type RatingsParam struct {
	//in:query
	//example: 1
	Page string `json:"page"`

	//in:query
	//example: 20
	Limit string `json:"limit"`

	//in:query
	//example: Realbridge
	Company string `json:"company"`

	//in:query
	//example: Recruiting Manager
	JobTitle string `json:"jobtitle"`
}
