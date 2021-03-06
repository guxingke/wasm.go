package interpreter

func localGet(vm *vm, args interface{}) {
	idx := args.(uint32)
	val := vm.getLocal(vm.local0Idx + idx)
	vm.pushU64(val)
}
func localSet(vm *vm, args interface{}) {
	idx := args.(uint32)
	val := vm.popU64()
	vm.setLocal(vm.local0Idx+idx, val)
}
func localTee(vm *vm, args interface{}) {
	idx := args.(uint32)
	val := vm.popU64()
	vm.pushU64(val)
	vm.setLocal(vm.local0Idx+idx, val)
}

func globalGet(vm *vm, args interface{}) {
	idx := args.(uint32)
	val := vm.globals[idx].Get()
	vm.pushU64(val)
}
func globalSet(vm *vm, args interface{}) {
	idx := args.(uint32)
	val := vm.popU64()
	vm.globals[idx].Set(val)
}
