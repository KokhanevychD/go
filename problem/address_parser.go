package problem

import (
	"sort"
	"strings"
)

// ByState func to sort str input by states and names
// Given a string with friends to visit in different states:
//    ad3="John Daggett, 341 King Road, Plymouth MA
//    Alice Ford, 22 East Broadway, Richmond VA
//    Sal Carpenter, 73 6th Street, Boston MA"
// results:
//    Massachusetts
//    .....^John Daggett 341 King Road Plymouth Massachusetts
//    .....^Sal Carpenter 73 6th Street Boston Massachusetts
//    ^Virginia
//    .....^Alice Ford 22 East Broadway Richmond Virginia
func ByState(str string) string {
	var resStr string
	var results map[string][]string
	results = make(map[string][]string)
	var states = map[string]string{
		"AZ": "Arizona",
		"CA": "California",
		"ID": "Idaho",
		"IN": "Indiana",
		"MA": "Massachusetts",
		"OK": "Oklahoma",
		"PA": "Pennsylvania",
		"VA": "Virginia",
	}

	data := strings.Split(str, "\n")
	for _, value := range data {
		for key, stateValue := range states {
			if strings.Contains(value, key) {
				value = strings.Replace(value, key, stateValue, 1)
				value = strings.ReplaceAll(value, ",", "")
				results[states[key]] = append(results[states[key]], value)
				break
			}
		}
	}

	var keys []string
	for key := range results {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		if resStr != "" {
			resStr += "\n "
		}
		resStr += key
		valueStr := results[key]
		sort.Strings(valueStr)
		for _, str := range valueStr {
			resStr += "\n..... " + str
		}
	}
	return resStr
}
