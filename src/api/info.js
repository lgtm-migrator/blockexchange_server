const app = require("../app");

app.get('/api/info', function(req, res){
  console.log("GET /api/info");

	res.json({
		api_version_major: 1,
		api_version_minor: 1,
		name: process.env.BLOCKEXCHANGE_NAME || "unknown",
		owner: process.env.BLOCKEXCHANGE_OWNER || "unknown",
	});
});
