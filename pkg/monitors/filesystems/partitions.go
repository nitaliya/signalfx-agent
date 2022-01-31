//go:build !windows
// +build !windows

package filesystems

import gopsutil "github.com/shirou/gopsutil/disk"

func getPartitions(all bool) ([]gopsutil.PartitionStat, error) {
	return gopsutil.Partitions(all)
}
