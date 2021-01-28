import { plainToClass } from "class-transformer";
import { UserInfo } from "../../model/UserInfo";


export class UserRepository {

    provider: UserFakeProvider;

    constructor() {
        this.provider = new UserFakeProvider();
    }

    async getAdminUser(){
        let result:string = await this.provider.getAdminUser();
        let json = JSON.parse(result);
        let users = plainToClass(UserInfo, json);
        return users;
    }

    async getUser(){
        let result:string = await this.provider.getUser();
        let json = JSON.parse(result);
        let users = plainToClass(UserInfo, json);
        return users;
    }

    async postUser(password:string, newPasswrod:string){
        return true;
    }

    async postLogin(id:string, password:string){
        return true;
    }

    async getLogout(){
        return true;
    }
}

class UserFakeProvider {

    constructor(){}

    getAdminUser():string {
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
}
