import React, { FC } from 'react';
import { Button } from 'antd';
import './App.css';
import Dashboard from "./pages/Dashboard/Components/Dashboard";

import { createBrowserHistory } from "history";
import { Router, Route, Switch, Redirect } from "react-router-dom";
import Login from './pages/Login/Login';

const hist = createBrowserHistory();

const App: FC = () => (
  <div className="App">
    <Router history={hist}>
      <Switch>
        <Route path="/login"><Login></Login></Route>
        <Route path="/dashboard"><Dashboard></Dashboard></Route>
        <Redirect from="/" to="/login" />
      </Switch>
    </Router>,
  </div>
);

export default App;
