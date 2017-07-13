package main_test

import (
	"fmt"
	"go-portal-api/routers"
	"io"
	"net/http/httptest"
)

var (
	server   *httptest.Server
	reader   io.Reader //Ignore this for now
	usersUrl string
)

func init() {

	server = httptest.NewServer(routers.InitRoutes()) //Creating new server with the user handlers

	usersUrl = fmt.Sprintf("%s/users", server.URL) //Grab the address for the API endpoint

	AuthTest_Init()
}

/*

func TestGetAllUsers_ShouldReturnOK(t *testing.T) {

	request, err := http.NewRequest("GET", usersUrl, reader) //Create request with JSON body

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
	}
}

func TestCreateUser_ShouldReturnOK(t *testing.T) {
	userJson := `{"username": "dennis", "password": "junior", "email": "dennis@gmail.com"}`

	reader = strings.NewReader(userJson) //Convert string to reader

	request, err := http.NewRequest("POST", usersUrl, reader) //Create request with JSON body

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
	}
}

func TestCreateUser_EmptyUsername_ReturnsError400(t *testing.T) {
	userJson := `{"username": "", "password": "junior", "email": "dennis@gmail.com"}`

	reader = strings.NewReader(userJson) //Convert string to reader

	request, err := http.NewRequest("POST", usersUrl, reader) //Create request with JSON body

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if res.StatusCode == 200 {
		t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
	}
}

func TestCreateUser_EmptyPassword_ReturnsError400(t *testing.T) {
	userJson := `{"username": "testuser", "password": "", "email": "dennis@gmail.com"}`

	reader = strings.NewReader(userJson) //Convert string to reader

	request, err := http.NewRequest("POST", usersUrl, reader) //Create request with JSON body

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if res.StatusCode == 200 {
		t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
	}
}

func TestCreateUser_EmptyEmail_ReturnsError400(t *testing.T) {
	userJson := `{"username": "testusername", "password": "testpassword", "email": ""}`

	reader = strings.NewReader(userJson) //Convert string to reader

	request, err := http.NewRequest("POST", usersUrl, reader) //Create request with JSON body

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if res.StatusCode == 200 {
		t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
	}
}

//TODO Nil conditions
func TestCreateUser_NilEmail_ReturnsError400(t *testing.T) {
	userJson := `{"username": "testusername", "password": "testpassword"}`

	reader = strings.NewReader(userJson) //Convert string to reader

	request, err := http.NewRequest("POST", usersUrl, reader) //Create request with JSON body

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if res.StatusCode == 200 {
		t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
	}
}
*/
