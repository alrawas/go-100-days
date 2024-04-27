package main

import (
	"alrawas/100daysofgo/maths/svg"
	"os"
	"time"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
