import { plainToClass } from "class-transformer";
import { Quota } from "../../model/Quota";
import { Template } from "../../model/Template";
import { UserInfo } from "../../model/UserInfo";

export class AdminRepository {

    provider: AdminFakeProvider;

    constructor() {
        this.provider = new AdminFakeProvider();
    }

    async getUser(){
        let result:string = await this.provider.getUser();
        let json = JSON.parse(result);
        let users = plainToClass(UserInfo, json as any[]);
        return users;
    }

    async postAdminTemplate(){
        let result:string = await this.provider.postAdminTemplate();
        let json = JSON.parse(result);
        let template = plainToClass(Template, json);
        return template;
    }

    async postAdminQuota(){
        let result:string = await this.provider.postAdminQuota();
        let json = JSON.parse(result);
        let quota = plainToClass(Quota, json);
        return quota;
    }

    async postUserQuota(){
        let result:string = await this.provider.postUserQuota();
        let json = JSON.parse(result);
        let quota = plainToClass(Quota, json);
        return quota;
    }
}

class AdminFakeProvider {

    constructor(){}

    getUser():string {
        return `[
                    {
                        "id": "string",
                        "user_quota": {
                            "instance": 1,
                            "cpu": 2,
                            "memmory": 3,
                            "nvidia_gpu": 4,
                            "storage": 5
                        },
                        "group_quota": {
                            "instance": 6,
                            "cpu": 7,
                            "memmory": 8,
                            "nvidia_gpu": 9,
                            "storage": 10
                        }
                    }
                ]`;
    }

    postAdminTemplate():string {
        return `{
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
        }`;
    }

    postAdminQuota():string {
        return `{
            "instance": 0,
            "cpu": 0,
            "memory": 0,
            "nvidia_gpu": 0,
            "storage": 0
        }`;
    }

    postUserQuota():string {
        return `{
            "instance": 0,
            "cpu": 0,
            "memory": 0,
            "nvidia_gpu": 0,
            "storage": 0
        }`;
    }
}
