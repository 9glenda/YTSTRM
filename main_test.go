package main

import "testing"

func TestExtractVideoID(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "valid ID",
			arg:  "DQeg0ZalP70",
			want: "DQeg0ZalP70",
		},
		{
			name: "valid HTTPS URL",
			arg:  "https://www.youtube.com/watch?v=DQeg0ZalP70",
			want: "DQeg0ZalP70",
		},
		{
			name: "valid HTTP URL without www",
			arg:  "http://youtube.com/watch?v=DQeg0ZalP70",
			want: "DQeg0ZalP70",
		},
		{
			name: "valid URL without protocol",
			arg:  "www.youtube.com/watch?v=DQeg0ZalP70",
			want: "DQeg0ZalP70",
		},
		{
			name: "valid HTTP URL",
			arg:  "http://www.youtube.com/watch?v=DQeg0ZalP70",
			want: "DQeg0ZalP70",
		},
		{
			name: "valid URL without protocol",
			arg:  "www.youtube.com/watch?v=DQeg0ZalP70",
			want: "DQeg0ZalP70",
		},
		{
			name: "URL with extra parameters",
			arg:  "https://www.youtube.com/watch?v=DQeg0ZalP70&t=123",
			want: "DQeg0ZalP70",
		},
		{
			name: "URL with playlist parameter",
			arg:  "https://www.youtube.com/watch?v=DQeg0ZalP70&list=PLG4M6LK1K9Rg-M6UvxcF43UAyNcEOhMFd",
			want: "DQeg0ZalP70",
		},
		{
			name: "invalid ID (too short)",
			arg:  "DQeg0ZalP7",
			want: "",
		},
		{
			name: "invalid ID (too long)",
			arg:  "DQeg0ZalP7012345678901234567890123456789012345678901234567890123456789",
			want: "",
		},
		{
			name: "invalid URL (not a YouTube URL)",
			arg:  "https://www.example.com/watch?v=DQeg0ZalP70",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractVideoID(tt.arg)
			if got != tt.want {
				t.Errorf("extractVideoID(%q) = %q; want %q", tt.arg, got, tt.want)
			}
		})
	}
}
