package helpers

import (
	"fmt"
	"log"
	"strings"

	"github.com/michaelpeterswa/MQTT-Influx-Connector/internal/structs"
)

func BuildTopic(s structs.SubTopic) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", s.Type, s.Location, s.Room, s.Name, s.Field)
}

func GetSubTopicFromString(s string) structs.SubTopic {
	var st structs.SubTopic
	t := strings.Split(s, "/")
	if len(t) == 5 {
		st.Type = t[0]
		st.Location = t[1]
		st.Room = t[2]
		st.Name = t[3]
		st.Field = t[4]
		return st
	}
	log.Fatal("String was malformed. Couldn't get SubTopic")
	return st
}
