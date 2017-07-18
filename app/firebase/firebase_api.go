package firebase

var (
	FbAPI FirebaseAPI
)

const (
	signInEndpoint         = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
	signUpEndpoint         = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/signupNewUser?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
	getAccountInfoEndpoint = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/getAccountInfo?key=AIzaSyAkR4u8iQBLckmYNtWSx9fmJNSilyWc__A"
)

type SignUpRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type SignUpResponse struct {
	Kind         string `json:"kind"`
	IDToken      string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalID      string `json:"localId"`
}

type AccountInfoReponse struct {
	Kind  string `json:"kind"`
	Users []struct {
		LocalID          string `json:"localId"`
		Email            string `json:"email"`
		EmailVerified    bool   `json:"emailVerified"`
		ProviderUserInfo []struct {
			ProviderID  string `json:"providerId"`
			FederatedID string `json:"federatedId"`
			Email       string `json:"email"`
			RawID       string `json:"rawId"`
		} `json:"providerUserInfo"`
		PasswordHash      string  `json:"passwordHash"`
		PasswordUpdatedAt float64 `json:"passwordUpdatedAt,omitempty"`
		ValidSince        string  `json:"validSince"`
		CreatedAt         string  `json:"createdAt"`
	} `json:"users"`
}

type FirebaseAPI interface {
	Signin()
	Signup(e, p string, rs bool) (*SignUpResponse, error)
}
