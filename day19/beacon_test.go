package day19_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day19"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/19/input", "19.in")
	if err != nil {
		log.Fatal(err)
	}
}

func TestCountBeacons(t *testing.T) {
	tests := []struct {
		file        string
		beaconCount int
		maxDistance int
	}{
		{"19.ex", 79, 3621},
		{"19.in", 303, 9621},
	}
	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			t.Fatal(err)
		}
		beaconCount, maxDistance := day19.CountBeaconsAndScannerDistance(in)
		format := "%s:\n\tgot %d, wanted %d, for beacon count\n\tgot %d wanted %d for maximum distance\n"
		if beaconCount != tt.beaconCount || maxDistance != tt.maxDistance {
			t.Fatalf(format, tt.file, beaconCount, tt.beaconCount, maxDistance, tt.maxDistance)
		} else {
			t.Logf(format, tt.file, beaconCount, tt.beaconCount, maxDistance, tt.maxDistance)
		}
	}
}
