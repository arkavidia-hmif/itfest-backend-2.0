package types

import (
	"context"
	"database/sql/driver"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"regexp"
	"strings"
)

type CareerInterest string

const (
	SoftwareEngineering  CareerInterest = "software-engineering"
	ProductManagement    CareerInterest = "product-management"
	UIDesigner           CareerInterest = "ui-designer"
	UXDesigner           CareerInterest = "ux-designer"
	UXResearcher         CareerInterest = "ux-researcher"
	ITConsultant         CareerInterest = "it-consultant"
	GameDeveloper        CareerInterest = "game-developer"
	CyberSecurity        CareerInterest = "cyber-security"
	BusinessAnalyst      CareerInterest = "business-analyst"
	BusinessIntelligence CareerInterest = "business-intelligence"
	DataScientist        CareerInterest = "data-scientist"
	DataAnalyst          CareerInterest = "data-analyst"
)

func (careerInterest *CareerInterest) Scan(value interface{}) error {
	*careerInterest = CareerInterest(value.(string))
	return nil
}

func (careerInterest CareerInterest) Value() (driver.Value, error) {
	return string(careerInterest), nil
}

func (CareerInterest) GormDataType() string {
	return "string"
}

type CareerInterests []CareerInterest

func (careerInterests *CareerInterests) Scan(values interface{}) error {
	regex, err := regexp.Compile(`[a-zA-Z\-]+`)
	if err != nil {
		return nil
	}

	words := regex.FindAllString(values.(string), -1)
	*careerInterests = []CareerInterest{}
	for _, word := range words {
		*careerInterests = append(*careerInterests, CareerInterest(word))
	}
	return nil
}

func (careerInterests CareerInterests) Value() (driver.Value, error) {
	var values []string
	for _, careerInterest := range careerInterests {
		values = append(values, string(careerInterest))
	}

	return values, nil
}

func (careerInterests CareerInterests) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	interests, _ := careerInterests.Value()

	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{strings.Join(interests.([]string), " ")},
	}
}

func (CareerInterests) GormDataType() string {
	return "string"
}
