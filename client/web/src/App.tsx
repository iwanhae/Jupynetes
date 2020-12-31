import React, { FC } from 'react';
import { Button } from 'antd';
import './App.css';
import Dashboard from "./pages/Dashboard/Dashboard.js"

const App: FC = () => (
  <div className="App">
    <Dashboard/>
  </div>
);

export default App;
