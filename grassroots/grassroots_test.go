package grassroots

import (
	"testing"
	"time"
)

func Test_ScheduleTime(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name     string
		args     args
		schedule []MatchScheduleItem
		want     string
	}{
		{
			name: "two day",
			schedule: []MatchScheduleItem{
				{
					MatchDay:      1,
					StartDateTime: time.Date(2023, time.January, 20, 17, 30, 0, 0, time.Local),
				},
				{
					MatchDay:      2,
					StartDateTime: time.Date(2023, time.January, 27, 17, 30, 0, 0, time.Local),
				},
			},
			want: "Fri 20 Jan 2023 (5:30PM), Fri 27 Jan 2023 (5:30PM)",
		},
		{
			name: "test match",
			schedule: []MatchScheduleItem{
				{
					MatchDay:      1,
					StartDateTime: time.Date(2023, time.January, 20, 11, 00, 0, 0, time.Local),
				},
				{
					MatchDay:      2,
					StartDateTime: time.Date(2023, time.January, 21, 11, 00, 0, 0, time.Local),
				},
				{
					MatchDay:      3,
					StartDateTime: time.Date(2023, time.January, 22, 11, 00, 0, 0, time.Local),
				},
				{
					MatchDay:      4,
					StartDateTime: time.Date(2023, time.January, 23, 11, 00, 0, 0, time.Local),
				},
				{
					MatchDay:      5,
					StartDateTime: time.Date(2023, time.January, 24, 11, 00, 0, 0, time.Local),
				},
			},
			want: "Fri 20 Jan 2023 (11:00AM), Sat 21 Jan 2023 (11:00AM), Sun 22 Jan 2023 (11:00AM), Mon 23 Jan 2023 (11:00AM), Tue 24 Jan 2023 (11:00AM)",
		},
		{
			name: "one day",
			schedule: []MatchScheduleItem{
				{
					MatchDay:      1,
					StartDateTime: time.Date(2023, time.January, 20, 17, 30, 0, 0, time.Local),
				},
			},
			want: "Fri 20 Jan 2023 (5:30PM)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ScheduleTime(tt.schedule); got != tt.want {
				t.Errorf("\nScheduleTime() = %q\n          want = %q", got, tt.want)
			}
		})
	}
}
