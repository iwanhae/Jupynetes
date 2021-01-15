import { plainToClass } from "class-transformer";
import { UserInfo } from "../../model/UserInfo";


export class TemplateRepository {

    provider: TemplateFakeProvider;

    constructor() {
        this.provider = new TemplateFakeProvider();
    }

    async postAdminTemplate(){
        let result:string = await this.provider.postAdminTemplate();
        let json = JSON.parse(result);
        let users = plainToClass(UserInfo, json);
        return users;
    }

    async getTemplate(){
        let result:string = await this.provider.getTemplate();
        let json = JSON.parse(result);
        let users = plainToClass(UserInfo, json);
        return users;
    }
}

class TemplateFakeProvider {

    constructor(){}

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

    getTemplate():string {
        return `[
            {
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
            }
        ]`;
    }
}
