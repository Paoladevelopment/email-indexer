package models

type IndexBody struct {
	IndexName string  `json:"index"`
	Records   []Email `json:"records"`
}

// The API docs of zincsearch indicates this structure to create a new index
type Index struct {
	Name        string     `json:"name"`
	StorageType string     `json:"storage_type"`
	Shards      uint       `json:"shard_num"`
	Mappings    Properties `json:"mappings"`
}

type Properties struct {
	Indexes map[string]IndexField `json:"properties"`
}

type IndexField struct {
	Type          string `json:"type"`
	Index         bool   `json:"index"`
	Store         bool   `json:"store"`
	Sortable      bool   `json:"sortable"`
	Aggregatable  bool   `json:"aggregatable"`
	Highlightable bool   `json:"highlightable"`
	Format        string `json:"format"`
}

func NewEmailIndex() Index {
	return Index{
		Name:        "emails",
		StorageType: "disk",
		Shards:      3,
		Mappings: Properties{
			map[string]IndexField{
				"message_id": {
					Type:  "keyword",
					Index: true,
					Store: false,
				},
				"date": {
					Type:         "text",
					Index:        true,
					Store:        false,
					Sortable:     true,
					Aggregatable: true,
				},
				"from": {
					Type:          "text",
					Index:         true,
					Store:         false,
					Highlightable: true,
				},
				"to": {
					Type:          "text",
					Index:         true,
					Store:         false,
					Highlightable: true,
				},
				"x_to": {
					Type:          "text",
					Index:         true,
					Store:         false,
					Highlightable: true,
				},
				"subject": {
					Type:          "text",
					Index:         true,
					Store:         false,
					Highlightable: true,
				},
				"cc": {
					Type:          "text",
					Index:         true,
					Store:         false,
					Highlightable: true,
				},
				"x_cc": {
					Type:          "text",
					Index:         true,
					Store:         false,
					Highlightable: true,
				},
				"bcc": {
					Type:          "text",
					Index:         true,
					Store:         false,
					Highlightable: true,
				},
				"x_bcc": {
					Type:          "text",
					Index:         true,
					Store:         false,
					Highlightable: true,
				},
				"content": {
					Type:          "text",
					Index:         true,
					Store:         false,
					Highlightable: true,
				},
			},
		},
	}
}
