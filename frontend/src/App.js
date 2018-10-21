import React, { Component } from 'react';
import './App.css';

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            name: '',
        };
    };

    send() {
        fetch('http://192.168.99.100:31961', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({name: this.textField.value})
        })
        .then(response => {
            return response.json()
        })
        .then(data => {
            this.setState({name: data.name})
        })
        .catch(console.error)
    }

    render() {
        const responseComponent = this.state.name ?
            <p>{"Hello " + this.state.name + "!"}</p> :
            null;
        return (
            <div className="App">
                <header className="App-header">
                    <input ref={ref => this.textField = ref} />
                    <button onClick={this.send.bind(this)}>Отправить</button>
                    {responseComponent}
                </header>
            </div>
        );
    }
}

export default App;
