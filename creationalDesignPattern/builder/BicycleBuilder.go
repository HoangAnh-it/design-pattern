package builder

/*implement buildProcess*/
type BicycleBuilder struct {
	vehicle Vehicle
}

func (c *BicycleBuilder) SetWheels() BuildProcess {
	c.vehicle.Wheels = 2
	return c
}

func (c *BicycleBuilder) SetSeats() BuildProcess {
	c.vehicle.Seats = 1
	return c
}

func (c *BicycleBuilder) SetStructure() BuildProcess {
	c.vehicle.Structure = "bicycle"
	return c
}

func (c *BicycleBuilder) GetVehicle() Vehicle {
	return c.vehicle
}
