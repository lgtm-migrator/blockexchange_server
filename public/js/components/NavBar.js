import LoginStatus from './login/LoginStatus.js';
import store from '../store/login.js';

export default {
	components: {
		"login-status": LoginStatus
	},
	data: function(){
		return {
			store: store
		};
	},
	template: /*html*/`
		<nav class="navbar navbar-expand-lg navbar-dark bg-dark">
			<router-link to="/" class="navbar-brand">Block exchange</router-link>
			<ul class="navbar-nav mr-auto">
				<li class="nav-item">
					<router-link to="/" class="nav-link">
						<i class="fa fa-question"></i> {{ $t("nav.about") }}
					</router-link>
				</li>
				<li class="nav-item">
					<router-link to="/login" class="nav-link">
						<i class="fa fa-sign-in"></i> {{ $t("nav.login") }}
					</router-link>
				</li>
				<li class="nav-item">
					<router-link to="/mod" class="nav-link">
						<i class="fa fa-download"></i> {{ $t("nav.mod") }}
					</router-link>
				</li>
				<li class="nav-item">
					<router-link to="/users" class="nav-link">
						<i class="fa fa-users"></i> {{ $t("nav.users" )}}
					</router-link>
				</li>
				<li class="nav-item">
					<router-link to="/search" class="nav-link">
						<i class="fa fa-search"></i> {{ $t("nav.search") }}
					</router-link>
				</li>
				<li class="nav-item" v-if="store.loggedIn">
					<router-link to="/collections" class="nav-link">
						<i class="fa fa-list"></i> {{ $t("nav.collections") }}
					</router-link>
				</li>
			</ul>
			<form class="form-inline my-2 my-lg-0">
				<login-status/>
			</form>
	</nav>
	`
};
