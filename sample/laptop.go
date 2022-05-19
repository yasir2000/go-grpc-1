package sample

// this module implements a constructor for all Laptop Object structure sample in functions
import (
	"github.com/golang/protobuf/ptypes"
	"gitlab.com/yasir2000/go-grpc-1/pb"
)

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

// NewCPU returns a new sample CPU
func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)
	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

// NewGPU returns a new sample GPU
func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memGB := randomInt(2, 6)

	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: &pb.Memory{
			Value: uint64(memGB),
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	return gpu
}

// NewSSD returns a new sample SSD
func NewSSD() *pb.Storage {
	memGB := randomInt(128, 1024)

	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(memGB),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return ssd
}

// NewHDD returns a new sample HDD
func NewHDD() *pb.Storage {
	memTB := randomInt(1, 6)

	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(memTB),
			Unit:  pb.Memory_TERABYTE,
		},
	}
	return hdd
}

//returns a new sample Screen
func NewScreen() *pb.Laptop {
	brand := randomlaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pb.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3500),
		Releaseyear: uint32(randomInt(2015, 2019)),
		UpdatesAt:   ptypes.TimestampNow(),
	}
	return laptop
}

// RandomLaptopScore returns a random laptop score
func RandomLaptopScore() float64 {
	return float64(randomInt(1, 10))
}
