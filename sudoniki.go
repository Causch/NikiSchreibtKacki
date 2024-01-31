package main

import (
  "fmt"
  "math/rand"
  "time"
  "strings"
)

var prefixes         = []string{"WAHNSINN", "SENSATION", "BOMBE", "UNFASSBAR", "BOMBE PLATZT", "ESKALATION","ENDLICH","UNFASSBARE ENTHÜLLUNGEN", "EILMELDUNG","ES KNALLT","ABSTURZ", "DEMÜTIGUNG"}
var topics           = []string{ 
                                "%s: %s zerstört %s", 
                                "%s: %s gewinnt - GIGANTISCHE STRAFE fuer %s", 
                                "%s: %s schafft unfassbares! %s heult", 
                                "%s: %s SIEGT UND %s MUSS WEINEN", 
                                "%s: %s BOXT %s WEG",
                                "%s: %s VERNICHTET %s",
                                "%s: %s DROHT UND %s BRICHT ZUSAMMEN",
                                "%s: %s BEENDET %ss KARRIERE", 
                                "%s: %s GEWINNT UND %s RASTET AUS", 
                                "%s: %s ZÜNDET BOMBE UND %s KRIEGT ANGST",
                                "%s: %s GEWINNT UND DECKT %ss BETRUG AUF",
                              }
var gute_leute       =[]string{ "AfD","Weidel", "Höcke", "Rene Springer",}
var schlechte_leute  =[]string{"Kanzler","Kanzler Scholz","Habeck","Gruene Politikerin","Scholz","ZDF","ARD", "GEZ","Ampel","SPD","Drosten","RKI-Chef","Lauterbach"}

func generateTitle() string {
  rand.Seed(time.Now().UnixNano())

  prefix := prefixes[rand.Intn(len(prefixes))]
  topic := topics[rand.Intn(len(topics))]

  gut      := gute_leute      [rand.Intn(len(gute_leute))]
  schlecht := schlechte_leute [rand.Intn(len(schlechte_leute))]

  return fmt.Sprintf(topic, prefix, gut,schlecht)
}

func main() {
  for i := 0; i < 5; i++ {
    title := generateTitle()
    fmt.Printf("%s\n\n",strings.ToUpper(title))
  }
}