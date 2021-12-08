package gitdata

type UserData struct {
	Login               string
	Id                  int64
	Node_id             string
	Avatar_url          string
	Gravatar_id         string
	Url                 string
	Html_url            string
	Followers_url       string
	Following_url       string
	Gists_url           string
	Starred_url         string
	Subscriptions_url   string
	Organizations_url   string
	Repos_url           string
	Events_url          string
	Received_events_url string
	Type                string
	Site_admin          bool
	Name                string
	Company             string
	Blog                string
	Location            string
	Email               string
	Hireable            string
	Bio                 string
	Twitter_username    string
	Public_repos        int
	Public_gists        int
	Followers           int
	Following           int
	Created_at          string
	Updated_at          string
}

type owner struct {
	Login               string
	Id                  int64
	Node_id             string
	Avatar_url          string
	Gravatar_id         string
	Url                 string
	Html_url            string
	Followers_url       string
	Following_url       string
	Gists_url           string
	Starred_url         string
	Subscriptions_url   string
	Organizations_url   string
	Repos_url           string
	Events_url          string
	Received_events_url string
	Type                string
	Site_admin          bool
}

type license struct {
	Key     string
	Name    string
	Spdx_id string
	Url     string
	Node_id string
}

type MainReposData struct {
	Id                int64
	Node_id           string
	Name              string
	Full_name         string
	Private           bool
	Owner             owner
	Html_url          string
	Fork              bool
	Url               string
	Forks_url         string
	Collaborators_url string
	Commits_url       string
	Issues_url        string
	Stargazers_url    string
	Created_at        string
	Updated_at        string
	Pushed_at         string
	Size              int
	Stargazers_count  int
	Language          string
	Forks_count       int
	Archived          bool
	License           license
	Open_issues       int
	Default_branch    string
}
