package config

import (
	"doppler/iprange"
	"errors"
	"github.com/cloudfoundry/gosteno"
	"github.com/cloudfoundry/loggregatorlib/cfcomponent"
	"time"
)

const HeartbeatInterval = 10 * time.Second

type Config struct {
	cfcomponent.Config
	EtcdUrls                      []string
	EtcdMaxConcurrentRequests     int
	Index                         uint
	DropsondeIncomingMessagesPort uint32
	OutgoingPort                  uint32
	LogFilePath                   string
	MaxRetainedLogMessages        uint32
	WSMessageBufferSize           uint
	SharedSecret                  string
	SkipCertVerify                bool
	BlackListIps                  []iprange.IPRange
	JobName                       string
	Zone                          string
	ContainerMetricTTLSeconds     int
	SinkInactivityTimeoutSeconds  int
}

func (c *Config) Validate(logger *gosteno.Logger) (err error) {
	if c.MaxRetainedLogMessages == 0 {
		return errors.New("Need max number of log messages to retain per application")
	}

	if c.BlackListIps != nil {
		err = iprange.ValidateIpAddresses(c.BlackListIps)
		if err != nil {
			return err
		}
	}

	err = c.Config.Validate(logger)
	return
}
