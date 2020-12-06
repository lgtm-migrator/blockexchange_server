import StartPage from './components/pages/StartPage.js';
import LoginPage from './components/pages/LoginPage.js';
import SearchPage from './components/pages/SearchPage.js';
import RegisterPage from './components/pages/RegisterPage.js';

export default [{
  path: "/", component: StartPage
},{
  path: "/login", component: LoginPage
},{
  path: "/search", component: SearchPage
},{
	path: "/register", component: RegisterPage
}];
