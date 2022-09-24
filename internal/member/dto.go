package member

type CreateRequest struct {
	UserName       string `json:"userName" validate:"required" example:"andy"`
	MembershipType string `json:"membershipType" validate:"required" example:"toss"`
}

type CreateResponse struct {
	ID             string `json:"id" example:"354660dc-f798-11ec-b939-0242ac120002"`
	UserName       string `json:"user_name" example:"andy"`
	MembershipType string `json:"membership_type" example:"toss"`
}

type UpdateRequest struct {
	ID             string `json:"ID" validate:"omitempty"`
	UserName       string `json:"userName" validate:"required"`
	MembershipType string `json:"membershipType" validate:"required"`
}

type UpdateResponse struct {
	ID             string `json:"id"`
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}

type GetResponse struct {
	ID             string `json:"id" example:"354660dc-f798-11ec-b939-0242ac120002"`
	UserName       string `json:"user_name" example:"andy"`
	MembershipType string `json:"membership_type"  example:"toss"`
}

type Fail400GetResponse struct {
	Message string `json:"message" example:"Bad Request"`
}
