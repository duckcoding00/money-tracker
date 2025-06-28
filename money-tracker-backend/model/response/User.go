package response

type (
	LoginResponse struct {
		Username    string `json:"username"`
		AccessToken string `json:"access_token"`
	}
)
