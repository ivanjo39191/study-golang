package main

import (
	"fmt"
)

func main() {
	var s string
	s = "how are you?"
	fmt.Println(s)
	s = `how are you?`
	fmt.Println(s)
	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"
	fmt.Println(s)
	s = `
<html>
	<body>"Hello"</body>
</html>
`
	fmt.Println(s)

	fmt.Println("C:\\my\\dir\\file")
	fmt.Println(`C:\my\dir\file`)
}
