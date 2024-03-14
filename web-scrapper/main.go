package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type PokemonProduct struct{
	url,image,name,price string
}


func main() {
	var pokemonProducts []PokemonProduct

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request){
		fmt.Println("Visiting: ",r.URL)
	})

	c.OnError(func(_ *colly.Response,err error){
		log.Println("Something went wrong:",err)
	})

	c.OnResponse(func(r *colly.Response){
		fmt.Print("Page visited:", r.Request.URL)
	})
	c.OnHTML("li.product",func(e *colly.HTMLElement){
		pokemonProduct := PokemonProduct{}
		
		pokemonProduct.url = e.ChildAttr("a","href")
		pokemonProduct.image = e.ChildAttr("img","src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.price = e.ChildText(".price")

		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})
	c.OnScraped(func(r *colly.Response){
		fmt.Println(r.Request.URL,"scraped !")
	})


	c.Visit("https://scrapeme.live/shop/")

	//openingg the csv file

	file,err :=  os.Create("products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file",err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	headers := []string{
		"url","image","name","price",
	}

	writer.Write(headers)

	for _,pokemonProduct := range pokemonProducts {
		record := []string{
			pokemonProduct.url,
			pokemonProduct.image,
			pokemonProduct.name,
			pokemonProduct.price,
		}
		writer.Write(record)
	}
defer writer.Flush()


}