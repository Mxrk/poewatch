package poewatch

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
)

type Client struct {
	// todo
}

var baseURL = "https://api.poe.watch/"

//GetChangeID returns the latest change ID from the top of the river and the time
func GetChangeID() ID {
	ids := ID{}
	req := get(baseURL + "status")

	err := json.Unmarshal([]byte(req), &ids)
	if err != nil {
		log.Println(err)
	}

	return ids
}

//GetLeagues list the current leagues. Event leagues appear first.
func GetLeagues() Leagues {
	leagues := Leagues{}
	req := get(baseURL + "leagues")

	err := json.Unmarshal([]byte(req), &leagues)
	if err != nil {
		log.Println(err)
	}

	return leagues
}

//GetItemData returns all found items via the stash api
func GetItemData() ItemData {
	data := ItemData{}
	req := get(baseURL + "itemdata")

	err := json.Unmarshal([]byte(req), &data)
	if err != nil {
		log.Println(err)
	}

	return data
}

//GetCharacters returns the player character names found through the stash api
func GetCharacters(name string) Characters {
	characters := Characters{}
	req := get(baseURL + "account?name=" + name)

	err := json.Unmarshal([]byte(req), &characters)
	if err != nil {
		log.Println(err)
	}

	return characters
}

//GetAccount returns player account name via a char name
func GetAccount(name string) Accounts {
	accounts := Accounts{}
	req := get(baseURL + "char?name=" + name)

	err := json.Unmarshal([]byte(req), &accounts)
	if err != nil {
		log.Println(err)
	}

	return accounts
}

//GetCategories returns all listed categories
func GetCategories() Categories {
	categories := Categories{}
	req := get(baseURL + "categories")

	err := json.Unmarshal([]byte(req), &categories)
	if err != nil {
		log.Println(err)
	}

	return categories
}

//GetItems returns price and item for specified league and category.
//It's advisable to use a singular request to the compact API, rather than multiple requests to this API.
func GetItems(league, category string) Items {
	items := Items{}
	req := get(baseURL + "get?league=" + league + "&category=" + category)

	err := json.Unmarshal([]byte(req), &items)
	if err != nil {
		log.Println(err)
	}

	return items
}

//GetItemsFilter returns price and item for specified league and category with a given filter.
func GetItemsFilter(league, category string, args FilterArgs) Items {
	items := Items{}

	values := url.Values{}

	values.Add("league", league)
	values.Add("category", category)

	if args.GemLevel != 0 {
		values.Add("gemLevel", strconv.Itoa(args.GemLevel))
	}

	if args.GemCorrupted {
		values.Add("gemCorrupted", strconv.FormatBool(args.GemCorrupted))
	}

	if args.LowConfidence {
		values.Add("lowConfidence", strconv.FormatBool(args.LowConfidence))
	}

	if args.ItemLevel != 0 {
		values.Add("itemLevel", strconv.Itoa(args.ItemLevel))
	}

	if args.LinkCount != 0 {
		values.Add("linkCount", strconv.Itoa(args.LinkCount))
	}

	req := get(baseURL + "get" + values.Encode())
	err := json.Unmarshal([]byte(req), &items)
	if err != nil {
		log.Println(err)
	}

	return items
}

//GetCompact returns price data of all items provided in the active league
func GetCompact(league string) Compact {
	compactData := Compact{}
	req := get(baseURL + "compact?league=" + league)

	err := json.Unmarshal([]byte(req), &compactData)
	if err != nil {
		log.Println(err)
	}

	return compactData
}

//GetItem returns information about a specific item.
func GetItem(itemID string) Item {
	item := Item{}
	req := get(baseURL + "id?id=" + itemID)

	err := json.Unmarshal([]byte(req), &item)
	if err != nil {
		log.Println(err)
	}

	return item
}

//GetItemHistory returns prices from past leagues.
func GetItemHistory(itemID, league string) History {
	history := History{}
	req := get(baseURL + "history?id=" + itemID + "&league=" + league)

	err := json.Unmarshal([]byte(req), &history)
	if err != nil {
		log.Println(err)
	}

	return history
}

//GetEnchants returns Enchant data for a given Item in a league.
func GetEnchants(itemID string, league string) []Enchants {
	enchants := []Enchants{}
	req := get(baseURL + "enchants?id=" + itemID + "&league=" + league)

	err := json.Unmarshal([]byte(req), &enchants)
	if err != nil {
		log.Println(err)
	}

	return enchants
}

//GetCorruptions returns corruption data for a given item
func GetCorruptions(itemID, league string) []Corruptions {
	corruptions := []Corruptions{}
	req := get(baseURL + "corruptions?id=" + itemID + "&league=" + league)

	err := json.Unmarshal([]byte(req), &corruptions)
	if err != nil {
		log.Println(err)
	}

	return corruptions
}
