package components

import (
	"testing"
	"time"
)

func TestGetThisMonday(t *testing.T) {
	t.Run("should return the date of the current Monday", func(t *testing.T) {
		got := getThisMonday()
		if got.Weekday() != time.Monday {
			t.Errorf("GetThisMonday() = %v; want a Monday", got)
		}
	})
}