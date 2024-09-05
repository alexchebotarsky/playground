package wavyword

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		word     string
		waviness int
	}
	tests := []struct {
		name    string
		args    args
		want    *WavyWord
		wantErr bool
	}{
		{
			name: "should create wavy word with 5 lines for waviness=2",
			args: args{
				word:     "wavyword",
				waviness: 2,
			},
			want: &WavyWord{
				Word: "wavyword",
				Lines: [][]string{
					{" ", " ", " ", " ", " ", " ", "r", " "},
					{" ", " ", " ", " ", " ", "o", " ", "d"},
					{"w", " ", " ", " ", "w", " ", " ", " "},
					{" ", "a", " ", "y", " ", " ", " ", " "},
					{" ", " ", "v", " ", " ", " ", " ", " "},
				},
			},
		},
		{
			name: "should create wavy word with 3 lines for waviness=1",
			args: args{
				word:     "wavyword",
				waviness: 1,
			},
			want: &WavyWord{
				Word: "wavyword",
				Lines: [][]string{
					{" ", " ", " ", "y", " ", " ", " ", "d"},
					{"w", " ", "v", " ", "w", " ", "r", " "},
					{" ", "a", " ", " ", " ", "o", " ", " "},
				},
			},
		},
		{
			name: "should create wavy word with 1 line if waviness=0",
			args: args{
				word:     "wavyword",
				waviness: 0,
			},
			want: &WavyWord{
				Word: "wavyword",
				Lines: [][]string{
					{"w", "a", "v", "y", "w", "o", "r", "d"},
				},
			},
		},
		{
			name: "should return error if word is empty",
			args: args{
				word:     "",
				waviness: 2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return error if waviness is negative",
			args: args{
				word:     "wavyword",
				waviness: -1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.word, tt.args.waviness)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
