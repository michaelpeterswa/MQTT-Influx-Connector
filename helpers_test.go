package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTopic(t *testing.T) {
	tests := []struct {
		name  string
		st    SubTopic
		topic string
	}{
		{
			name: "testing buildTopic() basic",
			st: SubTopic{
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
			assert.Equal(t, tc.topic, buildTopic(tc.st))
		})
	}
}

func TestGetSubTopicFromString(t *testing.T) {
	tests := []struct {
		name string
		str  string
		st   SubTopic
	}{
		{
			name: "testing buildTopic() basic",
			str:  "a/b/c/d/e",
			st: SubTopic{
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
			assert.Equal(t, tc.st, getSubTopicFromString(tc.str))
		})
	}
}
