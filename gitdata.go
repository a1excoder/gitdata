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

	if repos_number <= 100 {
		resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos?per_page=100&sort=updated&direction=desc&page=1", user_name))
		if err != nil {
			return resp.StatusCode, nil, err
		}

		if resp.StatusCode == 404 {
			return resp.StatusCode, nil, fmt.Errorf("Not Found")
		}

		json_decode, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, nil, err
		}

		err = json.Unmarshal(json_decode, &repos)
		if err != nil {
			return resp.StatusCode, nil, err
		}

		return 200, repos, nil
	}

	temp, temp2 := float32(repos_number)/100.0, repos_number
	if temp == float32(int(temp)) {
		temp = float32(int(temp))
	} else {
		temp = float32(int(temp) + 1)
	}
	fmt.Println(temp)
	for i := 1; i <= int(temp); i++ {
		fmt.Println(temp2)
		temp_repos := make([]MainReposData, 0, temp2)

		resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos?per_page=100&sort=updated&direction=desc&page=%d", user_name, i))
		if err != nil {
			return resp.StatusCode, nil, err
		}

		if resp.StatusCode == 404 {
			return resp.StatusCode, nil, fmt.Errorf("Not Found")
		}

		json_decode, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, nil, err
		}

		err = json.Unmarshal(json_decode, &temp_repos)
		if err != nil {
			return resp.StatusCode, nil, err
		}

		for _, rep_ := range temp_repos {
			repos = append(repos, rep_)
		}

		temp2 -= 100
	}

	return 200, repos, nil
}
