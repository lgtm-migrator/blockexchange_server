import Start from './components/Start.js';
import Search from './components/Search.js';
import UserSchemaList from './components/UserSchemaList.js';
import SchemaDetail from './components/SchemaDetail.js';
import UserList from './components/UserList.js';
import Login from './components/Login.js';
import Register from './components/Register.js';
import Mod from './components/Mod.js';

export default {
  "/": Start,
  "/users": UserList,
  "/login": Login,
  "/register": Register,
  "/search": Search,
  "/mod": Mod,
  "/schema/:username": UserSchemaList,
  "/schema/:username/:schemaname": SchemaDetail
};
