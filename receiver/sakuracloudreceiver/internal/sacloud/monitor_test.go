package sacloud

import (
	"testing"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/stretchr/testify/require"
)

func TestMonitor_monitorCPUTimeValue(t *testing.T) {
	cases := []struct {
		name   string
		in     []*iaas.MonitorCPUTimeValue
		expect *iaas.MonitorCPUTimeValue
	}{
		{
			name:   "input is nil",
			in:     nil,
			expect: nil,
		},
		{
			name: "input has only 1 value",
			in: []*iaas.MonitorCPUTimeValue{
				{
					Time:    time.Now(),
					CPUTime: 1.0,
				},
			},
			expect: nil,
		},
		{
			name: "second last value is used: with 2 values",
			in: []*iaas.MonitorCPUTimeValue{
				{
					Time:    time.Unix(1, 0),
					CPUTime: 1.0,
				},
				{
					Time:    time.Unix(2, 0),
					CPUTime: 2.0,
				},
			},
			expect: &iaas.MonitorCPUTimeValue{
				Time:    time.Unix(1, 0),
				CPUTime: 1.0,
			},
		},
		{
			name: "second last value is used: with over 2 values",
			in: []*iaas.MonitorCPUTimeValue{
				{
					Time:    time.Unix(0, 0),
					CPUTime: 0.0,
				},
				{
					Time:    time.Unix(1, 0),
					CPUTime: 1.0,
				},
				{
					Time:    time.Unix(2, 0),
					CPUTime: 2.0,
				},
			},
			expect: &iaas.MonitorCPUTimeValue{
				Time:    time.Unix(1, 0),
				CPUTime: 1.0,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := monitorCPUTimeValue(tc.in)
			require.Equal(t, tc.expect, actual, tc.name)
		})
	}
}
