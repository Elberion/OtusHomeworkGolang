package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type countWord struct {
	Word  string
	Count int
}

func Top10(incText string) []string {
	resultSlice := make([]string, 0, 10)
	words := strings.Fields(incText)
	countMap := make(map[string]int, len(words))

	for _, v := range words {
		countMap[v]++
	}
	sortedSlice := make([]countWord, 0, len(countMap))

	for k, v := range countMap {
		sortedSlice = append(sortedSlice, countWord{k, v})
	}
	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i].Count > sortedSlice[j].Count ||
			((sortedSlice[i].Count == sortedSlice[j].Count) && (sortedSlice[i].Word < sortedSlice[j].Word))
	})

	for _, v := range sortedSlice {
		resultSlice = append(resultSlice, v.Word)
	}

	if len(resultSlice) < 10 {
		return resultSlice
	}
	return resultSlice[:10]
}
