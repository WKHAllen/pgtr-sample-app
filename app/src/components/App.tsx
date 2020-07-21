import React from 'react';
import { BrowserRouter as Router, Route, Link, Switch } from 'react-router-dom';
import '../css/App.css';
import Test1 from './Test1';
import Test2 from './Test2';
import Test3 from './Test3';

export default class App extends React.Component {
	render() {
		return (
			<Router>
				<div className="App">
					<ul className="App-header">
						<li>
							<Link to="/">Home</Link>
						</li>
						<li>
							<Link to="/test1">Test 1</Link>
						</li>
						<li>
							<Link to="/test2">Test 2</Link>
						</li>
						<li>
							<Link to="/test3">Test 3</Link>
						</li>
					</ul>
					<Switch>
						<Route exact path="/test1" component={Test1}></Route>
						<Route exact path="/test2" component={Test2}></Route>
						<Route exact path="/test3" component={Test3}></Route>
					</Switch>
				</div>
			</Router>
		);
	}
}
