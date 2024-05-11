package dto

type SignInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}
