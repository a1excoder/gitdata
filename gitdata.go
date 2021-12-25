package gitdata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (udP *UserData) GetUserData(user_name string) (int, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s", user_name))
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

func GetRepos(user_name string, repos_number int) (int, []MainReposData, error) {
	repos := make([]MainReposData, 0, repos_number)
	tmp_repos := make([]MainReposData, 0, 100)

	tmp := float32(repos_number) / 100.0
	if tmp != float32(int(tmp)) {
		tmp = float32(int(tmp) + 1)
	}
	for i := 1; i <= int(tmp); i++ {

		resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos?per_page=100&sort=updated&direction=desc&page=%d", user_name, i))
		if err != nil {
			return resp.StatusCode, nil, err
		}

		if resp.StatusCode == 404 {
			return resp.StatusCode, nil, fmt.Errorf("Not Found")
		}

		if resp.StatusCode == 403 {
			json_decode, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return resp.StatusCode, nil, err
			}

			return resp.StatusCode, nil, fmt.Errorf(string(json_decode))
		}

		if resp.StatusCode != 200 {
			return resp.StatusCode, nil, fmt.Errorf("error code %d", resp.StatusCode)
		}

		json_decode, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, nil, err
		}

		err = json.Unmarshal(json_decode, &tmp_repos)
		if err != nil {
			return resp.StatusCode, nil, err
		}
		repos = append(repos, tmp_repos...)
		tmp_repos = tmp_repos[:0]
	}

	return 200, repos, nil
}

func GetSingleRepos(user_name, repos_name string) (int, MainReposData, error) {
	repos := MainReposData{}

	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s", user_name, repos_name))
	if err != nil {
		return resp.StatusCode, repos, err
	}

	if resp.StatusCode == 404 {
		return resp.StatusCode, repos, fmt.Errorf("Not Found")
	}

	json_decode, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, repos, err
	}

	err = json.Unmarshal(json_decode, &repos)
	if err != nil {
		return resp.StatusCode, repos, err
	}

	return resp.StatusCode, repos, nil
}
