package date

import "time"

const (
	LayoutDateTime    string = "2006-01-02T15:04:05.000000"
	LocationAsiaTokyo string = "Asia/Tokyo"
)

func LocationTokyo() *time.Location {
	loc, _ := time.LoadLocation(LocationAsiaTokyo)
	return loc
}
