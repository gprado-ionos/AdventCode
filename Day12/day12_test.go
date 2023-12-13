package main

import "testing"

var recordsEntry = []string{
	"???.### 1,1,3",
	".??..??...?##. 1,1,3",
	"?#?#?#?#?#?#?#? 1,3,1,6",
	"????.#...#... 4,1,1",
	"????.######..#####. 1,6,5",
	"?###???????? 3,2,1",
}

func Test_getEmptySpaceToExpand(t *testing.T) {

	records := buildRecords(recordsEntry)
	arrangements := getPossibleArrangements(records)

	if arrangements != 21 {
		t.Errorf("Result was incorrect, got: %d galaxies, want: %d.", arrangements, 21)
	}
}
