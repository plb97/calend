// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package calend

import (
	"testing"
)

type donnees_t struct {
	y, m int
	d float64
}
func (r donnees_t) equal(q donnees_t) bool {
	return r.y == q.y && r.m == q.m && r.d == q.d
}

type test_calend_t struct {
	jj  int // jour julien
	amj donnees_t // annee mois jour
	js  Jours_e // jour de la semaine
	dt  *JJ_t // date
}

// calendrier gregorien
var test_gcalend = []test_calend_t {
	{jj:2448449, amj:donnees_t{y:1991,m:7,d:11}, js:JEUDI,dt:Creer(2448449,0)},
	{jj:2451545, amj:donnees_t{y:2000,m:1,d:1}, js:SAMEDI,dt:Creer(2451545,0)},
	{jj:2457578, amj:donnees_t{y:2016,m:7,d:8}, js:VENDREDI,dt:Creer(2457578,0)},
	{jj:2299161, amj:donnees_t{y:1582,m:10,d:15}, js:VENDREDI,dt:Creer(2299161,0)},
}
func Test_Gamj2jj(t *testing.T) {
	for _, tt := range test_gcalend {
		attendu := tt.dt
		obtenu := Gamj2jj(tt.amj.y, tt.amj.m, tt.amj.d)
		if !obtenu.Egal(attendu) {
			t.Errorf("Gamj2jj: attendu %v, obtenu %v", attendu, obtenu)
		}
		if obtenu.Jour() != tt.js {
			t.Errorf("Gamd2jj: attendu js %d, obtenu %d", tt.js, obtenu.Jour())
		}
	}
}
func Test_Gjj2amj(t *testing.T) {
	for _, tt := range test_gcalend {
		attendu := tt.amj
		y, m, d := Gjj2amj(tt.dt)
		obtenu := donnees_t{y:y, m:m, d:d}
		if !obtenu.equal(attendu) {
			t.Errorf("Gjj2amj: attendu %v, obtenu %v", attendu, obtenu)
		}
	}
}

// calendrier julien
var test_jcalend = []test_calend_t {
	{jj:1842713, amj:donnees_t{y:333,m:1,d:27}, js:SATURDAY,dt:Creer(1842713,0)},
	{jj:2299160, amj:donnees_t{y:1582,m:10,d:4}, js:THURSDAY,dt:Creer(2299160,0)},
}
func Test_Jamj2jj(t *testing.T) {
	for _, tt := range test_jcalend {
		attendu := tt.dt
		obtenu := Jamj2jj(tt.amj.y, tt.amj.m, tt.amj.d)
		if !obtenu.Egal(attendu) {
			t.Errorf("Jamj2jj: attendu %v, obtenu %v", attendu, obtenu)
		}
		if obtenu.Jour() != tt.js {
			t.Errorf("Jamj2jj: attendu js %d, obtenu %d", tt.js, obtenu.Jour())
		}
	}
}
func Test_Jjj2amj(t *testing.T) {
	for _, tt := range test_jcalend {
		attendu := tt.amj
		y, m, d := Jjj2amj(tt.dt)
		obtenu := donnees_t{y:y, m:m, d:d}
		if !obtenu.equal(attendu) {
			t.Errorf("Jjj2amj: attendu %v, obtenu %v", attendu, obtenu)
		}
	}
}

// calendrier musulman
var test_mcalend = []test_calend_t {
	// samedi 1er janvier 2000
	{jj:2451545, amj:donnees_t{y:1420,m:int(RAMADAN),d:24}, js:ESSABAT,dt:Creer(2451545,0)},
	// vendredi 8 juillet 2016
	{jj:2457578, amj:donnees_t{y:1437,m:int(CHAWWAL),d:2}, js:ELJOMOA,dt:Creer(2457578,0)},
}
 func Test_Mamj2jj(t *testing.T) {
	for _, tt := range test_mcalend {
		attendu := tt.dt
		obtenu := Mamj2jj(tt.amj.y, tt.amj.m, tt.amj.d)
		if !obtenu.Egal(attendu) {
			t.Errorf("Mamj2jj: attendu %v, obtenu %v", attendu, obtenu)
		}
		if obtenu.Jour() != tt.js {
			t.Errorf("Mamj2jj: attendu js %d, obtenu %d", tt.js, obtenu.Jour())
		}
	}
}
func Test_Mjj2amj(t *testing.T) {
	for _, tt := range test_mcalend {
		attendu := tt.amj
		y, m, d := Mjj2amj(tt.dt)
		obtenu := donnees_t{y:y, m:m, d:d}
		if !obtenu.equal(attendu) {
			t.Errorf("Mjj2amj: attendu %v, obtenu %v", attendu, obtenu)
		}
	}
}
