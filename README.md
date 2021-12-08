# gitdata
library for working with github api, written in Golang

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
