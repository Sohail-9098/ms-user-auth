package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/sohail-9098/ms-user-auth/user"
	"github.com/sohail-9098/ms-user-auth/util"
)

func TestE2E_Login(t *testing.T) {
	user := &user.User{Username: "Chek", Password: "qwerty"}
	userJson, err := json.Marshal(user)
	util.HandleError("error marshal json: ", err)
	fmt.Println(string(userJson))
	req, err := http.NewRequest("GET", "http://localhost:4000/login", bytes.NewBuffer(userJson))
	util.HandleError("error creating request: ", err)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	util.HandleError("error http request: ", err)
	fmt.Println(res.StatusCode)
	body, err := io.ReadAll(res.Body)
	util.HandleError("error reading body: ", err)
	fmt.Println(string(body))
}
