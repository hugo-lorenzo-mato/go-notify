package cards

type DefaultAdaptiveCard struct {
	Type    string    `json:"type,omitempty"`
	Body    []Body    `json:"body,omitempty"`
	Actions []Actions `json:"actions,omitempty"`
	Schema  string    `json:"$schema,omitempty"`
	Version string    `json:"version,omitempty"`
}
type Items struct {
	Type                string `json:"type,omitempty"`
	Weight              string `json:"weight,omitempty"`
	Text                string `json:"text,omitempty"`
	Wrap                bool   `json:"wrap,omitempty"`
	HorizontalAlignment string `json:"horizontalAlignment,omitempty"`
	Spacing             string `json:"spacing,omitempty"`
	IsSubtle            bool   `json:"isSubtle,omitempty"`
}
type Columns struct {
	Type  string  `json:"type,omitempty"`
	Items []Items `json:"items,omitempty"`
	Width string  `json:"width,omitempty"`
}
type Facts struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
}
type Body struct {
	Type    string    `json:"type,omitempty"`
	Size    string    `json:"size,omitempty"`
	Weight  string    `json:"weight,omitempty"`
	Text    string    `json:"text,omitempty"`
	Columns []Columns `json:"columns,omitempty"`
	Wrap    bool      `json:"wrap,omitempty"`
	Facts   []Facts   `json:"facts,omitempty"`
}
type ActionsCardBody struct {
	Type        string `json:"type,omitempty"`
	ID          string `json:"id,omitempty"`
	Placeholder string `json:"placeholder"`
	IsMultiline bool   `json:"isMultiline"`
}
type Actions struct {
	Type  string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
}
type Card struct {
	Type    string            `json:"type,omitempty"`
	Style   string            `json:"style,omitempty"`
	Body    []ActionsCardBody `json:"body,omitempty"`
	Actions []CardActions     `json:"actions,omitempty"`
	Schema  string            `json:"$schema,omitempty"`
}
type CardActions struct {
	Type  string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
	Card  Card   `json:"card,omitempty"`
}

var AdaptiveCardExample = `
{
    "type": "AdaptiveCard",
    "body": [
        {
            "type": "TextBlock",
            "size": "Medium",
            "weight": "Bolder",
            "text": "Publish Adaptive Card schema"
        },
        {
            "type": "ColumnSet",
            "columns": [
                {
                    "type": "Column",
                    "items": [
                        {
                            "type": "TextBlock",
                            "weight": "Bolder",
                            "text": "Matt Hidinger",
                            "wrap": true
                        },
                        {
                            "type": "TextBlock",
                            "horizontalAlignment": "Left",
                            "spacing": "None",
                            "text": "Created {{DATE(2017-02-14T06:08:39Z, SHORT)}}",
                            "isSubtle": true,
                            "wrap": true
                        }
                    ],
                    "width": "stretch"
                }
            ]
        },
        {
            "type": "TextBlock",
            "text": "Now that we have defined the main rules and features of the format, we need to produce a schema and publish it to GitHub. The schema will be the starting point of our reference documentation.",
            "wrap": true
        },
        {
            "type": "FactSet",
            "facts": [
                {
                    "title": "Bamboo plan:",
                    "value": "Adaptive Card"
                },
                {
                    "title": "Jira Key:",
                    "value": "Backlog"
                }
            ]
        }
    ],
    "actions": [
        {
            "type": "Action.ShowCard",
            "title": "Comment",
            "card": {
                "type": "AdaptiveCard",
                "style": "emphasis",
                "body": [
                    {
                        "type": "Input.Text",
                        "id": "comment",
                        "placeholder": "Enter your comment",
                        "isMultiline": true
                    }
                ],
                "actions": [
                    {
                        "type": "Action.Submit",
                        "title": "OK"
                    }
                ],
                "$schema": "http://adaptivecards.io/schemas/adaptive-card.json"
            }
        }
    ],
    "$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
    "version": "1.0"
}
`
