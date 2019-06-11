package elasticstore

func mapping(indexName string) string {

	return `{
	"settings": {
		"analysis": {
			"analyzer": {
				"my_analyzer": {
					"type": "custom",
					"filter": [
						"lowercase"
					],
					"tokenizer": "whitespace"
				}
			}
		}
	},

	"mappings":{
		"properties":{
			"CarID":{
				"type": "text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"UserID":{
				"type":"text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"DateStart":{
				"type":"text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"DateReturned":{
				"type":"text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"Rate":{
				"type":"float",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"TotalDue":{
				"type":"float",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"AmtPaid":{
				"type":"float",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"AccountCreationDate":{
				"type":"text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			}
		}
	}	
}
`
}
