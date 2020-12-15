const bodyParser = require('body-parser');
const jsonParser = bodyParser.json({limit: '20mb'});
const logger = require("../logger");

const app = require("../app");
const schemapart_dao = require("../dao/schemapart");
const schema_dao = require("../dao/schema");
const serializer = require("../util/serializer");

const tokenmiddleware = require("../middleware/token");
const rolecheck = require("../util/rolecheck");
const tokencheck = tokenmiddleware(claims => rolecheck.can_upload(claims.role));

app.post('/api/schemapart', tokencheck, jsonParser, function(req, res){
  logger.debug("POST /api/schemapart", req.body.schema_id, req.body.offset_x, req.body.offset_y, req.body.offset_z);

  return schema_dao.get_by_id(req.body.schema_id)
  .then(schema => {
    // check user id in claims
    if (schema.user_id != +req.claims.user_id){
      res.status(401).end();
      return;
    }

    if (schema.completed) {
      res.status(500).end();
      return;
    }

    const serialized_data = serializer.serialize(req.body.data);

    schemapart_dao.create({
      schema_id: schema.id,
      offset_x: req.body.offset_x,
      offset_y: req.body.offset_y,
      offset_z: req.body.offset_z,
      data: serialized_data.data,
      metadata: serialized_data.metadata
    })
    .then(id_obj => res.json(id_obj))
    .catch(() => res.status(500).end());
  });
});


// curl 127.0.0.1:8080/api/schemapart/1/0/0/0
app.get('/api/schemapart/:schema_id/:offset_x/:offset_y/:offset_z', function(req, res){
  logger.debug("GET /api/schemapart", req.params);

  schemapart_dao.get_by_id_and_offset(
    req.params.schema_id,
    req.params.offset_x,
    req.params.offset_y,
    req.params.offset_z
  )
  .then(schemapart => {
    if (schemapart) {
      const data = serializer.deserialize(schemapart);

      res.json({
        schema_id: req.params.schema_id,
        offset_x: req.params.offset_x,
        offset_y: req.params.offset_y,
        offset_z: req.params.offset_z,
        data: {
          node_ids: data.node_ids,
          param1: data.param1,
          param2: data.param2,
          metadata: data.metadata,
          size: data.size,
					node_mapping: data.node_mapping
        }
      });
    } else
      res.status(404).end();
  })
  .catch(e => {
		console.error(e);
		res.status(500).end();
	});
});
