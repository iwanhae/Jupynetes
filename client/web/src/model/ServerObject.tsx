
import assert from "assert";
import { Flavor } from "./Flavor";
import { Template } from "./Template";

export enum ServerStatus {
  ERROR,
  STOPPED,
  INITIALIZING,
  RUNNING
}

var dateFormat = require('dateformat');

export class ServerObject {
  id: string;
  name: string;
  description: string;
  template: Template;
  flavor: Flavor;
  created_at: string;
  status: string;
  message: string;
  last_transition_time: string;
  last_probe_time: string;
  owner: string[];

  constructor() {
    this.id = "3";
    this.name = "sdfdf";
    this.description = "dfdf";
    this.template = new Template();
    this.flavor = new Flavor();
    this.created_at = "";
    this.status = "error";
    this.message = "dfdf";
    this.last_transition_time = "";
    this.last_probe_time = "";
    this.owner = [];
  }

  getNDaysAgo():string {

    if(this.last_transition_time == null) {
      return "undefined";
    }

    console.log("last_transition_time: " + this.last_transition_time);

    let now:number = Date.now();
    let lastTransitionTime:number = Date.parse(this.last_transition_time);
    let diffInMilliSeconds:number = now - lastTransitionTime;
    let diffDays = Math.floor(diffInMilliSeconds / 86400000); // days
    let diffHrs = Math.floor((diffInMilliSeconds % 86400000) / 3600000); // hours
    let diffMins = Math.round(((diffInMilliSeconds % 86400000) % 3600000) / 60000); // m

    console.log(diffInMilliSeconds);

    if(diffHrs < 1) {
      return diffMins + " 분";
    } else if (diffDays < 1) {
      return diffHrs + " 시간";
    } else {
      return diffDays + " 일";
    }
  }

  getFormattedDate(date:string):string {
    return dateFormat(date, "yyyy년 mm월 dd일 h시 MM분 ss초");
  }

  getFormmatedCreatedAt():string {
    return this.getFormattedDate(this.created_at);
  }

  getStatus():ServerStatus {
    switch(this.status) { 
      case "error": { 
          return ServerStatus.ERROR;
      } 
      case "stopped": { 
          return ServerStatus.STOPPED; 
      }
      case "initializing": {
        return ServerStatus.INITIALIZING;
      }
      case "running": {
        return ServerStatus.RUNNING;
      }
      default: {
        assert(false);
      } 
    } 
  }
  
}