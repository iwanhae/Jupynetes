import React, { Component } from 'react';


import ThemeProvider from '@material-ui/styles/ThemeProvider';
import { createMuiTheme,  MuiThemeProvider } from '@material-ui/core/styles';
import LoginScreen from './pages/login/LoginScreen';

const theme = createMuiTheme({
  typography: {
    fontFamily: "Roboto",
    fontSize: 12,
  },
});

class Main extends React.Component {
  render() {
    return (
      <ThemeProvider theme={theme}>
        <div>
          <LoginScreen></LoginScreen>
        </div>
      </ThemeProvider>
    );
  }
}

export default Main;