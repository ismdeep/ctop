package connector

import (
	"testing"
	"time"

	api "github.com/fsouza/go-dockerclient"
)

func TestCalcUptime(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		insp *api.Container
		want string
	}{
		{
			name: "running container returns duration",
			insp: &api.Container{
				State: api.State{
					Running:   true,
					StartedAt: time.Now().Add(-2 * time.Minute),
				},
			},
			want: "2 minutes",
		},
		{
			name: "exited container returns placeholder",
			insp: &api.Container{
				State: api.State{
					StartedAt:  time.Now().Add(-10 * time.Minute),
					FinishedAt: time.Now().Add(-5 * time.Minute),
				},
			},
			want: "-",
		},
		{
			name: "created container returns placeholder",
			insp: &api.Container{
				State: api.State{},
			},
			want: "-",
		},
		{
			name: "future start time returns placeholder",
			insp: &api.Container{
				State: api.State{
					Running:   true,
					StartedAt: time.Now().Add(2 * time.Minute),
				},
			},
			want: "-",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := calcUptime(tt.insp); got != tt.want {
				t.Fatalf("calcUptime() = %q, want %q", got, tt.want)
			}
		})
	}
}
