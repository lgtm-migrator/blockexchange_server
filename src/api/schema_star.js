const bodyParser = require('body-parser');
const jsonParser = bodyParser.json();
const logger = require("../logger");

const app = require("../app");
const userschemastar_dao = require("../dao/userschemastar");

const { MANAGEMENT } = require("../permissions");
const tokenmiddleware = require("../middleware/token");
const permissioncheck = require("../middleware/permissioncheck");

app.get('/api/schema/:id/star', jsonParser, function(req, res){
  logger.debug("GET /api/schema/:id/star", req.params.id);

	userschemastar_dao.find_by_schema_id(req.params.id)
	.then(result => res.json(result || []))
	.catch(() => res.status(500).end());

});

app.put('/api/schema/:id/star', tokenmiddleware, permissioncheck(MANAGEMENT), jsonParser, function(req, res){
  logger.debug("POST /api/schema/:id/star", req.params.id, req.claims.user_id);

	userschemastar_dao.create(req.claims.user_id, req.params.id)
	.then(() => res.end())
	.catch(() => res.status(500).end());

});

app.delete('/api/schema/:id/star', tokenmiddleware, permissioncheck(MANAGEMENT), jsonParser, function(req, res){
  logger.debug("DELETE /api/schema/:id/star", req.params.id, req.claims.user_id);

	userschemastar_dao.remove(req.claims.user_id, req.params.id)
	.then(() => res.end())
	.catch(() => res.status(500).end());

});
