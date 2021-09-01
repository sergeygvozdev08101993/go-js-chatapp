import React, { Component } from "react";
import "./App.css";
import Header from "./components/Header/Header"
import { connect, sendMsg } from "./api";
import ChatHistory from "./components/ChatHistory/ChatHistory";

class App extends Component {

  constructor(props) {
    super(props);
    this.state = {
        chatHistory: []
    }
  }

  componentDidMount() {
      connect((msg) => {
          console.log("new message")
          this.setState(prevState => ({
              chatHistory: [...this.state.chatHistory, msg]
          }))
          console.log(this.state);
      });
  }

    send() {
    console.log("hello");
    sendMsg("hello");
  }

  render() {
    return (
        <div className="App">
          <Header />
            <ChatHistory chatHistory={this.state.chatHistory}/>
          <button onClick={this.send}>Click</button>
        </div>
    )
  }
}

export default App;
