package space

type Planet string

const earthAgeSeconds = 31557600

func Age( s float64, p Planet) float64 {
	switch p {
	case "Mercury":
		return s/(earthAgeSeconds*0.2408467)
	case "Venus":
		return s/(earthAgeSeconds*0.61519726)
	case "Earth":
		return s/(earthAgeSeconds)
	case "Mars":
		return s/(earthAgeSeconds*1.8808158)
	case "Jupiter":
		return s/(earthAgeSeconds*11.862615)
	case "Saturn":
		return s/(earthAgeSeconds*29.447498)
	case "Uranus":
		return s/(earthAgeSeconds*84.016846)
	case "Neptune":
		return s/(earthAgeSeconds*164.79132)
	default: 
		return 0
	}
}
