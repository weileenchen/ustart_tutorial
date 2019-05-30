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
			"FirstName":{
				"type": "text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"LastName":{
				"type":"text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"DOB":{
				"type":"text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"BillingInformation": {
				"type":"nested", 
				"properties":{
					"StreetAddress":{
						"type":"text",
						"analyzer":"my_analyzer",
						"fields":{
							"raw":{
								"type":"keyword"
							}
						}
					},
					"City":{
						"type":"text",
						"analyzer":"my_analyzer",
						"fields":{
							"raw":{
								"type":"keyword"
							}
						}
					},
					"State":{
						"type":"text",
						"analyzer":"my_analyzer",
						"fields":{
							"raw":{
								"type":"keyword"
							}
						}
					},
					"Zip":{
						"type":"text",
						"analyzer":"my_analyzer",
						"fields":{
							"raw":{
								"type":"keyword"
							}
						}
					},
					"CardNumber":{
						"type":"text",
						"analyzer":"my_analyzer",
						"fields":{
							"raw":{
								"type":"keyword"
							}
						}
					},
					"CV":{
						"type":"text",
						"analyzer":"my_analyzer",
						"fields":{
							"raw":{
								"type":"keyword"
							}
						}
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
