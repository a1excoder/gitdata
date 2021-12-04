package gitdata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (udP *UserData) GetUserData(user_name string) (int, error) {
	resp, err := http.Get("https://api.github.com/users/" + user_name)
	if err != nil {
		return resp.StatusCode, err
	}

	if resp.StatusCode == 404 {
		return resp.StatusCode, fmt.Errorf("Not Found")
	}

	json_decode, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}

	err = json.Unmarshal(json_decode, udP)
	return resp.StatusCode, err
}
