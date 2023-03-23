package averagesalary

import (
	"testing"
)

/*
Write a function that returns the average salary from a set of salaries,
excluding the max and minimum salaries.
*/
func TestAverage(t *testing.T) {

	salaries := []int{300, 500, 1000, 900, 200}
	got := Average(salaries)
	want := 566

	if want != int(got) {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
