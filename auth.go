package main

import (
	"encoding/json"
	"errors"
)

const (
	MOJANG_AUTH_SERVER = "https://authserver.mojang.com"
)

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
	Error             *string   `json:"error"`
	ErrorMessage      *string   `json:"errorMessage"`
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

	bytes, err := post(MOJANG_AUTH_SERVER+"/authenticate", authRequest)
	if err != nil {
		return AuthResponse{}, err
	}
	var authResponse AuthResponse
	err = json.Unmarshal(bytes, &authResponse)
	if err != nil {
		return AuthResponse{}, err
	}

	return authResponse, nil
}

type InvalidateRequest struct {
	AccessToken string `json:"accessToken"`
	ClientToken string `json:"clientToken"`
}

func invalidate(authToken string) error {
	if config.ClientToken == "" {
		return errors.New("there is no client token set")
	}

	invalidateRequest := InvalidateRequest{
		AccessToken: authToken,
		ClientToken: config.ClientToken,
	}

	bytes, err := post(MOJANG_AUTH_SERVER+"/invalidate", invalidateRequest)
	if err != nil || len(bytes) == 0 {
		return err
	}
	return nil
}
