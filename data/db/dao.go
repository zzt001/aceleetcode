package db

import (
	"github.com/zzt001/aceleetcode/data/crawler"
	gorp "gopkg.in/gorp.v1"
)

func UpdateProblemTable(dbMap *gorp.DbMap) error {
	dto, err := crawler.CrawlData()
	if err != nil {
		return err
	}
	pbs, err := dto.ToProblems()
	if err != nil {
		return err
	}
	// drop table first and then create again
	if err := dropProblemTable(dbMap); err != nil {
		return err
	}
	if err := createProblemTable(dbMap); err != nil {
		return err
	}
	// insert new data
	tx, err := dbMap.Db.Begin()
	if err != nil {
		return err
	}
	for _, pb := range pbs {
		if err := dbMap.Insert(pb); err != nil {
			return err
		}
	}
	return tx.Commit()
}
