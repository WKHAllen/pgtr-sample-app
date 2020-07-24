import React from 'react';
import { Link } from 'react-router-dom';
import '../css/Home.css';

export default class Home extends React.Component {
	render() {
		return (
			<div className="Home">
				<h1>Navigation</h1>
				<div className="Nav-Group">
					<span><Link to="/people" className="btn btn-primary">People</Link></span>
					<span><Link to="/quotes" className="btn btn-primary">Quotes</Link></span>
				</div>
			</div>
		);
	}
}
