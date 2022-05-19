package sample

// this module include utility functions that select random values and options from
// list or group of lists that sufficie Laptop object structure
import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"gitlab.com/yasir2000/go-grpc-1/pb"
)

// executed only once and before all functions, purpose here is to generate random from Seed value
// which is Unix time here, so for every execution attempt time changes the seed value of randomness
func init() {
	rand.Seed(time.Now().UnixNano())
}

// chooses random string from set of strings by identifying length of target string
// then return a slice of that randomly chosen string
func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

// returns random boolean either true or false
func randomBool() bool {
	return rand.Intn(2) == 1
}

// returns random integer value from between max and min by taking two values then
// min value added to random  integer reminder or integer of  difference between
// max and min plus one to give value between upper and lower
func randomInt(min, max int) int {
	return min + rand.Int()%(max-min+1)
}

// returns random float64 value from between max and min by taking two values then
// min value added to random  float64 reminder or float64 of  difference between
// max and min plus one to give value between upper and lower
func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

//// returns random float32 value from between max and min by taking two values then
// min value added to random  float32 reminder or float32 of  difference between
// max and min plus one to give value between upper and lower
func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

// returns random UUID by generating new UUID string everytime
func randomID() string {
	return uuid.New().String()
}

// returns selective random option from three Keyboard Layout types
func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

//returns a pointer to screen resolution, which is a random of value in
// a range multiplied  by a ratio, the reason pointer here is returned
// /is that screen resolution  is a struct and proto functions expects a
//pointer than a value, so the returned is to change pointer value
func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9

	resolution := &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}
	return resolution
}

// returns a one from two options of Laptop Screen types randomally
func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

// returns a CPU type string randomally from group string list of CPU's
func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	}
	return randomStringFromSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
		"Ryzen 3 PRO 3200GE",
	)
}

// returns a string option from randimally group of GPU list
func randomGPUBrand() string {
	return randomStringFromSet("Nvidia", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "Nvidia" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		)
	}
	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-56",
	)
}

// returns a Laptop Model brand randomally from  list of brands

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo")
}

// returns a Laptop Name randomally from list of brand groups of names and models
func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	default:
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad P53")

	}
}
