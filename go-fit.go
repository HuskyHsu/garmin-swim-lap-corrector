package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/tormoder/fit"
)

func test() {
	// Read our FIT test file data
	testData, err := os.ReadFile("17365012152_ACTIVITY_NEW.FIT")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Decode the FIT file data
	fit, err := fit.Decode(bytes.NewReader(testData))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Inspect the TimeCreated field in the FileId message
	fmt.Println(fit.FileId.TimeCreated)

	// Get the actual activity
	activity, err := fit.Activity()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the latitude and longitude of the first Record message
	// fmt.Println(len(activity.Laps))
	for _, lap := range activity.Laps {
		fmt.Printf("%v - %.2fs - %dm\n", lap.StartTime, float64(lap.TotalTimerTime)/1000, lap.TotalDistance/100)
		if lap.NumActiveLengths > 0 {
			for _, length := range activity.Lengths[lap.FirstLengthIndex : lap.FirstLengthIndex+lap.NumActiveLengths] {
				fmt.Printf("  %v - %.2fs\n", length.StartTime, float64(length.TotalTimerTime)/1000)
			}
		}
	}
}
