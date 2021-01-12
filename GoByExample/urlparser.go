package main

import (
	"fmt"
	"net"
	"net/url"
	"path"
)

func urlParser() {

	s := "https://jira-sd-test.mc1.oracleiaas.com/projects/COM/queues/custom/205/COM-62247"
	fmt.Println(s)

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme: ", u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println("User Creds: ", p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("Host: ", host)
	fmt.Println("Port: ", port)

	fmt.Println("Path:", u.Path)
	fmt.Println("Path Base:", string(path.Base(u.Path)))
	fmt.Println("Fragment:", u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}
