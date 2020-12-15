const bodyParser = require('body-parser');
const jsonParser = bodyParser.json();
const logger = require("../logger");

const events = require("../events");

const app = require("../app");
const schema_dao = require("../dao/schema");
const tokenmiddleware = require("../middleware/token");
const rolecheck = require("../util/rolecheck");
const tokencheck = tokenmiddleware(claims => rolecheck.can_upload(claims.role));

app.post('/api/schema', tokencheck, jsonParser, function(req, res){
  logger.debug("POST /api/schema", req.body);

  schema_dao.create({
    user_id: +req.claims.user_id,
    name: req.body.name,
    description: req.body.description,
    max_x: req.body.max_x,
    max_y: req.body.max_y,
    max_z: req.body.max_z,
    part_length: req.body.part_length,
		license: req.body.license
  })
  .then(inserted_data => res.json(inserted_data))
  .catch(e => {
    console.error(e);
    res.status(500).end();
  });

});



app.post('/api/schema/:id/complete', tokencheck, jsonParser, function(req, res){
  logger.debug("POST /api/schema/id/complete", req.params.id, req.body);

  return schema_dao.get_by_id(req.params.id)
  .then(schema => {
    // check user id in claims
    if (schema.user_id != +req.claims.user_id){
      res.status(401).end();
      return;
    }

    // check if already completed
    if (schema.complete){
      res.status(500).end();
      return;
    }

    return schema_dao.finalize(schema.id)
    .then(() => res.end())
    .then(() => events.emit("new-schema", schema))
    .catch(() => res.status(500).end());
    });
});
