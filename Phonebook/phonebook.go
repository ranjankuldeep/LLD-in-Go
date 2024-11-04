package Phonebook

import (
	"fmt"
	"strconv"
	"strings"
)

type Helper struct {
	identifierStr string
}

func (h *Helper) print(val string) {
	fmt.Printf("log: %v %v", h.identifierStr, val)
}

func (h *Helper) println(val string) {
	fmt.Println(fmt.Printf("log: %v %v", h.identifierStr, val))
}

func NewHelper(value string) *Helper {
	return &Helper{}
}

type vehicleInfo struct {
	ticketId      string
	vehicleNumber string
	spotId        string
	vehicleType   int
}

type vInfo struct {
	vehicleType int
	isParked    bool
}

type ParkSystem struct {
	parkLot [][][]vInfo
	booking [][][]vehicleInfo
	helper  *Helper
}

func NewParkSystem(slots [][][]vInfo, logger *Helper) *ParkSystem {
	booking := make([][][]vehicleInfo, len(slots))
	for i := range booking {
		booking[i] = make([][]vehicleInfo, len(slots[i]))
		for j := range booking[i] {
			booking[i][j] = make([]vehicleInfo, len(slots[i][j]))
			for k := range booking[i][j] {
				booking[i][j][k] = vehicleInfo{}
			}
		}
	}
	return &ParkSystem{parkLot: slots,
		helper:  logger,
		booking: booking,
	}
}

func (p *ParkSystem) park(vehicleType int, vehicleNumber string, ticketId string, parkingStrategy int) string {
	switch parkingStrategy {
	case 0:
		for i := range p.parkLot {
			for j := range p.parkLot[i] {
				for k := range p.parkLot[i][j] {
					if p.parkLot[i][j][k].vehicleType == vehicleType {
						spotId := fmt.Sprintf("%v-%v-%v", i, j, k)
						if !p.parkLot[i][j][k].isParked {
							// Mark the slot to be booked
							p.parkLot[i][j][k].isParked = true
							// Confirm the booking of the vehicle
							p.booking[i][j][k] = vehicleInfo{
								ticketId:      ticketId,
								vehicleNumber: vehicleNumber,
								spotId:        spotId,
							}
						} else {
							return "already parked"
						}
						return spotId
					}
				}
			}
		}

	case 1:
		maxFreeSpotCount := -1
		bestFloor := -1

		for floor := range p.parkLot {
			count := p.getFreeSpotsCount(floor, vehicleType)
			if count > maxFreeSpotCount || (count == maxFreeSpotCount && bestFloor > floor) {
				maxFreeSpotCount = count
				bestFloor = floor
			}
		}

		if bestFloor != -1 {
			for j := range p.parkLot[bestFloor] {
				for k := range p.parkLot[bestFloor][j] {
					if p.parkLot[bestFloor][j][k].vehicleType == vehicleType {
						spotId := fmt.Sprintf("%v-%v-%v", bestFloor, j, k)
						if !p.parkLot[bestFloor][j][k].isParked {
							// Mark the slot to booked
							p.parkLot[bestFloor][j][k].isParked = true
							// Confirm the address of the vehicle
							p.booking[bestFloor][j][k] = vehicleInfo{
								ticketId:      ticketId,
								vehicleNumber: vehicleNumber,
								spotId:        spotId,
							}
						} else {
							return "already parked"
						}
						return spotId
					}
				}
			}
		}

	}
	return "nil"
}

func (p *ParkSystem) removeVehicle(spotId string) bool {
	// Optimized way of removing the vehicle
	parts := strings.Split(spotId, "-")
	floor, _ := strconv.Atoi(parts[0])
	row, _ := strconv.Atoi(parts[1])
	col, _ := strconv.Atoi(parts[2])

	if !p.parkLot[floor][row][col].isParked {
		return false
	}
	// Remove the parking
	p.parkLot[floor][row][col].isParked = false
	p.booking[floor][row][col] = vehicleInfo{} // empty vehicleInfo
	return true
}

func (p *ParkSystem) searchVehicle(vehicleNumber string) string {
	for i := range p.booking {
		for j := range p.booking[i] {
			for k := range p.booking[i][j] {
				if p.booking[i][j][k].vehicleNumber == vehicleNumber {
					return p.booking[i][j][k].spotId
				}
			}
		}
	}
	return ""
}

func (p *ParkSystem) getFreeSpotsCount(floor int, vehicleType int) int {
	count := 0
	for i := range p.parkLot[floor] {
		for j := range p.parkLot[floor][i] {
			if p.parkLot[floor][i][j].vehicleType == vehicleType {
				count++
			}
		}
	}
	return count
}
