package kubeclient

//Template variable key
const (
	VariableFlavorCPU       = "flavor.cpu"
	VariableFlavorMem       = "flavor.mem"
	VariableFlavorStorage   = "flaovr.storage"
	VariableFlavorNvidiaGpu = "flaovr.nvidia-gpu"
)

//In-Kubernetes consts
const (
	LabelAppName = "jupynetes.name"
)
