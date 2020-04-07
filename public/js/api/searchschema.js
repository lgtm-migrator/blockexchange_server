
export const find_by_description = keywords => m.request({
	method: "POST",
	data: {
		keywords: keywords
	},
	url: "api/searchschema"
});

export const find_recent = count => m.request({
	url: `api/searchrecent/${count}`
});