package formatters

import "github.com/tidwall/gjson"

func FormatMetaData(metadataResponse string) string {
	wikipedia := gjson.Get(metadataResponse, "stops.1.wikipedia.abstract").String()
	distance := gjson.Get(metadataResponse, "distance").String() + "km"
	return "Distance from Norway: " + distance + "\n\n" + "Wikiedia description" + "\n" + wikipedia

}
