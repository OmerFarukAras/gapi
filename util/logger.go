package util

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

// Info dt.Format("01-02 15:04:05")

func Info(s string, a any) {
	dt := time.Now()
	yellow := color.New(color.FgHiBlack).SprintFunc()
	fmt.Println("\n"+yellow(dt.Format("15:04:05"))+" "+s, "")
	fmt.Println(a)
}
