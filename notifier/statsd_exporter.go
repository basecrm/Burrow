package notifier

import (
	"github.com/mailgun/metrics"
	"strconv"
	"strings"
)

type StatsdExporter struct {
	Client       metrics.Client
	LagThreshold int64
}

func (exporter *StatsdExporter) Export(msg Message) error {
	for _, partition := range msg.Partitions {
		if partition.Lag > exporter.LagThreshold {
			topic := strings.Replace(partition.Topic, ".", "_", -1)
			m := exporter.Client.Metric("consumer_lag", msg.Cluster, msg.Group, topic, strconv.Itoa(int(partition.Partition)))
			exporter.Client.Gauge(m, int64(partition.Lag), 1)
		}

	}
	return nil
}
