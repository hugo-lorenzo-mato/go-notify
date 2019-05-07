package cards

type DefaultMessageCard struct {
	Type            string            `json:"@type"`
	Context         string            `json:"@context"`
	Summary         string            `json:"summary"`
	ThemeColor      string            `json:"themeColor"`
	Title           string            `json:"title"`
	Sections        []Sections        `json:"sections"`
	PotentialAction []PotentialAction `json:"potentialAction"`
}
type SectionsFacts struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Sections struct {
	ActivitySubtitle string          `json:"activitySubtitle"`
	Facts            []SectionsFacts `json:"facts"`
	Text             string          `json:"text"`
}
type Inputs struct {
	Type        string `json:"@type"`
	ID          string `json:"id"`
	Title       string `json:"title"`
	IsMultiline bool   `json:"isMultiline"`
}
type PotentialActions struct {
	Type   string `json:"@type"`
	Name   string `json:"name"`
	Target string `json:"target"`
}
type Targets struct {
	Os  string `json:"os"`
	URI string `json:"uri"`
}
type PotentialAction struct {
	Type    string             `json:"@type"`
	Name    string             `json:"name"`
	Inputs  []Inputs           `json:"inputs,omitempty"`
	Actions []PotentialActions `json:"actions,omitempty"`
	Targets []Targets          `json:"targets,omitempty"`
}

var DefaultMessageCardExample = `
{
	"@type": "MessageCard",
	"@context": "https://schema.org/extensions",
	"summary": "GO-NOTiFY - new message",
	"themeColor": "0078D7",
	"title": "TITLE",
	"sections": [
		{
			"activitySubtitle": "DATE",
			"facts": [
				{
					"name":  "Plan:",
					"value": "PLAN"
				}
			],
			"text": "messageText"
		}
	],
	"potentialAction": [
		{
			"@type": "OpenUri",
			"name": "Ver",
			"targets": [
				{
					"os": "default",
					"uri": "URI"
				}
			]
		}
	]
}
`
