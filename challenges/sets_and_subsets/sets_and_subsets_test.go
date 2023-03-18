package setsandsubsets

import (
	"testing"
)

/*
Write a function which checks to see whether a subset
of flights is included in another set.
*/
func TestIsSubset(t *testing.T) {
	t.Run("Should return true for subset", func(t *testing.T) {
		firstFlights := []Flight{
			{Origin: "GLA", Destination: "CDG", Price: 1000},
			{Origin: "GLA", Destination: "JFK", Price: 5000},
			{Origin: "GLA", Destination: "SNG", Price: 3000},
		}

		secondFlights := []Flight{
			{Origin: "GLA", Destination: "CDG", Price: 1000},
			{Origin: "GLA", Destination: "JFK", Price: 5000},
			{Origin: "GLA", Destination: "SNG", Price: 3000},
			{Origin: "GLA", Destination: "AMS", Price: 500},
		}

		got := IsSubset(firstFlights, secondFlights)
		want := true

		if got != want {
			t.Errorf("Got %t but wanted %t", got, want)
		}
	})

	t.Run("Should return false if not a subset", func(t *testing.T) {
		firstFlights := []Flight{
			{Origin: "GLA", Destination: "CDG", Price: 1000},
			{Origin: "GLA", Destination: "JFK", Price: 5000},
			{Origin: "GLA", Destination: "SNG", Price: 3000},
		}

		secondFlights := []Flight{
			{Origin: "GLA", Destination: "CDG", Price: 1000},
			{Origin: "GLA", Destination: "SNG", Price: 3000},
			{Origin: "GLA", Destination: "AMS", Price: 500},
		}

		got := IsSubset(firstFlights, secondFlights)
		want := false

		if got != want {
			t.Errorf("Got %t but wanted %t", got, want)
		}
	})

	t.Run("Should return false if empty child set", func(t *testing.T) {
		firstFlights := []Flight{}

		secondFlights := []Flight{
			{Origin: "GLA", Destination: "CDG", Price: 1000},
			{Origin: "GLA", Destination: "SNG", Price: 3000},
			{Origin: "GLA", Destination: "AMS", Price: 500},
		}

		got := IsSubset(firstFlights, secondFlights)
		want := false

		if got != want {
			t.Errorf("Got %t but wanted %t", got, want)
		}
	})

	t.Run("Should return false if empty parent set", func(t *testing.T) {
		firstFlights := []Flight{
			{Origin: "GLA", Destination: "CDG", Price: 1000},
			{Origin: "GLA", Destination: "SNG", Price: 3000},
			{Origin: "GLA", Destination: "AMS", Price: 500},
		}

		secondFlights := []Flight{}

		got := IsSubset(firstFlights, secondFlights)
		want := false

		if got != want {
			t.Errorf("Got %t but wanted %t", got, want)
		}
	})

	t.Run("Should handle duplicate child flights", func(t *testing.T) {
		firstFlights := []Flight{
			{Origin: "GLA", Destination: "CDG", Price: 1000},
			{Origin: "GLA", Destination: "CDG", Price: 1000},
		}

		secondFlights := []Flight{
			{Origin: "GLA", Destination: "CDG", Price: 1000},
			{Origin: "GLA", Destination: "SNG", Price: 3000},
			{Origin: "GLA", Destination: "AMS", Price: 500},
		}

		got := IsSubset(firstFlights, secondFlights)
		want := false

		if got != want {
			t.Errorf("Got %t but wanted %t", got, want)
		}
	})

	t.Run("Should handle duplicate parent flights", func(t *testing.T) {
		firstFlights := []Flight{
			{Origin: "GLA", Destination: "CDG", Price: 1000},
		}

		secondFlights := []Flight{
			{Origin: "GLA", Destination: "CDG", Price: 1000},
			{Origin: "GLA", Destination: "CDG", Price: 1000},
			{Origin: "GLA", Destination: "SNG", Price: 3000},
			{Origin: "GLA", Destination: "AMS", Price: 500},
		}

		got := IsSubset(firstFlights, secondFlights)
		want := true

		if got != want {
			t.Errorf("Got %t but wanted %t", got, want)
		}
	})
}
