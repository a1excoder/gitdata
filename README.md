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
