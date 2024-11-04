package Phonebook

import (
	"testing"
)

func TestCountParkSpot(t *testing.T) {
	// Input given in this format.
	parking := [][][]int{
		{
			{4, 4, 2, 2},
			{2, 4, 2, 0},
			{0, 2, 2, 2},
			{4, 4, 4, 0},
		},
	}
	parkLot := make([][][]vInfo, len(parking))
	for i := range parking {
		parkLot[i] = make([][]vInfo, len(parking[i]))
		for j := range parking[i] {
			parkLot[i][j] = make([]vInfo, len(parking[i][j]))
			for k := range parking[i][j] {
				parkLot[i][j][k] = vInfo{
					vehicleType: parking[i][j][k],
					isParked:    false,
				}
			}
		}
	}
	logger := NewHelper("Phonebook")

	pkSystem := NewParkSystem(parkLot, logger)
	freeSlot := pkSystem.getFreeSpotsCount(0, 2)
	if freeSlot != 7 {
		t.Errorf("FreeSpot is %d, want 3", freeSlot)
	}
}

func TestPark(t *testing.T) {
	parking := [][][]int{
		{
			{4, 4, 2, 2},
			{2, 4, 2, 0},
			{0, 2, 2, 2},
			{4, 4, 4, 0},
		},
	}
	parkLot := make([][][]vInfo, len(parking))
	for i := range parking {
		parkLot[i] = make([][]vInfo, len(parking[i]))
		for j := range parking[i] {
			parkLot[i][j] = make([]vInfo, len(parking[i][j]))
			for k := range parking[i][j] {
				parkLot[i][j][k] = vInfo{
					vehicleType: parking[i][j][k],
					isParked:    false,
				}
			}
		}
	}
	logger := NewHelper("Phonebook")

	pkSystem := NewParkSystem(parkLot, logger)
	spotId := pkSystem.park(2, "1234", "kdkd", 0)
	if spotId != "0-0-2" {
		t.Errorf("Spot is %s, want 0-0-2", spotId)
	}
	spotId2 := pkSystem.park(4, "1234", "kdkd", 1)
	if spotId2 != "0-0-0" {
		t.Errorf("Spot is %s, want 0-0-0", spotId2)
	}
}

func TestRemoveVehicle(t *testing.T) {
	parking := [][][]int{
		{
			{4, 4, 2, 2},
			{2, 4, 2, 0},
			{0, 2, 2, 2},
			{4, 4, 4, 0},
		},
	}
	parkLot := make([][][]vInfo, len(parking))
	for i := range parking {
		parkLot[i] = make([][]vInfo, len(parking[i]))
		for j := range parking[i] {
			parkLot[i][j] = make([]vInfo, len(parking[i][j]))
			for k := range parking[i][j] {
				parkLot[i][j][k] = vInfo{
					vehicleType: parking[i][j][k],
					isParked:    false,
				}
			}
		}
	}
	logger := NewHelper("Phonebook")
	pkSystem := NewParkSystem(parkLot, logger)
	spotId := pkSystem.park(2, "1234", "kdkd", 0)

	checkBool := pkSystem.removeVehicle(spotId)
	if !checkBool {
		t.Errorf("Vehicle should be removed, expected true, got %v", checkBool)
	}

	checkRemoval := pkSystem.removeVehicle("0-2-3")
	if checkRemoval {
		t.Errorf("Vehicle should not be removed, expected false, got %v", checkRemoval)
	}
}

func TestSearchVehicle(t *testing.T) {
	parking := [][][]int{
		{
			{4, 4, 2, 2},
			{2, 4, 2, 0},
			{0, 2, 2, 2},
			{4, 4, 4, 0},
		},
	}
	parkLot := make([][][]vInfo, len(parking))
	for i := range parking {
		parkLot[i] = make([][]vInfo, len(parking[i]))
		for j := range parking[i] {
			parkLot[i][j] = make([]vInfo, len(parking[i][j]))
			for k := range parking[i][j] {
				parkLot[i][j][k] = vInfo{
					vehicleType: parking[i][j][k],
					isParked:    false,
				}
			}
		}
	}
	logger := NewHelper("Phonebook")
	pkSystem := NewParkSystem(parkLot, logger)
	spotId := pkSystem.park(2, "1234", "kdkd", 0)
	spotIdBooking := pkSystem.searchVehicle("1234")

	if spotId != spotIdBooking {
		t.Errorf("Vehicle spotId didn't match, expected %s, got %s", spotId, spotIdBooking)
	}
}
