package utils

import (
	"flag"
	"os"
)

var OldZoneId int
var NewZoneId int
var InsertIgnore bool
var WriteFile bool

func init() {
	flag.IntVar(&OldZoneId, "old-zone", -1, "old zone id")
	flag.IntVar(&NewZoneId, "new-zone", -1, "new zone id")
	flag.BoolVar(&InsertIgnore, "ignore", false, "insert statement ignore exist or not (default false)")
	flag.BoolVar(&WriteFile, "export", false, "export sql to file (default false)")
	flag.Parse()
	CheckArgs()
}
func CheckArgs() {
	if OldZoneId == -1 || NewZoneId == -1 {
		flag.Usage()
		os.Exit(1)
	}
}
