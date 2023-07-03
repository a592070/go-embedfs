import './App.css';
import {useLayoutEffect} from "react";
import {Link} from "react-router-dom";

function App() {
    useLayoutEffect(() => {
        document.title = 'App';
    }, []);

  return (
    <div className="App">
        <p>
            <Link to={`/about`}>About</Link>
        </p>
        <p>
            <Link to={`/404`}>404</Link>
        </p>
    </div>
  );
}

export default App;
