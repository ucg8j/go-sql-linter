package fs

import (
	"reflect"
	"testing"
)

func TestTrailingWhitespace(t *testing.T) {
	type args struct {
		lines []string
		lint  bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "2 line lint false",
			args: args{
				lines: []string{
					"IAMSTRING ",
					"thisbestringtwo  ",
				},
				lint: false,
			},
			want: []string{"IAMSTRING", "thisbestringtwo"},
		},
		{
			name: "1 line lint true",
			args: args{
				lines: []string{
					"IAMSTRING ",
				},
				lint: true,
			},
			want: []string{"line 1, issue = Trailing whitespace"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrailingWhitespace(tt.args.lines, tt.args.lint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrailingWhitespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultipleNewLines(t *testing.T) {
	type args struct {
		lines []string
		lint  bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "multiple new lines lint false",
			args: args{
				lines: []string{"ab", "", "", "", "c"},
				lint:  false,
			},
			want: []string{"ab", "", "c"},
		},
		{
			name: "multiple new lines lint true",
			args: args{
				lines: []string{"ab", "", "", "c"},
				lint:  true,
			},
			want: []string{"line 2, issue = Multiple new lines"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultipleNewLines(tt.args.lines, tt.args.lint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultipleNewLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapitaliseKeywords(t *testing.T) {
	type args struct {
		lines []string
		lint  bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "capitalise keyword lint true",
			args: args{
				lines: []string{"select FROM wHeRe"},
				lint:  false,
			},
			want: []string{"SELECT FROM WHERE"},
		},
		{
			name: "capitalise keyword lint true",
			args: args{
				lines: []string{"select"},
				lint:  true,
			},
			want: []string{"line 1, issue = SELECT Keyword not capitalised"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CapitaliseKeywords(tt.args.lines, tt.args.lint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CapitaliseKeywords() = %v, want %v", got, tt.want)
			}
		})
	}
}
