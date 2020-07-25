import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import '../css/App.css';
import Home from './Home';
import Header from './Header';
import People from './People';
import Quotes from './Quotes';
import NotFound from './errors/NotFound';

export default class App extends React.Component {
	render() {
		return (
			<Router>
				<div className="App">
					<Header />
					<div className="App-Body">
						<Switch>
							<Route exact path="/" component={Home}></Route>
							<Route exact path="/people" component={People}></Route>
							<Route exact path="/quotes" component={Quotes}></Route>
							<Route exact path="/error/404" component={NotFound}></Route>
						</Switch>
					</div>
				</div>
			</Router>
		);
	}
}
