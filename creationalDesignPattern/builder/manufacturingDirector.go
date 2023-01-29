package builder

type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) SetBuilder(builder BuildProcess) {
	f.builder = builder
}

func (f *ManufacturingDirector) StartManufacturing() {
	f.builder.SetSeats().SetWheels().SetStructure()
}
