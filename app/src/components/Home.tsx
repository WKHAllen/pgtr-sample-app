import React from 'react';
import { Link } from 'react-router-dom';

export default class NotFound extends React.Component {
	render() {
		return (
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
		);
	}
}
