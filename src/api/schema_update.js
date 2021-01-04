const bodyParser = require('body-parser');
const jsonParser = bodyParser.json();
const logger = require("../logger");

const app = require("../app");
const schema_dao = require("../dao/schema");

const { MANAGEMENT } = require("../permissions");
const tokenmiddleware = require("../middleware/token");
const permissioncheck = require("../middleware/permissioncheck");

app.put("/api/schema/:id", tokenmiddleware, permissioncheck(MANAGEMENT), jsonParser, function(req, res){
  logger.debug("PUT /api/schema/:id", req.params.id, req.body);
  const new_schema = req.body;

  return schema_dao.get_by_id(new_schema.id)
  .then(old_schema => {
    if (old_schema.user_id != req.claims.user_id){
      res.status(401).end();
      return;
    }

    // update schema
    return schema_dao.update(new_schema)
    .then(s => res.json(s));
  });
});
