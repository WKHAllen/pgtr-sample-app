import React from 'react';
import { Link } from 'react-router-dom';

export default class Home extends React.Component {
	render() {
		return (
			<ul className="App-header">
				<li>
					<Link to="/people">People</Link>
				</li>
				<li>
					<Link to="/quotes">Quotes</Link>
				</li>
			</ul>
		);
	}
}
