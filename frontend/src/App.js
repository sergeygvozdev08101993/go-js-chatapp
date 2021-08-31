import React, { Component } from "react";
import "./App.css";
import Header from "./components/Header/Header"
import { connect, sendMsg } from "./api";

class App extends Component {

  constructor(props) {
    super(props);
    connect();
  }

  send() {
    console.log("hello");
    sendMsg("hello");
  }

  render() {
    return (
        <div className="App">
          <Header />
          <button onClick={this.send}>Click</button>
        </div>
    )
  }
}

export default App;
