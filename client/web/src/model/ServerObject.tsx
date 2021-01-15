import { Flavor } from "./Flavor";
import { Template } from "./Template";
import { TemplateVariable } from "./TemplateVariables";

export class ServerObject {
  id: string;
  name: string;
  description: string;
  template: Template;
  flavor: Flavor;
  created_at: string;
  status: string;
  message: string;
  last_translation_time: string;
  last_probe_time: string;
  owner: string[];

  constructor() {
    this.id = "";
    this.name = "";
    this.description = "";
    this.template = new Template();
    this.flavor = new Flavor();
    this.created_at = "";
    this.status = "";
    this.message = "";
    this.last_translation_time = "";
    this.last_probe_time = "";
    this.owner = [];
  }
}