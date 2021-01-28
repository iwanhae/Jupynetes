import { plainToClass } from "class-transformer";
import { ServerResponse } from "http";
import { ServerObject } from "../../model/ServerObject";

export class ServerRepository {

    provider: ServerFakeProvider;

    constructor() {
        this.provider = new ServerFakeProvider();
    }

    async getServers(){
        let result:string = await this.provider.getServers();
        console.log("getServers: " + result);
        let json = JSON.parse(result);
        console.log("getServers: " + json);
        let servers = plainToClass(ServerObject, json as any[]);
        console.log("length of server" + servers.length);
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
                "id": "1",
                "name": "사과서버",
                "description": "사과서버입니다.",
                "template": {
                "id": 0,
                "name": "사과템플릿",
                "description": "사과템플릿입니다.",
                "bdoy": "string",
                "variables": [
                    {
                    "name": "string",
                    "value": "string"
                    }
                ]
                },
                "flavor": {
                "cpu": 55,
                "memory": 22,
                "nvidia_gpu": 44
                },
                "created_at": "2021-01-02T08:30:00Z",
                "status": "running",
                "message": "string",
                "last_transition_time": "2021-01-02T08:30:00Z",
                "last_probe_time": "2021-01-02T08:30:00Z",
                "owner": [
                "2016920036",
                "admin"
                ]
            },
            {
                "id": "2",
                "name": "오렌지서버",
                "description": "사과서버입니다.",
                "template": {
                "id": 0,
                "name": "사과템플릿",
                "description": "사과템플릿입니다.",
                "bdoy": "string",
                "variables": [
                    {
                    "name": "string",
                    "value": "string"
                    }
                ]
                },
                "flavor": {
                "cpu": 55,
                "memory": 22,
                "nvidia_gpu": 44
                },
                "created_at": "2021-01-11T08:30:00Z",
                "status": "error",
                "message": "string",
                "last_transition_time": "2021-01-11T08:30:00Z",
                "last_probe_time": "2021-01-11T08:30:00Z",
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
