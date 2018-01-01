package space

//I tried doing this with map however for some reason mapping had worse benchmarks
//than switch... which is very interesting.

//Planet is name of the planet
type Planet string

var earthYear float64 = 31557600
var mercuryYear = earthYear * 0.2408467
var venusYear = earthYear * 0.61519726
var marsYear = earthYear * 1.8808158
var jupiterYear = earthYear * 11.862615
var saturnYear = earthYear * 29.447498
var uranusYear = earthYear * 84.016846
var neptuneYear = earthYear * 164.79132

//Age is the age determined by a different planet
func Age(seconds float64, planet Planet) float64 {
	var newAge float64
	switch planet {
	case "Earth":
		newAge = seconds / earthYear
	case "Mercury":
		newAge = seconds / mercuryYear
	case "Venus":
		newAge = seconds / venusYear
	case "Mars":
		newAge = seconds / marsYear
	case "Jupiter":
		newAge = seconds / jupiterYear
	case "Saturn":
		newAge = seconds / saturnYear
	case "Uranus":
		newAge = seconds / uranusYear
	case "Neptune":
		newAge = seconds / neptuneYear
	}
	return newAge
}
