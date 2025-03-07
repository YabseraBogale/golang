package main

import (
	"github.com/YabseraBogale/golang/browser_engine/parser"
)

func main() {
	p := parser.Parser{
		Pos: 0,
		Input: `
			<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>

    </body>
</html>
	`,
	}

}
