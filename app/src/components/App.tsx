import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import '../css/App.css';
import NotFound from './errors/NotFound';
import Home from './Home';
import Test1 from './Test1';
import Test2 from './Test2';
import Test3 from './Test3';

export default class App extends React.Component {
	render() {
		return (
			<Router>
				<div className="App">
					<Switch>
						<Route exact path="/" component={Home}></Route>
						<Route exact path="/test1" component={Test1}></Route>
						<Route exact path="/test2" component={Test2}></Route>
						<Route exact path="/test3" component={Test3}></Route>
						<Route exact path="/error/404" component={NotFound}></Route>
					</Switch>
				</div>
			</Router>
		);
	}
}
