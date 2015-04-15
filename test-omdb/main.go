/*
Copyright 2014 Kaissersoft Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"github.com/caerostris/go-omdb"
	"log"
)

func main() {
	//Testing SearchMovies
	res, err := omdb.SearchMovies("The fifth element", "")
	if err != nil {
		log.Fatal(err)
	}

	//Testing GetMovieByTitle
	res2, err := omdb.GetMovieByTitle("True Grit", "1969", "full", "false")
	if err != nil {
		log.Fatal(err)
	}

	//Testing GetMovieByImdbId
	res3, err := omdb.GetMovieByImdbID("tt2015381", "short", "true")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Search[0].Title)
	fmt.Println(res2.Title)
	fmt.Println(res3.Title)
}
