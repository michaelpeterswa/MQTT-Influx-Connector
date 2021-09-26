package main

import (
	"fmt"
	"log"
	"strings"
)

func buildTopic(s SubTopic) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", s.Type, s.Location, s.Room, s.Name, s.Field)
}

func getSubTopicFromString(s string) SubTopic {
	var st SubTopic
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
