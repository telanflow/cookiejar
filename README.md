## CookieJar
A cookiejar drive with Go.

## Example

```go
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/ziiber/cookiejar"
)

func main() {

	// Netscape HTTP Cookie File
	jar, _ = cookiejar.NewFileJar("cookie.txt", nil)

	client := &http.Client{
		Jar: jar,
	}

	req, _ := http.NewRequest(http.MethodGet, "https://telan.me", nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	c, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(c))
}
```
