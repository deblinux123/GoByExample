package main

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("12.34", 64)

	fmt.Println(f)

	i, _ := strconv.ParseInt("12", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("ox1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("786", 0, 64)

	fmt.Println(u)

	k, _ := strconv.Atoi("123")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)

	// parsing the url
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	url, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(url.Scheme)

	fmt.Println(url.User)
	fmt.Println(url.User.Username())

	p, _ := url.User.Password()
	fmt.Println(p)

	fmt.Print(url.Host)

	host, port, _ := net.SplitHostPort(url.Host)

	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(url.Path)
	fmt.Println(url.Fragment)
	fmt.Println(url.RawQuery)
	m, _ := url.Parse(url.RawQuery)
	fmt.Println(m)

}
