package models

type SearchZincRequest struct {
	SearchType string `json:"search_type"`
	Query      struct {
		Term  string `json:"term"`
		Field string `json:"field"`
	}
	SortFields []string `json:"sort_fields"`
	From       uint     `json:"from"`
	MaxResults uint     `json:"max_results"`
	Source     []string `json:"_source"`
	Highlight  struct {
		Fields map[string]interface{} `json:"fields"`
	}
}

type SearchZincResponse struct {
	Hits struct {
		Total struct {
			Value uint `json:"value"`
		} `json:"total"`
		Hits []EmailHit `json:"hits"`
	} `json:"hits"`
}

type EmailHit struct {
	Source struct {
		Email
	} `json:"_source"`
	Highlight struct {
		Content []string `json:"content"`
	} `json:"highlight"`
}
