package main

type AuthAgent struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

type AuthRequest struct {
	Agent       AuthAgent `json:"agent"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	ClientToken *string   `json:"clientToken"`
	RequestUser bool      `json:"requestUser"`
}

type AuthResponse struct {
	AccessToken       string    `json:"accessToken"`
	ClientToken       string    `json:"clientToken"`
	SelectedProfile   Profile   `json:"selectedProfile"`
	AvailableProfiles []Profile `json:"availableProfiles"`
	Error             string    `json:"error"`
	ErrorMessage      string    `json:"errorMessage"`
}

type Profile struct {
	Agent string `json:"agent"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

func auth(username string, password string) (AuthResponse, error) {
	authRequest := AuthRequest{
		Agent: AuthAgent{
			Name:    "minecraft",
			Version: 1,
		},
		Username: username,
		Password: password,
	}
	if config.ClientToken != "" {
		authRequest.ClientToken = &config.ClientToken
	}

	var authResponse AuthResponse
	err := post("https://authserver.mojang.com/authenticate", authRequest, &authResponse)
	if err != nil {
		return AuthResponse{}, err
	}

	return authResponse, nil
}
