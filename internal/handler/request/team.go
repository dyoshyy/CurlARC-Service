package request

type CreateTeamRequest struct {
	Name string `json:"name"`
}

type UpdateTeamRequest struct {
	Name string `json:"name"`
}
