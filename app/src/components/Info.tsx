import React from 'react';
import '../css/Info.css';

export default class Info extends React.Component {
    render() {
        return (
            <div className="Info">
                This app shows off the capabilities of the PGTR stack.
                It can be used as a template for more practical applications.
                To use it, <a href="https://github.com/WKHAllen/pgtr-sample-app" target="_blank" rel="noopener noreferrer">see the source</a>.
                Details on the stack and instructions on how to build upon the template application are provided.
                Click the buttons above to see the app in use.
                You can always navigate back home by clicking on the logo or text in the header.
            </div>
        );
    }
}
