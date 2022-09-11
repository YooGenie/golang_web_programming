package internal

type Request struct {
	ID             string `json:"ID" validate:"omitempty"`
	UserName       string `json:"userName" validate:"required"`
	MembershipType string `json:"membershipType" validate:"required"`
}

type CreateResponse struct {
	ID             string `json:"id"`
	MembershipType string `json:"membership_type"`
}

type UpdateResponse struct {
	ID             string `json:"id"`
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}

type GetResponse struct {
	ID             string `json:"id"`
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}
