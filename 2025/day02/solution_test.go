package day02

import "testing"

const example = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestPart1(t *testing.T) {
	got := Solution{}.Part1(example)
	want := 1227775554

	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Solution{}.Part2(example)
	want := 4174379265

	if got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}

func TestGenerateDoubled(t *testing.T) {
	doubled := generateDoubled(100)
	// Should be: 11, 22, 33, 44, 55, 66, 77, 88, 99
	want := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}

	if len(doubled) != len(want) {
		t.Errorf("generateDoubled(100) got %d items, want %d", len(doubled), len(want))
		return
	}

	for i, v := range want {
		if doubled[i] != v {
			t.Errorf("generateDoubled(100)[%d] = %d, want %d", i, doubled[i], v)
		}
	}
}

func TestParse(t *testing.T) {
	ranges := Parse("11-22,95-115")

	if len(ranges) != 2 {
		t.Fatalf("Parse() got %d ranges, want 2", len(ranges))
	}

	if ranges[0].Lo != 11 || ranges[0].Hi != 22 {
		t.Errorf("ranges[0] = %v, want {11, 22}", ranges[0])
	}

	if ranges[1].Lo != 95 || ranges[1].Hi != 115 {
		t.Errorf("ranges[1] = %v, want {95, 115}", ranges[1])
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution{}.Part1(example)
	}
}

func BenchmarkPart1_RealInput(b *testing.B) {
	input := `2558912-2663749,1-19,72-85,82984-100358,86-113,193276-237687,51-69,779543-880789,13004-15184,2768-3285,4002-4783,7702278-7841488,7025-8936,5858546565-5858614010,5117615-5149981,4919-5802,411-466,126397-148071,726807-764287,7454079517-7454227234,48548-61680,67606500-67729214,9096-10574,9999972289-10000034826,431250-455032,907442-983179,528410-680303,99990245-100008960,266408-302255,146086945-146212652,9231222-9271517,32295166-32343823,32138-36484,4747426142-4747537765,525-652,333117-414840,13413537-13521859,1626-1972,49829276-50002273,69302-80371,8764571787-8764598967,5552410836-5552545325,660-782,859-1056`

	for i := 0; i < b.N; i++ {
		Solution{}.Part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution{}.Part2(example)
	}
}
