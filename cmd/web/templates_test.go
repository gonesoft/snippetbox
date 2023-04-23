package main

import (
	"testing"
	"time"
)

func TestReadableDate(t *testing.T) {
	tm := time.Date(2020, 12, 17, 10, 0, 0, 0, time.UTC)
	hd := readableDate(tm)

	if hd != "17 Dec 2020 at 10:00" {
		t.Errorf("want %q; got %q", "17 Dec 2020 at 10:00", hd)
	}
}

func Test_readableDate(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "UTC",
			args: args{
				t: time.Date(2020, 12, 17, 10, 0, 0, 0, time.UTC),
			},
			want: "17 Dec 2020 at 10:00",
		},
		{
			name: "Empty",
			args: args{
				t: time.Time{},
			},
			want: "",
		},
		{
			name: "CET",
			args: args{
				t: time.Date(2020, 12, 17, 10, 0, 0, 0, time.FixedZone("CET", 1*60*60)),
			},
			want: "17 Dec 2020 at 10:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readableDate(tt.args.t); got != tt.want {
				t.Errorf("readableDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
