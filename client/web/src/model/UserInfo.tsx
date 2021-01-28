
import { Quota } from "./Quota";

export class UserInfo {
  id: string;
  
  userQuota: Quota;
  groupQutoa: Quota;
  

  constructor(){
    this.id = "";
    this.userQuota = new Quota();
    this.groupQutoa = new Quota();
  }
}