package sacloud

import (
	"sort"
	"time"

	"github.com/sacloud/iaas-api-go"
)

func monitorCondition(end time.Time) *iaas.MonitorCondition {
	end = end.Truncate(time.Second)
	start := end.Add(-time.Hour)
	return &iaas.MonitorCondition{
		Start: start,
		End:   end,
	}
}

func monitorDatabaseValue(values []*iaas.MonitorDatabaseValue) *iaas.MonitorDatabaseValue {
	if len(values) > 1 {
		// Descending
		sort.Slice(values, func(i, j int) bool { return values[i].Time.After(values[j].Time) })
		return values[1]
	}
	return nil
}

func monitorCPUTimeValue(values []*iaas.MonitorCPUTimeValue) *iaas.MonitorCPUTimeValue {
	if len(values) > 1 {
		// Descending
		sort.Slice(values, func(i, j int) bool { return values[i].Time.After(values[j].Time) })
		return values[1]
	}
	return nil
}

func monitorDiskValue(values []*iaas.MonitorDiskValue) *iaas.MonitorDiskValue {
	if len(values) > 1 {
		// Descending
		sort.Slice(values, func(i, j int) bool { return values[i].Time.After(values[j].Time) })
		return values[1]
	}
	return nil
}

func monitorInterfaceValue(values []*iaas.MonitorInterfaceValue) *iaas.MonitorInterfaceValue {
	if len(values) > 1 {
		// Descending
		sort.Slice(values, func(i, j int) bool { return values[i].Time.After(values[j].Time) })
		return values[1]
	}
	return nil
}

func monitorRouterValue(values []*iaas.MonitorRouterValue) *iaas.MonitorRouterValue {
	if len(values) > 1 {
		// Descending
		sort.Slice(values, func(i, j int) bool { return values[i].Time.After(values[j].Time) })
		return values[1]
	}
	return nil
}

func monitorFreeDiskSizeValue(values []*iaas.MonitorFreeDiskSizeValue) *iaas.MonitorFreeDiskSizeValue {
	if len(values) > 1 {
		// Descending
		sort.Slice(values, func(i, j int) bool { return values[i].Time.After(values[j].Time) })
		return values[1]
	}
	return nil
}

func monitorConnectionValue(values []*iaas.MonitorConnectionValue) *iaas.MonitorConnectionValue {
	if len(values) > 2 {
		// Descending
		sort.Slice(values, func(i, j int) bool { return values[i].Time.After(values[j].Time) })
		return values[1]
	}
	return nil
}

func monitorLinkValue(values []*iaas.MonitorLinkValue) *iaas.MonitorLinkValue {
	if len(values) > 2 {
		// Descending
		sort.Slice(values, func(i, j int) bool { return values[i].Time.After(values[j].Time) })
		return values[1]
	}
	return nil
}

func monitorLocalRouterValue(values []*iaas.MonitorLocalRouterValue) *iaas.MonitorLocalRouterValue {
	if len(values) > 1 {
		// Descending
		sort.Slice(values, func(i, j int) bool { return values[i].Time.After(values[j].Time) })
		return values[1]
	}
	return nil
}
