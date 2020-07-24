import React from 'react';
import '../css/Quotes.css';

interface QuotesState {
	quote: string
}

export default class Quotes extends React.Component<{}, QuotesState> {
	constructor(props: any) {
		super(props);
		this.state = {
			quote: ''
		};
	}

	getJSON(url: string): Promise<Response> {
		return fetch(url).then(res => res.json());
	}

	getQuoteById(event: React.ChangeEvent<HTMLInputElement>) {
		this.getJSON(`/api/quote/id/${event.target.value}`)
			.then(console.log);
	}

	getRandomQuote() {
		this.getJSON('/api/quote/random')
			.then(console.log);
	}

	getAllQuotes() {
		this.getJSON('/api/quotes')
			.then(console.log);
	}

	render() {
		return (
			<div className="Quotes">
				<div className="form-group">
					<label htmlFor="quote-id">Quote ID</label>
					<input id="quote-id"
						   className="form-control"
						   type="number"
						   onChange={(event) => this.getQuoteById(event)} />
				</div>
				<div className="form-group">
					<button className="btn btn-primary"
							type="button"
							onClick={() => this.getRandomQuote()}>
						Random quote
					</button>
				</div>
				<div className="form-group">
					<button className="btn btn-primary"
							type="button"
							onClick={() => this.getAllQuotes()}>
						All quotes
					</button>
				</div>
				<div>{this.state.quote}</div>
			</div>
		);
	}
}
