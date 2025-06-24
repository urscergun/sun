# sun
A simple Go wrapper for https://sunrisesunset.io/api/ to get sunrise and sunset information based on date, latitude and longitude.

Example:

```golang
package main

import (
	"fmt"
	"time"

	"github.com/urscergun/sun"
)

func main() {
	now := time.Now()

	var date string = now.Format("2006-01-02")
	var latitude string = "45.763935"
	var longitude string = "21.2238336"

	sunTimes, err := sun.FetchSunTimes(date, latitude, longitude)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Date: %s\nLocation: %s, %s\nSunrise: %s\nSunset: %s\n", date, latitude, longitude, sunTimes.Sunrise, sunTimes.Sunset)

}
