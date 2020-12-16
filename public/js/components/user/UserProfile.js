import store from '../../store/login.js';

export default {
	data: function(){
		return {
			store: store
		};
	},
	template: /*html*/`
		<div class="row">
			<div class="col-md-6">
				<div class="card">
					<div class="card-header">
						Update profile
					</div>
					<div class="card-body">
						<input type="text"
							class="form-control"
							placeholder="Username"
							v-model="store.claims.username"
						/>
						<input type="text"
							class="form-control"
							placeholder="E-Mail"
							v-model="store.claims.mail"
						/>
						<a class="btn btn-primary btn-block">
							<i class="fa fa-save"/> Update profile
						</a>
					</div>
				</div>
			</div>
			<div class="col-md-6">
				<div class="card">
				  <div class="card-header">
				    Profile info
				  </div>
				  <div class="card-body">
						<ul>
							<li>
								ID: <b>{{ store.claims.user_id }}</b>
							</li>
							<li>
								Name: <b>{{ store.username }}</b>
							</li>
							<li>
								Role: <span class="badge badge-success">{{ store.claims.role }}</span>
							</li>
							<li>
								Type: <span class="badge badge-success">{{ store.claims.type }}</span>
							</li>
						</ul>
				  </div>
				</div>
			</div>
		</div>
	`
};