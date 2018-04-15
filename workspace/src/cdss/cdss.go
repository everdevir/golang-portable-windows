// hx
package main

import (
	"log"
	"net/http"
	"time"
	//"github.com/remizovm/geonames"
)

type Visit struct {
	ID           uintptr
	Admission    *Admission
	Practitioner *Practitioner
	
  Place        Geo
  Time         time.Time
}

//TODO `Admission` must be renamed to something meaningful for both outpatient and inpatient care
type Admission struct {
	ID            uintptr
  AdmittingPhysician *Practitioner
  Patient       struct {
    *Person
    // Immutable //TODO what to do with immutable traits in order not to ask repeatedly?
    Ethnicity     string //TODO List its different kinds in a database document
	  Birth struct {
		  Date  time.Time
		  Place Geo
	  }
    // Mutable
    Age           uint8
    AgeGroup      uint8  //TODO List its different kinds in a database document
    Occupation    string //TODO List its different kinds in a database document
    MaritalStatus uint8  //TODO List its different kinds in a database document
  }
	Time struct {
		Start time.Time
		End   time.Time
	}
  After  *Admission // because of has: which admission is after this admission?
	Before *Admission // because of has: which admission is before this admission?
  Place struct {
		Hospital string
		Section  string
	}
  Visits []*Visit
  MedicalHistory {
    QuotedChiefComplaint string
    PresentingProblems   []string //TODO List its different kinds in a database document. Perhaps `interface` type would be better to save both int,string items!
  }
}

type title int

const (
	Mx title = iota + 1
	Mstr
	Mr
	Miss
	Ms
	Dr
	Prof

	Master    = Mstr
	Doctor    = Dr
	Professor = Prof

	allTitle = Mx | Mstr | Mr | Miss | Ms | Dr | Prof
)

type sex int

const (
	Male sex = iota + 1
	Female
  Unknown
  
	allSex = Unknown | Male | Female
)

type marital_status int

const (
	Married marital_status = iota + 1
	Divorced
	Single
)

type Geo struct {
	City      int
	Town      int
	State     int
	Country   int
	Continent int
}

type Address struct {
	HouseNumber int
	HouseName   string
	Street      string
	Locality    string
	Geo         Geo
	Postcode    int
	Telephone   int
}

type Person struct {
	ID uintptr

	Name struct {
		First  string
		Family string
		Maiden string
		Middle string
	}
	Sex   sex
	Title struct {
		Common  title
		Special string
	}

	/* the KEY below is a GeoID of a political district
	that provides National identifiction number */
	Citizenship map[int]struct {
		NationalID  string // Maybe some countries have alphabet characters in their national id numbers!
		HomeAddress []Address
	}
	Contact struct {
		Telephone struct {
			Home   []uint //TODO should be directed to Citizenship[].HomeAddress[].Telephone
			Work   []uint
			Mobile []uint
		}
		Email []string
	}
}

type Practitioner struct {
	Person

	/* the KEY below is a GeoID of a political district
	that provides National identifiction number */
	Physicianship map[int]struct {
		PhysicianID uint
		Office      []Address
	}

	Studentship []struct {
		StudentID  uint
		University string
		DateRange  string
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":8181", nil))
}
