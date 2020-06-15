package formatters

import "github.com/tidwall/gjson"

func FormatAdvisory(advisoryResponse string, countryCode string) string {
	advisory := gjson.Get(advisoryResponse, "data."+countryCode+".advisory")
	score := advisory.Get("score").String()
	message := advisory.Get("message").String()

	return "Advisory score (lower is safer): " + score + "\n" + "Advisory description: " + message
}
