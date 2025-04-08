package clients

import (
	"Status418/go/clients/responses"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type AuthClientInterface interface {
	GetUserInfo(token string) (*responses.UserInfo, error)
}

type AuthClient struct {
}

func NewAuthClient() *AuthClient {
	return &AuthClient{}
}

func (auth *AuthClient) GetUserInfo(token string) (*responses.UserInfo, error) {
	apiUrl := os.Getenv("API_USERINFO")

	client := &http.Client{}

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println("Error al crear la solicitud GET:", err)
		return nil, err
	}

	req.Header.Add("Authorization", token)

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al realizar la solicitud GET:", err)
		return nil, err
	}

	if response.StatusCode != 200 {
		fmt.Println("Error al realizar la solicitud GET")
		return nil, errors.New("La peticion respondio con error")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error al leer el cuerpo de la respuesta:", err)
		return nil, err
	}

	var userInfo responses.UserInfo

	if err := json.Unmarshal(responseBody, &userInfo); err != nil {
		fmt.Println("Error al deserializar el JSON:", err)
		return nil, err
	}

	fmt.Println("Código de estado:", response.Status)

	return &userInfo, nil
}
