package split

import (
	"reflect"
	"testing"
)

// // 测试
// func TestSplit(t *testing.T) {
// 	got := Split("我爱你", "爱")
// 	want := []string{"我", "你"}
// 	// got := Split("a:b:c", ":")
// 	// want := []string{"a", "b", "c"}
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("want:%v got:%v", want, got)
// 	}
// }

func TestSplit(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{
		"simple":     {input: "我爱你", sep: "爱", want: []string{"我", "你"}},
		"multi sep":  {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"multi sep2": {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"chinese":    {input: "沙河有沙河粉", sep: "沙", want: []string{"", "河有", "河粉"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name: %s failed, want:%v got:%v", name, tc.want, got)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	// b.N 不是固定的数
	for i := 0; i < b.N; i++ {
		Split("沙河有沙河粉", "沙")
	}
}
