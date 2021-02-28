import React from 'react';
import appState from '../../../AppState';
import InstanceList from '../Components/InstanceList';

export class InstanceListContainer extends React.Component {
  async componentDidMount() {
    await appState.getInstances();
    console.log('InstanceListContainer componentDidMount');
    console.log('lenth of appState.servers: ' + appState.servers.length);
    this.setState({});
  }

  render() {
    return <InstanceList></InstanceList>;
  }
}
