import { plainToClass } from "class-transformer";
import { Quota } from "../../model/Quota";
import { UserInfo } from "../../model/UserInfo";

export class QuotaRepository {

    provider: QuotaFakeProvider;

    constructor() {
        this.provider = new QuotaFakeProvider();
    }

    async postAdminQuota(){
        let result:string = await this.provider.postAdminQuota();
        let json = JSON.parse(result);
        let users = plainToClass(Quota, json);
        return users;
    }

    async postUserQuota(){
        let result:string = await this.provider.postUserQuota();
        let json = JSON.parse(result);
        let users = plainToClass(Quota, json);
        return users;
    }
}

class QuotaFakeProvider {

    constructor(){}

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
