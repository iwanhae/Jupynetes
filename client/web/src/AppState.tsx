import { observable, makeObservable, action } from "mobx";
import { ServerRepository } from './api/server/ServerRepository';
import { ServerObject } from "./model/ServerObject";


class AppState {
    @observable servers:ServerObject[] = [];

     constructor() {
        makeObservable(this, {
            servers: observable,
        });
    }

    @action getInstances = async () => {
        let serverRepository:ServerRepository =  new ServerRepository();
        let results = await serverRepository.getServers();
        console.log("getInstance: " + results.length);
        this.servers = results;
        console.log("appState length of servers: " + this.servers.length);
    }

    @action deleteServer = (instance:ServerObject) => {
       const newServers = this.servers.filter(item=> item !== instance);
       this.servers = newServers;
       console.log("server length: " + this.servers.length);
    };
}

const appState = new AppState();

export default appState;