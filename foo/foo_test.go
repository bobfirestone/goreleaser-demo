package foo

import "testing"

func TestFoo(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
	}{
		{"v0.1.4", args{"v0.1.4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Foo(tt.args.version)
		})
	}
}
