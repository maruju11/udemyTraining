package main

import "fmt"

func main() {
	name := "jorge maruju"
	jorge := "11"
	fmt.Println(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Go Palmeiras</title>
	</head>
	<body>
	<h1>` +
		name +
		`</h1>
<h1>` +
		jorge +
		`</h1>
	</body>
	</html>
	`)
}
