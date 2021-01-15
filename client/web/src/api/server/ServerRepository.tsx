import { plainToClass } from "class-transformer";
import { ServerObject } from "../../model/ServerObject";

export class ServerRepository {

    provider: ServerFakeProvider;

    constructor() {
        this.provider = new ServerFakeProvider();
    }

    async getServers(){
        let result:string = await this.provider.getServers();
        let json = JSON.parse(result);
        let servers = plainToClass(ServerObject, json as any[]);
        return servers;
    }

    async postServer(){
        let result:string = await this.provider.postServer();
        let json = JSON.parse(result);
        let servers = plainToClass(ServerObject, json);
        return servers;
    }

    async getServer(){
        let result:string = await this.provider.getServer();
        let json = JSON.parse(result);
        let servers = plainToClass(ServerObject, json);
        return servers;
    }

    async deleteServer(){
        let result:boolean = await this.provider.deleteServer();
        return result;
    }
}

class ServerFakeProvider {

    constructor(){}

    getServers():string {
        return `[
            {
                "id": "string",
                "name": "string",
                "description": "string",
                "template": {
                "id": 0,
                "name": "string",
                "description": "string",
                "bdoy": "string",
                "variables": [
                    {
                    "name": "string",
                    "value": "string"
                    }
                ]
                },
                "flavor": {
                "cpu": 0,
                "memory": 0,
                "nvidia_gpu": 0
                },
                "created_at": "2021-01-30T08:30:00Z",
                "status": "running",
                "message": "string",
                "last_transition_time": "2021-01-30T08:30:00Z",
                "last_probe_time": "2021-01-30T08:30:00Z",
                "owner": [
                "2016920036",
                "admin"
                ]
            }
        ]`;
    }

    postServer():string {
        return `{
            "id": "string",
            "name": "string",
            "description": "string",
            "template": {
                "id": 0,
                "name": "string",
                "description": "string",
                "bdoy": "string",
                "variables": [
                {
                    "name": "string",
                    "value": "string"
                }
                ]
            },
            "flavor": {
                "cpu": 0,
                "memory": 0,
                "nvidia_gpu": 0
            },
            "created_at": "2021-01-30T08:30:00Z",
            "status": "running",
            "message": "string",
            "last_transition_time": "2021-01-30T08:30:00Z",
            "last_probe_time": "2021-01-30T08:30:00Z",
            "owner": [
                "2016920036",
                "admin"
            ]
            }`;
    }

    getServer():string {
        return `{
            "id": "string",
            "name": "string",
            "description": "string",
            "template": {
                "id": 0,
                "name": "string",
                "description": "string",
                "bdoy": "string",
                "variables": [
                {
                    "name": "string",
                    "value": "string"
                }
                ]
            },
            "flavor": {
                "cpu": 0,
                "memory": 0,
                "nvidia_gpu": 0
            },
            "created_at": "2021-01-30T08:30:00Z",
            "status": "running",
            "message": "string",
            "last_transition_time": "2021-01-30T08:30:00Z",
            "last_probe_time": "2021-01-30T08:30:00Z",
            "owner": [
                "2016920036",
                "admin"
            ]
        }`;
    }

    deleteServer():boolean {
        return true;
    }

}
