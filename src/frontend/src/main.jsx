import React from 'react';
import ReactDOM from 'react-dom/client';
import Circle from './components/circle';
import './style.css';

function App() {
  return (
    <div className="circle-wrapper">
      {/* Full circle */}
      <Circle
        size={150}
        color="royalblue"
        visiblePercentage={100}
        style={{
          position: 'absolute',
          top: '50%',
          left: '50%',
          transform: 'translate(-50%, -50%)',
        }}
      />

      {/* Partial circle on top */}
      <Circle
        size={150}
        color="tomato"
        visiblePercentage={50}
        style={{
          position: 'absolute',
          top: '50%',
          left: '50%',
          transform: 'translate(-50%, -50%)',
        }}
      />
    </div>
  );
}

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);