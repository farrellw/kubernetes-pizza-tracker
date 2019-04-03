import React, { Component } from "react";
import Countdown from "./Countdown";
import "./App.css";

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <Countdown />
        </header>
      </div>
    );
  }
}

export default App;
