package api

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
		PasswordHash      string `json:"passwordHash"`
		PasswordUpdatedAt int64  `json:"passwordUpdatedAt"`
		ValidSince        string `json:"validSince"`
		CreatedAt         string `json:"createdAt"`
	} `json:"users"`
}
