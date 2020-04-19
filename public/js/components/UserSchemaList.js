import SchemaList from './schemalist/SchemaList.js';
import Breadcrumb from './Breadcrumb.js';

import { find_by_username } from '../api/searchschema.js';
import { remove } from '../api/schema.js';

export default class {
  constructor(vnode) {
    this.state = {
      list: [],
      username: vnode.attrs.username
    };
    this.search();
  }

  search(){
    find_by_username(this.state.username)
    .then(l => this.state.list = l);
  }

  removeItem(schema){
    remove(schema)
    .then(() => this.search());
  }

  view(){
    return [
			m(Breadcrumb, {
				links: [{
					name: "Home",
					link: "#!/"
				},{
					name: "User-schemas",
				},{
					name: this.state.username,
					active: true
				}]
			}),
			m(SchemaList, {
	      list: this.state.list,
				hide_user: true,
	      removeItem: schema => this.removeItem(schema)
	    })
		];
  }
}
