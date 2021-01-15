import { observable, makeObservable } from "mobx";
import { ServerRepository } from './api/server/ServerRepository';
import { ServerObject } from "./model/ServerObject";


export default class AppState {
    servers:ServerObject[] = [];

    constructor() {
        makeObservable(this, {
        servers: observable,
        });
    }

    async getInstances() {
        let serverRepository:ServerRepository =  new ServerRepository();
        let results = await serverRepository.getServers();
        this.servers.concat(results);
    }
}

const appState = new AppState();