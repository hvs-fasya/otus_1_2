package otus_1_2

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	svc := NewPackerSvc()
	tests := []struct {
		packed string
		want   string
	}{
		{`a4bc2d5e`, `aaaabccddddde`},
		{`a11bc2d5e`, `aaaaaaaaaaabccddddde`},
		{`a11bc2d11`, `aaaaaaaaaaabccddddddddddd`},
		{`abcd`, `abcd`},
		{`45`, ``},
		{``, ``},
		{`qwe\4\5`, `qwe45`},
		{`qwe\\a`, `qwe\a`},
		{`qwe\\5`, `qwe\\\\\`},
		{`\\5`, `\\\\\`},
		{`4\\5`, ``},
	}
	for _, test := range tests {
		res := svc.Unpack(test.packed)
		if res != test.want {
			t.Errorf("Unpack() = %s, want %s", res, test.want)
		}
	}
}
