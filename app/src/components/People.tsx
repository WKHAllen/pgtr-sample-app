import React from 'react';
import '../css/People.css';

interface PeopleState {
	person: string
}

export default class People extends React.Component<{}, PeopleState> {
	constructor(props: any) {
		super(props);
		this.state = {
			person: ''
		};
	}

	getJSON(url: string) {
		return fetch(url).then(res => res.json());
	}

	getPersonById(event: React.ChangeEvent<HTMLInputElement>) {
		this.getJSON(`/api/person/id/${event.target.value}`)
			.then(console.log);
	}

	getRandomPerson() {
		this.getJSON('/api/person/random')
			.then(console.log);
	}

	getAllPeople() {
		this.getJSON('/api/people')
			.then(console.log);
	}

	render() {
		return (
			<div className="People">
				<div className="form-group">
					<label htmlFor="person-id">Person ID</label>
					<input id="person-id"
						   className="form-control"
						   type="number"
						   onChange={(event) => this.getPersonById(event)} />
				</div>
				<div className="form-group">
					<button className="btn btn-primary"
							type="button"
							onClick={() => this.getRandomPerson()}>
						Random person
					</button>
				</div>
				<div className="form-group">
					<button className="btn btn-primary"
							type="button"
							onClick={() => this.getAllPeople()}>
						All people
					</button>
				</div>
				<div>{this.state.person}</div>
			</div>
		);
	}
}
