export class Quota {
  instance: number;
  cpu: number;
  memory: number;
  nvidia_gpu: number;
  storage: number;

  constructor() {
      this.instance = 0;
      this.cpu = 0;
      this.memory = 0;
      this.nvidia_gpu = 0;
      this.storage = 0;
  }
}