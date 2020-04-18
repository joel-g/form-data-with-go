package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
)

var path = "https://vc0.publicstuff.com/api/2.0/request_submit?device=iframe"

func main() {
	for _, pair := range coords {
		resp, err := http.PostForm(path, generateReportBody(pair[0], pair[1]))
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		fmt.Println(resp.StatusCode)
	}
}

func generateReportBody(lat, lon string) url.Values {
	return url.Values{
		"title":           {"Report Gatherings"},
		"description":     {randomAmmendment()},
		"request_type_id": {"1017049"},
		"is_private":      {"0"},
		"latitude":        {lat},
		"longitude":       {lon},
		"address":         {""},
		"user_address":    {""},
		"long_address":    {""},
		"zipcode":         {""},
		"space_id":        {"4833"},
		"client_id":       {"954"},
	}
}

func randomAmmendment() string {
	rights := []string{
		"Congress shall make no law respecting an establishment of religion, or prohibiting the free exercise thereof; or abridging the freedom of speech, or of the press; or the right of the people peaceably to assemble, and to petition the government for a redress of grievances.",
		"A well regulated militia, being necessary to the security of a free state, the right of the people to keep and bear arms, shall not be infringed.",
		"No soldier shall, in time of peace be quartered in any house, without the consent of the owner, nor in time of war, but in a manner to be prescribed by law.",
		"The right of the people to be secure in their persons, houses, papers, and effects, against unreasonable searches and seizures, shall not be violated, and no warrants shall issue, but upon probable cause, supported by oath or affirmation, and particularly describing the place to be searched, and the persons or things to be seized.",
		"No person shall be held to answer for a capital, or otherwise infamous crime, unless on a presentment or indictment of a grand jury, except in cases arising in the land or naval forces, or in the militia, when in actual service in time of war or public danger; nor shall any person be subject for the same offense to be twice put in jeopardy of life or limb; nor shall be compelled in any criminal case to be a witness against himself, nor be deprived of life, liberty, or property, without due process of law; nor shall private property be taken for public use, without just compensation.",
		"In all criminal prosecutions, the accused shall enjoy the right to a speedy and public trial, by an impartial jury of the state and district wherein the crime shall have been committed, which district shall have been previously ascertained by law, and to be informed of the nature and cause of the accusation; to be confronted with the witnesses against him; to have compulsory process for obtaining witnesses in his favor, and to have the assistance of counsel for his defense.",
		"In suits at common law, where the value in controversy shall exceed twenty dollars, the right of trial by jury shall be preserved, and no fact tried by a jury, shall be otherwise reexamined in any court of the United States, than according to the rules of the common law.",
		"Excessive bail shall not be required, nor excessive fines imposed, nor cruel and unusual punishments inflicted.",
		"The enumeration in the Constitution, of certain rights, shall not be construed to deny or disparage others retained by the people.",
		"The powers not delegated to the United States by the Constitution, nor prohibited by it to the states, are reserved to the states respectively, or to the people.",
	}
	return rights[rand.Intn(10)]
}

var coords = [][]string{{"47.62266", "-122.14746"}, {"47.62018", "-122.15174"}, {"47.61817", "-122.15405"}, {"47.61615", "-122.15705"}, {"47.61178", "-122.16218"}, {"47.60695", "-122.16971"}, {"47.60096", "-122.17655"}, {"47.59633", "-122.18569"}, {"47.5921", "-122.1864"}, {"47.58649", "-122.18436"}, {"47.58313", "-122.17973"}, {"47.57977", "-122.17097"}, {"47.57843", "-122.16683"}, {"47.57754", "-122.16269"}, {"47.57717", "-122.15372"}, {"47.57715", "-122.14572"}, {"47.57922", "-122.13737"}, {"47.58591", "-122.13304"}, {"47.59462", "-122.13622"}, {"47.5973", "-122.1401"}, {"47.59874", "-122.14784"}, {"47.60208", "-122.14239"}, {"47.60669", "-122.13591"}, {"47.61084", "-122.13012"}, {"47.61355", "-122.12514"}, {"47.61856", "-122.11878"}, {"47.62212", "-122.11054"}, {"47.62575", "-122.10951"}, {"47.62942", "-122.1083"}, {"47.63309", "-122.10779"}, {"47.63536", "-122.10949"}, {"47.63759", "-122.11186"}, {"47.63952", "-122.11729"}, {"47.63689", "-122.12815"}, {"47.6343", "-122.13332"}, {"47.62986", "-122.13643"}, {"47.62634", "-122.14194"}}

// eyeOutline = {"47.61149",-122.15261}, {"47.61121,-122.16093}, {"47.61048,-122.16548}, {"47.60916,-122.17209}, {"47.60692,-122.1787}, {"47.60462,-122.18527}, {"47.60162,-122.19183}, {"47.59828,-122.19823}, {"47.59471,-122.20325}, {"47.59125,-122.19843}, {"47.58848,-122.19293}, {"47.58636,-122.18897}, {"47.58471,-122.18682}, {"47.58353,-122.18398}, {"47.58216,-122.17907}, {"47.58136,-122.17594}, {"47.58078,-122.17246}, {"47.58068,-122.16609}, {"47.58012,-122.15973}, {"47.5804,-122.15458}, {"47.58068,-122.14873}, {"47.58183,-122.14401}, {"47.58345,-122.13962}, {"47.58505,-122.13412}, {"47.58689,-122.12965}, {"47.58908,-122.12483}, {"47.59149,-122.12036}, {"47.59391,-122.11674}, {"47.59679,-122.11347}, {"47.60054,-122.11782}, {"47.60429,-122.1239}, {"47.60669,-122.12952}, {"47.60864,-122.134}, {"47.6099,-122.14123},

// eyeCenter = {"47.61114,-122.15844}, {"47.60981,-122.16501}, {"47.60721,-122.17055}, {"47.60556,-122.17297}, {"47.60368,-122.17436}, {"47.60157,-122.17542}, {"47.59945,-122.17578}, {"47.59666,-122.177}, {"47.59387,-122.17651}, {"47.59107,-122.17636}, {"47.58851,-122.17346}, {"47.58586,-122.17061}, {"47.58356,-122.16621}, {"47.58207,-122.16284}, {"47.58174,-122.1581}, {"47.58265,-122.15417}, {"47.58403,-122.15025}, {"47.58747,-122.1448}, {"47.59018,-122.14241}, {"47.59335,-122.14055}, {"47.59582,-122.1404}, {"47.59853,-122.14059}, {"47.60076,-122.14134}, {"47.60298,-122.14209}, {"47.60527,-122.14362}, {"47.60709,-122.14651}, {47.60993,-122.15248},
