package core

import (
	"reflect"
	"testing"
)

func Test_toStrings(t *testing.T) {
	type args struct {
		e []*Extract
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
			[]*Extract{
					{
						Object:"123",
					},
					{
						Object:"456",
					},
				},
			},
			want: []string{"123", "456"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toStrings(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
