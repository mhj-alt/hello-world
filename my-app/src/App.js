// import logo from './logo.svg';
import './App.css';

async function whenClick() {
  try {
    const response = await fetch('http://localhost:8080/ping', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });
    
    const data = await response.text();
    
    console.log('Response status:', response.status);
    console.log('Response data:', data);
    
    alert(`Response: ${data}`);
  } catch (error) {
    console.error('Error:', error);
    alert(`Error: ${error.message}`);
  }
}

function App() {
  return (
    <div className="App">
      <header className="App-header">
        {/* <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p> */}

        <div>

          <button onClick={whenClick}>ping</button>
        </div>
      </header>
    </div>
  );
}

export default App;
