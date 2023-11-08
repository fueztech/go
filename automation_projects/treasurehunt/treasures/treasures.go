package treasures

import (
	"math/rand"
	"time"
)

var Scriptures []SpiritualTreasure

type SpiritualTreasure struct {
	Scripture string `json:"scripture"`
	Text      string `json:"text"`
	Writer    string `json:"writer"`
}

type houseLocations struct {
	Room string `json:"Hiding Room"`
}

type timeToday struct {
	Time string `json:"time"`
}

// Embed the Structs together as shown below
type scriptureAndLocation struct {
	SpiritualTreasure SpiritualTreasure `json:"Spiritual Treasure"`
	HouseLocation     houseLocations    `json:"House Location"`
	Time              time.Time         `json:"Current Time"`
}

func ScriptureCollection() scriptureAndLocation {
	var Scriptures = []SpiritualTreasure{
		{Scripture: "Proverbs 29:23", Text: "Trembling at Men is a Snare", Writer: "Solomon"},
		{Scripture: "Psalms 46:10", Text: "Give in and Know that I am God", Writer: "Psalmist"},
		{Scripture: "Malachi 3:6", Text: "I am Jehovah, I do not Change..", Writer: "Malachi"},
		{Scripture: "Isaiah 42:5", Text: "The One who gives breath to the people", Writer: "Isaiah"},
		{Scripture: "Matthew 7:12", Text: "All Things Therefore that your want..", Writer: "Matthew"},
		{Scripture: "Job 33:6", Text: "God's own spirit made me", Writer: "Moses"},
		{Scripture: "Proverbs 18:14", Text: "A Person's spirit can sustain him through illness", Writer: "Solomon"},
		{Scripture: "Acts 17:27", Text: "He is not far off from each one of us", Writer: "Paul"},
		{Scripture: "Proverbs 10:19", Text: "When words are many, there does not fail to be transgression"},
		{Scripture: "John 8:54", Text: "If I glorify myself my glory is nothing", Writer: "John"},
		{Scripture: "1 Samuel 30:6", Text: "David strengthened himself by Jehovah his God", Writer: "Samuel"},
		{Scripture: "Ephesians 5:16", Text: "[Make] the best use of your time", Writer: "Paul"},
		{Scripture: "Ecclessiastes 3:11", Text: "He has made everything beautiful in its time", Writer: "Solomon"},
		{Scripture: "Daniel 7:9,10", Text: "The Ancient of Days sat down", Writer: "Daniel"},
		{Scripture: "Proverbs 22:3", Text: "The Shrewd One Sees the Danger", Writer: "Proverbs"},
		{Scripture: "Zechariah 8:6", Text: "Should it seem to difficult also to me?", Writer: "Zechariah"},
		{Scripture: "Psalms 32:8", Text: "I will give you insight and instruct you in the way you should go", Writer: "Psalmist"},
		{Scripture: "Exodus 23:2", Text: "You must not follow the crowd", Writer: "Moses"},
		{Scripture: "1 Corinthians 15:33", Text: "Bad Associations spoil useful habits...", Writer: "Paul"},
		{Scripture: "Matthew 6:33", Text: "Keep on seeking first the kingdom...", Writer: "Matthew"},
		{Scripture: "Romans 15:4", Text: "All things that were written beforehand were written for our instruction...", Writer: "Paul"},
		{Scripture: "Genesis 18:25", Text: "Will the Judge of all the earth not do what is right?", Writer: "Moses"},
		{Scripture: "Ecclessiastes 10:4", Text: "If the anger of a ruler flares up against you...", Writer: "Solomon"},
		{Scripture: "1 Peter 3:8", Text: "All of you have fellow feeling", Writer: "Peter"},
		{Scripture: "1 Corinthians 3:9", Text: "For we are God's fellow workers", Writer: "Paul"},
	}

	var houseRooms = []houseLocations{
		{Room: "Kitchen"},
		{Room: "Bathroom"},
		{Room: "Dining Room"},
		{Room: "Living Room"},
		{Room: "Girls Bedroom"},
		{Room: "Office"},
		{Room: "Master Bedroom"},
		{Room: "Hallway"},
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex1 := rand.Intn(len(Scriptures))
	randomIndex2 := rand.Intn(len(houseRooms))

	currentTime := time.Now()

	return scriptureAndLocation{
		SpiritualTreasure: Scriptures[randomIndex1],
		HouseLocation:     houseRooms[randomIndex2],
		Time:              currentTime,
	}
}

func AddScriptures(newTreasure SpiritualTreasure) {
	// Add the new set of scriptures to the slice of the scriptures slice
	Scriptures = append(Scriptures, newTreasure)
}
