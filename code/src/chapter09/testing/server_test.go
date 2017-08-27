package testing

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	testCases := []struct {
		gmt  string
		loc  string
		want string
	}{
		{"12:31", "Europe/Zuri", "13:31"},
		{"12:31", "America/New_York", "7:31"},
		{"08:08", "Australia/Sydney", "18:08"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s in %s", tc.gmt, tc.loc), func(t *testing.T) {
			t.Parallel()
			loc, err := time.LoadLocation(tc.loc)
			if err != nil {
				t.Fatal("could not load location")
			}
			gmt, _ := time.Parse("15:04", tc.gmt)
			if got := gmt.In(loc).Format("15:04"); got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}
