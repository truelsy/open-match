package main

import (
	"ns-open-match/internal/appmain"
	"ns-open-match/nbanow/director"
)

func main() {
	appmain.RunApplication("director", director.BindService)
}
