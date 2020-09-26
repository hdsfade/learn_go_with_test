package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'",expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 4)
	fmt.Println(repeated)
	//output: bbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i:=0;i<b.N;i++ {
		Repeat("a", 6)
	}
}

