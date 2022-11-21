package pipeline

import (
	"fmt"
	"testing"
)

func TestLaunchPipeline(t *testing.T) {
	tableTests := [][]int{
		{3, 14},
		{5, 55},
		{10, 385},
	}

	var res int
	for i, test := range tableTests {
		t.Run(fmt.Sprintf("Test ke %d", i+1), func(t *testing.T) {
			res = LaunchPipeline(test[0])
			if res != test[1] {
				t.Fail()
			}

			t.Logf("%d == %d", res, test[1])
		})
	}
}
