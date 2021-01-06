package transport

type ErrResponse struct {
	Message string
}
type GetUserSummaryRequest struct {
	UserUUID string `json:"userUUID"`
}

type GetUserSummaryResponse struct {
	Summary *UserSummary `json:"summary"`
}

type GetUserVaccineDetailsRequest struct {
	UserUUID string `json:"userUUID"`
}

type GetUserVaccineDetailsResponse struct {
	VaccineDetails []*VaccineDetails `json:"vaccineDetails"`
}

type GetCountryVaccinesRequest struct {
	Country string `json:"userDetails"`
}

type GetCountryVaccinesResponse struct {
	RequiredVaccineDetails []*RequiredVaccineDetails `json:"requiredVaccineDetails"`
}

// model objects
type UserSummary struct {
	TotalVaccinations  int64             `json:"totalVaccinations"`
	LatestVaccinations []*VaccineDetails `json:"latestVaccinations"`
}

type VaccineDetails struct {
	Name            string `json:"name"`
	VaccineID       string `json:"vaccineID"`
	UserID          string `json:"userID"`
	UserIDType      string `json:"userIDType"`
	Authority       string `json:"authority"`
	AuthorityType   string `json:"authorityType"`
	Country         string `json:"country"`
	VaccinationDate int64  `json:"vaccinationDate"`
}

type RequiredVaccineDetails struct {
	Name                   string                    `json:"name"`
	Description            string                    `json:"description"`
	AuthorisedBodiesNearby []*AuthorisedBodiesNearby `json:"authorisedBodiesNearby"`
}

type AuthorisedBodiesNearby struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Call     string `json:"call"`
}
