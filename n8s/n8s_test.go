package n8n

import "testing"

func TestN8s(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"handle single word", args{"william"}, "w5m"},
		{"handle more than one word", args{"william gough"}, "w10h"},
		{"handle numbers", args{"2"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := N8s(tt.args.word); got != tt.want {
				t.Errorf("N8n() = %v, want %v", got, tt.want)
			}
		})
	}
}
