import logo from "./logo.svg";
import "./App.css";
import { Log } from "./utils";
import { Hello } from "./components";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <Hello />
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React <Log />
        </a>
      </header>
    </div>
  );
}

export default App;
