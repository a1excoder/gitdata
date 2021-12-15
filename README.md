# gitdata
library for working with github api, written in Golang

<h3>How to make a project</h3>

```bash
sudo apt install golang
mkdir project
cd project/
touch main.go
go mod init <project name>
go get github.com/a1excoder/gitdata
# write code in main.go
go build && ./<project name>
```


<h3>Example: get user data</h3>

```Golang
package main

import (
	"fmt"
	"log"

	"github.com/a1excoder/gitdata"
)

func main() {
	user := gitdata.UserData{}
	code_stat, err := user.GetUserData("a1excoder")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("code status: %d\navatar url: %s\nlocation: %s\n", code_stat, user.Avatar_url, user.Location)
}
```

<h3>output</h3>

```
code status: 200
avatar url: https://avatars.githubusercontent.com/u/49398367?v=4
location: Ukraine, Kyiv
```


<h3>Example: get repos data</h3>

```Golang
package main

import (
	"fmt"
	"log"

	gitdata "github.com/a1excoder/gitdata"
)

func main() {
	_, repos, err := gitdata.GetSingleRepos("a1excoder", "gitdata")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%s)=> writed on %s and master is %s\n", repos.Full_name, repos.Language, repos.Owner.Login)
}
```

<h3>output</h3>

```
a1excoder/gitdata)=> writed on Go and master is a1excoder
```



<h3>Example: get all user repos data</h3>

```Golang
package main

import (
	"fmt"
	"log"

	gitdata "github.com/a1excoder/gitdata"
)

func main() {
	user := gitdata.UserData{}
	_, err := user.GetUserData("a1excoder")
	if err != nil {
		log.Println(err)
		return
	}

	
	_, repos, err := gitdata.GetRepos("a1excoder", user.Public_repos)
	if err != nil {
		log.Println(err)
		return
	}

	for _, rep := range repos {
		fmt.Printf("%s", rep.Full_name)
		if rep.License.Name != "" {
			fmt.Printf("(%s)", rep.License.Name)
		}
		
		fmt.Print("\n")
	}
}
```

<h3>output</h3>

```
a1excoder/gitdata(MIT License)
a1excoder/simple
a1excoder/confman(MIT License)
a1excoder/ExtremeCodeOS(Other)
a1excoder/a1excoder
a1excoder/GetWeather
a1excoder/TubeDownload
a1excoder/GoWeather
a1excoder/newlib
a1excoder/installLinuxSoft
a1excoder/goWeatherBot(MIT License)
a1excoder/noSpyTraffic
a1excoder/a1excode-source
a1excoder/a1excode_parser
a1excoder/CNotes
a1excoder/pyWeatherBot
a1excoder/laravel_ajax
a1excoder/GoodCrypt
a1excoder/RESTful_API
```



<h3>Example: get all user repos data with error 403</h3>

```Golang
package main

import (
	"fmt"
	"log"
	"encoding/json"

	gitdata "github.com/a1excoder/gitdata"
)

func main() {
	user := gitdata.UserData{}
	_, err := user.GetUserData("a1excoder")
	if err != nil {
		log.Println(err)
		return
	}

	cod, repos, err := gitdata.GetRepos("a1excoder", user.Public_repos)
	if cod == 403 {
		data_error := &gitdata.CodeError403{}
		err := json.Unmarshal([]byte(fmt.Sprint(err)), data_error)
		if err != nil {
			log.Println(err)
		}

		fmt.Printf("message: %s\ndoc. url: %s\n", data_error.Message, data_error.Documentation_url)
		return
	}

	if err != nil {
		log.Println(err)
		return
	}

	for i, rep := range repos {
		fmt.Printf("%d) %s", i, rep.Full_name)
		if rep.License.Name != "" {
			fmt.Printf("(%s)", rep.License.Name)
		}

		fmt.Print("\n")
	}
}

```

<h3>output</h3>

```
message: API rate limit exceeded for 0.0.0.0. (But here's the good news: Authenticated requests get a higher rate limit. Check out the documentation for more details.)
doc. url: https://docs.github.com/rest/overview/resources-in-the-rest-api#rate-limiting
```

