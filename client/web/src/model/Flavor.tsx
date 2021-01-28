export class Flavor {
  cpu: number;
  memory: number;
  nvidia_gpu: number;

  constructor() {
      this.cpu = 0;
      this.memory = 0;
      this.nvidia_gpu = 0;
  }
}