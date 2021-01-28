import { TemplateVariable } from "./TemplateVariables";

export class Template {
  id: number;
  name: string;
  description: string;
  body: string;
  templateVariables: TemplateVariable[];


  constructor() {
      this.id = 0;
      this.name = "";
      this.description = "";
      this.body = "";
      this.templateVariables = [];
  }
}