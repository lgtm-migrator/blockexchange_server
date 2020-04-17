const app = require("./app");
const migrate = require("./migrate");

// user mgmt
require("./api/user");
require("./api/register");
require("./api/token");

// stats / info
require("./api/info");

// search
require("./api/searchschema");

// down / upload
require("./api/schema_get");
require("./api/schema_update");
require("./api/schema_create");
require("./api/schema_delete");
require("./api/schemamods");
require("./api/schemapart");

migrate().then(() => {
  app.listen(8080, () => console.log('Listening on http://127.0.0.1:8080'));
})
.catch(() => {
  process.exit(-1);
});
