// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testDataSet int

const (
	testDataSetDefault testDataSet = iota
	testDataSetAll
	testDataSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name        string
		metricsSet  testDataSet
		resAttrsSet testDataSet
		expectEmpty bool
	}{
		{
			name: "default",
		},
		{
			name:        "all_set",
			metricsSet:  testDataSetAll,
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "none_set",
			metricsSet:  testDataSetNone,
			resAttrsSet: testDataSetNone,
			expectEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, tt.name), settings, WithStartTime(start))

			expectedWarnings := 0

			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			allMetricsCount++
			mb.RecordSakuracloudServerCPUTimeDataPoint(ts, 1, "sakuracloud.server.id-val", "sakuracloud.server.name-val", "sakuracloud.server.zone-val")

			allMetricsCount++
			mb.RecordSakuracloudServerDiskReadDataPoint(ts, 1, "sakuracloud.server.id-val", "sakuracloud.server.name-val", "sakuracloud.server.zone-val", "sakuracloud.server.disk_id-val", 29)

			allMetricsCount++
			mb.RecordSakuracloudServerDiskWriteDataPoint(ts, 1, "sakuracloud.server.id-val", "sakuracloud.server.name-val", "sakuracloud.server.zone-val", "sakuracloud.server.disk_id-val", 29)

			allMetricsCount++
			mb.RecordSakuracloudServerNetworkInterfaceReceiveDataPoint(ts, 1, "sakuracloud.server.id-val", "sakuracloud.server.name-val", "sakuracloud.server.zone-val", "sakuracloud.server.interface_id-val", 34)

			allMetricsCount++
			mb.RecordSakuracloudServerNetworkInterfaceSendDataPoint(ts, 1, "sakuracloud.server.id-val", "sakuracloud.server.name-val", "sakuracloud.server.zone-val", "sakuracloud.server.interface_id-val", 34)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSakuracloudServerUpDataPoint(ts, 1, "sakuracloud.server.id-val", "sakuracloud.server.name-val", "sakuracloud.server.zone-val")

			res := pcommon.NewResource()
			metrics := mb.Emit(WithResource(res))

			if tt.expectEmpty {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if tt.metricsSet == testDataSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if tt.metricsSet == testDataSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "sakuracloud.server.cpu_time":
					assert.False(t, validatedMetrics["sakuracloud.server.cpu_time"], "Found a duplicate in the metrics slice: sakuracloud.server.cpu_time")
					validatedMetrics["sakuracloud.server.cpu_time"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "CPU usage time per core in milliseconds. Values range from 0 to 1000", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("sakuracloud.server.id")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.name")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.zone")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.zone-val", attrVal.Str())
				case "sakuracloud.server.disk.read":
					assert.False(t, validatedMetrics["sakuracloud.server.disk.read"], "Found a duplicate in the metrics slice: sakuracloud.server.disk.read")
					validatedMetrics["sakuracloud.server.disk.read"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Disk read throughput per server", ms.At(i).Description())
					assert.Equal(t, "bps", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("sakuracloud.server.id")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.name")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.zone")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.zone-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.disk_id")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.disk_id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.disk_index")
					assert.True(t, ok)
					assert.EqualValues(t, 29, attrVal.Int())
				case "sakuracloud.server.disk.write":
					assert.False(t, validatedMetrics["sakuracloud.server.disk.write"], "Found a duplicate in the metrics slice: sakuracloud.server.disk.write")
					validatedMetrics["sakuracloud.server.disk.write"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Disk write throughput per server", ms.At(i).Description())
					assert.Equal(t, "bps", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("sakuracloud.server.id")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.name")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.zone")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.zone-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.disk_id")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.disk_id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.disk_index")
					assert.True(t, ok)
					assert.EqualValues(t, 29, attrVal.Int())
				case "sakuracloud.server.network_interface.receive":
					assert.False(t, validatedMetrics["sakuracloud.server.network_interface.receive"], "Found a duplicate in the metrics slice: sakuracloud.server.network_interface.receive")
					validatedMetrics["sakuracloud.server.network_interface.receive"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Network interface incoming traffic per NIC", ms.At(i).Description())
					assert.Equal(t, "bps", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("sakuracloud.server.id")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.name")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.zone")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.zone-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.interface_id")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.interface_id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.interface_index")
					assert.True(t, ok)
					assert.EqualValues(t, 34, attrVal.Int())
				case "sakuracloud.server.network_interface.send":
					assert.False(t, validatedMetrics["sakuracloud.server.network_interface.send"], "Found a duplicate in the metrics slice: sakuracloud.server.network_interface.send")
					validatedMetrics["sakuracloud.server.network_interface.send"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Network interface outgoing traffic per NIC", ms.At(i).Description())
					assert.Equal(t, "bps", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("sakuracloud.server.id")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.name")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.zone")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.zone-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.interface_id")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.interface_id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.interface_index")
					assert.True(t, ok)
					assert.EqualValues(t, 34, attrVal.Int())
				case "sakuracloud.server.up":
					assert.False(t, validatedMetrics["sakuracloud.server.up"], "Found a duplicate in the metrics slice: sakuracloud.server.up")
					validatedMetrics["sakuracloud.server.up"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Indicates whether a server is up (1) or down (0)", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("sakuracloud.server.id")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.name")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("sakuracloud.server.zone")
					assert.True(t, ok)
					assert.EqualValues(t, "sakuracloud.server.zone-val", attrVal.Str())
				}
			}
		})
	}
}
