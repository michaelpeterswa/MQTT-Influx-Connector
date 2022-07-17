package helpers_test

import (
	"testing"

	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/helpers"
	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/structs"
	"github.com/stretchr/testify/assert"
)

func TestBuildTopic(t *testing.T) {
	tests := []struct {
		name  string
		st    structs.SubTopic
		topic string
	}{
		{
			name: "testing buildTopic() basic",
			st: structs.SubTopic{
				Type:     "a",
				Location: "b",
				Room:     "c",
				Name:     "d",
				Field:    "e",
			},
			topic: "a/b/c/d/e",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.topic, helpers.BuildTopic(tc.st))
		})
	}
}

func TestGetSubTopicFromString(t *testing.T) {
	tests := []struct {
		name string
		str  string
		st   structs.SubTopic
	}{
		{
			name: "testing buildTopic() basic",
			str:  "a/b/c/d/e",
			st: structs.SubTopic{
				Type:     "a",
				Location: "b",
				Room:     "c",
				Name:     "d",
				Field:    "e",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.st, helpers.GetSubTopicFromString(tc.str))
		})
	}
}
