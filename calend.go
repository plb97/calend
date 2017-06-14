// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package calend

import (
	"fmt"
	"github.com/plb97/fqa"
)

/*
Lundi, mardi, mercredi, jeudi, vendredi, samedi et dimanche sont les sept jours de la semaine.
Lunae dies, martis dies, mercurii dies, iovis dies, veneris dies, saturni dies, solis dies sunt septem dies septimanae.
*/
/*
Janvier, février, mars, avril, mai, juin, juillet, août, septembre, octobre, novembre et décembre sont les douze mois de l'année.
Ianuarius, Februarius, Martius, Aprilis, Maius, Iunius, Iulius, Augustus, September, October, November, December sunt duodecim menses anni.
*/
/*   mois musulmans
	1	Ianuarius	janvier		Mouharram		محرم
	2	Februarius	février		Safar			صفر
	3	Martius		mars		Rabi' Al-Awwal		ربيع الأول
	4	Aprilis		avril		Rabi' Al-Thani		ربيع الثاني
	5	Maius		mai		Joumada Al-Awwal	جمادى الأول
	6	Iunius		juin		Joumada Al-Thani	جمادى الثاني
	7	Iulius		juillet		Rajab			-رجب
	8	Augustus	août		Cha'ban			شعبان
	9	September	septembre	Ramadan			رمضان
	10	October		octobre		Chawwal			شوال
	11	November	novembre	Dhou Al-qa'da		ذو القعدة
	12	December	décembre	Dhou Al-Hijja		ذو الحجة
*/
/*	jours de la semaine
	0	Lunae dies	lundi		elethnayn	الإثنين
	1	martis dies	mardi		eththolatha	الثلاثاء
	2	mercurii dies	mercredi	elarbieaa	الأربعاء
	3	iovis dies	jeudi		elkhamiis	الخميس
	4	veneris dies	vendredi	eljomoa		الجمعة
	5	saturni dies	samedi		essabat		السبت
	6	solis dies	dimanche	elahad		الأحد
*/

// Enumeration des jours de la semaine
type Jours_e int
const (

	LUNAE_DIES Jours_e = iota
	MARTIS_DIES
	MERCURII_DIES
	IOVIS_DIES
	VENERIS_DIES
	SATURNI_DIES
	SOLIS_DIES
	
	LUNDI = LUNAE_DIES
	MARDI = MARTIS_DIES
	MERCREDI = MERCURII_DIES
	JEUDI = IOVIS_DIES
	VENDREDI = VENERIS_DIES
	SAMEDI = SATURNI_DIES
	DIMANCHE = SOLIS_DIES

	MONDAY = LUNAE_DIES
	TUESDAY = MARTIS_DIES
	WEDNESDAY = MERCURII_DIES
	THURSDAY = IOVIS_DIES
	FRIDAY = VENERIS_DIES
	SATURDAY = SATURNI_DIES
	SUNDAY = SOLIS_DIES

	ELETHNAYN = LUNAE_DIES		// الإثنين
	ETHTHOLATHA = MARTIS_DIES	// الثلاثاء
	ELARBIEAA = MERCURII_DIES	// الأربعاء
	ELKHAMIIS = IOVIS_DIES		// الخميس
	ELJOMOA = VENERIS_DIES		// الجمعة
	ESSABAT = SATURNI_DIES		// السبت
	ELAHAD = SOLIS_DIES		// الأحد

)

// Enumeration des mois de l'annee
type Mois_e int
const (

	IANUARIUS Mois_e = iota + 1
	FEBRUARIUS
	MARTIUS
	APRILIS
	MAIUS
	IUNIUS
	IULIUS
	AUGUSTUS
	SEPTEMBER
	OCTOBER
	NOVEMBER
	DECEMBER

	JANVIER = IANUARIUS
	FEVRIER = FEBRUARIUS
	MARS = MARTIUS
	AVRIL = APRILIS
	MAI = MAIUS
	JUIN = IUNIUS
	JUILLET = IULIUS
	AOUT = AUGUSTUS
	SEPTEMBRE = SEPTEMBER
	OCTOBRE = OCTOBER
	NOVEMBRE = NOVEMBER
	DECEMBRE = DECEMBER

	JANUARY = IANUARIUS
	FEBRUARY = FEBRUARIUS
	MARCH = MARTIUS
	APRIL = APRILIS
	MAY = MAIUS
	JUNE = IUNIUS
	JULY = IULIUS
	AUGUST = AUGUSTUS
	//SEPTEMBER = SEPTEMBER
	//OCTOBER = OCTOBER
	//NOVEMBER = NOVEMBER
	//DECEMBER = DECEMBER

	MOUHARRAM = IANUARIUS		// محرم
	SAFAR = FEBRUARIUS		// صفر
	RABI_AL_AWWAL = MARTIUS		// ربيع الأول
	RABI_AL_THANI = APRILIS		// ربيع الثاني
	JOUMADA_AL_AWWAL = MAIUS	// جمادى الأول
	JOUMADA_AL_THANI = IUNIUS	// جمادى الثاني
	RAJAB = IULIUS			// رجب
	CHA_BAN = AUGUSTUS		// شعبان
	RAMADAN = SEPTEMBER		// رمضان
	CHAWWAL = OCTOBER		// شوال
	DHOUL_AL_QA_DA = NOVEMBER	// ذو القعدة
	DHOUL_AL_HIJJA = DECEMBER	// ذو الحجة

)

// Type date sous la forme d'un compteur de jours 'n' et d'une fraction de jour 't'
type JJ_t struct {
	n int
	t float64
}
// creation d'une date
func Creer(n int, t float64) *JJ_t {
	j, t := fqa.Ent(t)
	return &JJ_t{n:n+j, t:t}
}
// elements d'une date ('r.n' compteur de jours et 'r.t' fraction de jour ou heure)
// 0 <= 'r.t' < 1
func (r *JJ_t) Elmt() (int, float64) {
	return r.n, r.t
}
// date sous forme de chaine
func (r JJ_t) String() string {
	return fmt.Sprintf("[n:%d t:%f]",r.n,r.t)
}
// verification de l'egalite avec la date
func (r *JJ_t) Egal(q *JJ_t) bool {
	if nil == q { return nil == r}
	return r.n == q.n && r.t == q.t
}
func (r *JJ_t) Jour() Jours_e {
	return Jours_e(r.n % 7)
}
// structure calendrier
type calend_t struct {
	f []*fqa.Fqa_t		// formes quasi affines
}
// calendrier sous forne d'une chaine
func (o calend_t) String() string {
	s := fmt.Sprintf("f:[")
	for i := 0; i < len(o.f); i++ {
		if 0 == i {
			s += fmt.Sprintf("%d:%v",i,o.f[i])
		} else {
			s += fmt.Sprintf(" %d:%v",i,o.f[i])
		}
	}
	s += "]"
	return s
}
// jour julien d'une date sous forme de tableau d'entiers
// exemple de date : []int{20,17,5,1} == 1er mai 2017
func (o *calend_t) jj(dt []int) int {
	m := len(o.f) // nombre de formes quasi affines du calendrier
	if m != len(dt) {
		panic("Parametre invalide")
	}
	var v int = 0 // calcul du jour julien a l'aide des formes quasi affines
	for i := 0; i < m; i++ {
		v += o.f[i].Valeur(dt[i]) // calcul de de la valeur a l'aide de la forme quasi affine
	}
	return v
}
// date sous forme de tableau d'entiers a partir d'un jour julien
// exemple de date : []int{20,17,5,1} == 1er mai 2017
func (o *calend_t) dt(jj int) []int {
	m := len(o.f) // nombre de formes quasi affines du calendrier
	v := make([]int,m) // initailisation du tableau
	r := jj // initialisation du reste
	for i := 0; i < m; i++ {
		v[i], r = o.f[i].Div_fqa(r) // calcul du 'quotient' et du 'reste' de la 'division' par la forme quasi affine
	}
	return v
}

var g_calend = calend_t {	// calendrier gregorien
	f:[]*fqa.Fqa_t{
		fqa.Creer(146097,4,6884480),
		fqa.Creer(1461,4,0),
		fqa.Creer(153,5,-457),
		fqa.Creer(1,1,-1),
	},
}

var j_calend = calend_t {	// calendier julien
	f:[]*fqa.Fqa_t{
		fqa.Creer(1461,4,6884472),
		fqa.Creer(153,5,-457),
		fqa.Creer(1,1,-1),
	},
}

var m_calend = calend_t {	// calendier musulman
	f:[]*fqa.Fqa_t{
		fqa.Creer(10631,30,58442583),
		fqa.Creer(325,11,-320),
		fqa.Creer(1,1,-1),
	},
}

// Calendrier gregorien
func Gamj2jj(a, m int, d float64) *JJ_t {
	a, m = fqa.Norm_am(fqa.Corrig_am(a,m))
	j, t := fqa.Ent(d)
	c, b := fqa.Divent(a,100)
	return &JJ_t{g_calend.jj([]int{c,b,m,j}), t}
}
func Gjj2amj(jj *JJ_t) (int, int, float64) {
	dt := g_calend.dt(jj.n)
	a,m := fqa.DNorm_am(100*dt[0]+dt[1],dt[2])
	j := dt[3]
	return a,m,float64(j) + jj.t
}

// Calendrier julien
func Jamj2jj(a, m int, d float64) *JJ_t {
	a, m = fqa.Norm_am(fqa.Corrig_am(a,m))
	j, t := fqa.Ent(d)
	return &JJ_t{j_calend.jj([]int{a,m,j}), t}
}
func Jjj2amj(jj *JJ_t) (int, int, float64) {
	dt := j_calend.dt(jj.n)
	a,m := fqa.DNorm_am(dt[0],dt[1])
	j := dt[2]
	return a,m,float64(j) + jj.t
}

// Calendrier musulman
func Mamj2jj(a, m int, d float64) *JJ_t {
	a, m = fqa.Corrig_am(a,m)
	j, t := fqa.Ent(d)
	return &JJ_t{m_calend.jj([]int{a,m,j}), t}
}
func Mjj2amj(jj *JJ_t) (int, int, float64) {
	dt := m_calend.dt(jj.n)
	a,m,j := dt[0],dt[1],dt[2]
	return a,m,float64(j) + jj.t
}

