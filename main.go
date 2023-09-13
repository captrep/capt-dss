package main

import (
	"capt-dss/dss"
	"fmt"
)

func main() {
	h := dss.NewHandler()

	serverStr := fmt.Sprintf(":8000")
	s := dss.NewServer(h)

	s.Serve(serverStr)
}
