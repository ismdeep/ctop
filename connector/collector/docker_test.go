package collector

import (
	"testing"

	api "github.com/fsouza/go-dockerclient"
)

func TestMemoryCache(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		stats *api.Stats
		want  int64
	}{
		{
			name:  "prefers total inactive file for cgroup v1",
			stats: newMemStats(0, 0, 200, 150, 100),
			want:  200,
		},
		{
			name:  "uses inactive file for cgroup v2",
			stats: newMemStats(0, 0, 0, 150, 100),
			want:  150,
		},
		{
			name:  "falls back to cache for older docker",
			stats: newMemStats(0, 0, 0, 0, 100),
			want:  100,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := memoryCache(tt.stats); got != tt.want {
				t.Fatalf("memoryCache() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestReadMem(t *testing.T) {
	t.Parallel()

	c := &Docker{}
	stats := newMemStats(1024, 2048, 256, 0, 0)

	c.ReadMem(stats)

	if c.MemUsage != 768 {
		t.Fatalf("MemUsage = %d, want 768", c.MemUsage)
	}
	if c.MemLimit != 2048 {
		t.Fatalf("MemLimit = %d, want 2048", c.MemLimit)
	}
	if c.MemPercent != 38 {
		t.Fatalf("MemPercent = %d, want 38", c.MemPercent)
	}
}

func TestReadMemClampsCacheToUsage(t *testing.T) {
	t.Parallel()

	c := &Docker{}
	stats := newMemStats(128, 1024, 256, 0, 0)

	c.ReadMem(stats)

	if c.MemUsage != 0 {
		t.Fatalf("MemUsage = %d, want 0", c.MemUsage)
	}
	if c.MemPercent != 0 {
		t.Fatalf("MemPercent = %d, want 0", c.MemPercent)
	}
}

func newMemStats(usage, limit, totalInactiveFile, inactiveFile, cache uint64) *api.Stats {
	stats := &api.Stats{}
	stats.MemoryStats.Usage = usage
	stats.MemoryStats.Limit = limit
	stats.MemoryStats.Stats.TotalInactiveFile = totalInactiveFile
	stats.MemoryStats.Stats.InactiveFile = inactiveFile
	stats.MemoryStats.Stats.Cache = cache
	return stats
}
