package notifier

import (
	"github.com/linkedin/Burrow/protocol"
	"github.com/mailgun/metrics"
	"strconv"
	"strings"
)

type StatsdReporter struct {
	Client       metrics.Client
	LagThreshold int64
}

func (reporter *StatsdReporter) Notify(msg Message) error {
	for _, partition := range msg.Partitions {
		if partition.Lag > reporter.LagThreshold {
			topic := strings.Replace(partition.Topic, ".", "_", -1)
			m := reporter.Client.Metric("consumer_lag", msg.Cluster, msg.Group, topic, strconv.Itoa(int(partition.Partition)))
			reporter.Client.Gauge(m, int64(partition.Lag), 1)
		}

	}
	return nil
}

func (reporter *StatsdReporter) Ignore(msg Message) bool {
	//always report lag
	return msg.Status <= protocol.StatusNotFound
}

func (reporter *StatsdReporter) NotifierName() string {
	return "statsd-report"
}
