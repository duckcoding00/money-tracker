package request

type (
	UserRequest struct {
		Username string `json:"username" validate:"required,min=8"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	ResetUser struct {
		Username string `json:"username" validate:"required,min=8"`
	}

	VerifyToken struct {
		Username string `json:"username" validate:"required,min=8"`
		Token    string `json:"token" validate:"required"`
	}

	NewPassword struct {
		Password string `json:"password" validate:"required,min=8"`
	}
)
