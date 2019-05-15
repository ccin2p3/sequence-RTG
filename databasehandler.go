package sequence

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/sqlboiler/boil"
	"sequence/models"
	"time"
)

func OpenDbandSetContext()(*sql.DB, context.Context){
	// Get a handle to the SQLite database, using mattn/go-sqlite3
	db, err := sql.Open("sqlite3", "sequence.sdb")
	if err != nil{
		panic(err)
	}
	// Configure SQLBoiler to use the sqlite database
	boil.SetDB(db)
	// Need to set a context for purposes I don't understand yet
	ctx := context.Background()     // Dark voodoo magic, https://golang.org/pkg/context/#Background
	return db, ctx
}

func GetPatternsFromDatabase(db *sql.DB, ctx context.Context) map[string]string{
	pmap := make(map[string]string)
	// This pulls 'all' of the patterns from the patterns database
	patterns, err := models.Patterns().All(ctx, db)
	if err !=nil {
		logger.DatabaseSelectFailed("patterns", "All", err.Error())
	}
	for _, p := range patterns{
		pmap[p.ID] = p.SequencePattern
	}
	return pmap
}

func GetPatternsWithExamplesFromDatabase(db *sql.DB, ctx context.Context) map[string]AnalyzerResult{
	pmap := make(map[string]AnalyzerResult)
	var patterns models.PatternSlice
	var err error
	if config.matchThresholdValue != "0"{
		patterns, err = models.Patterns(models.PatternWhere.ThresholdReached.EQ(true)).All(ctx, db)
		if err !=nil {
			logger.DatabaseSelectFailed("patterns", "Where threshold_reached=true", err.Error())
		}
	}else{
		patterns, err = models.Patterns().All(ctx, db)
		if err !=nil {
			logger.DatabaseSelectFailed("patterns", "All", err.Error())
		}
	}

	for _, p := range patterns{
		var s *models.Service
		s, err = models.Services(models.ServiceWhere.ID.EQ(p.ServiceID)).One(ctx,db)
		if err !=nil {
			logger.DatabaseSelectFailed("services", "Where id = " + p.ServiceID, err.Error())
		}
		ar := AnalyzerResult{PatternId:p.ID, Pattern:p.SequencePattern, ThresholdReached:p.ThresholdReached, DateCreated:p.DateCreated}
		var ex models.ExampleSlice
		ex, err = p.PatternExamples().All(ctx, db)
		if err !=nil {
			logger.DatabaseSelectFailed("examples", "All", err.Error())
		}
		for _, e := range ex{
			lr := LogRecord{Message:e.ExampleDetail, Service:s.Name}
			ar.Examples = append(ar.Examples, lr)
		}
		st, er := p.PatternStatistics().One(ctx, db)
		if er !=nil {
			logger.DatabaseSelectFailed("statistics", "One", err.Error())
		}
		ar.ExampleCount = int(st.OriginalMatchCount)
		pmap[p.ID]=ar
	}
	return pmap
}

func GetPatternsFromDatabaseByService(db *sql.DB, ctx context.Context, sid string) map[string]string{
	pmap := make(map[string]string)
	// This pulls 'all' of the patterns from the patterns database
	patterns, err := models.Patterns(models.PatternWhere.ServiceID.EQ(sid)).All(ctx, db)
	if err !=nil {
		logger.DatabaseSelectFailed("patterns", "Where Serviceid = " + sid, err.Error())
	}
	for _, p := range patterns{
		pmap[p.ID] = p.SequencePattern
	}
	return pmap
}

func GetServicesFromDatabase(db *sql.DB, ctx context.Context) map[string]string{
	// This pulls 'all' of the services from the services table
	smap := make(map[string]string)
	services, err := models.Services().All(ctx, db)
	if err !=nil {
		logger.DatabaseSelectFailed("services", "All", err.Error())
	}
	for _, p := range services{
		smap[p.ID] = p.Name
	}
	return smap
}

func AddService(ctx context.Context, tx *sql.Tx, id string, name string){
	// This pulls 'all' of the services from the services table
	var s models.Service
	s.ID = id
	s.Name = name
	s.DateCreated = time.Now()
	err := s.Insert(ctx, tx, boil.Whitelist("id", "name", "date_created"))
	if err != nil{
		logger.DatabaseInsertFailed("service", id, err.Error())
	}
}

func AddPattern(ctx context.Context, tx *sql.Tx, result AnalyzerResult, sID string){
	p := models.Pattern{ID:result.PatternId, SequencePattern:result.Pattern, DateCreated:time.Now(),ServiceID:sID, ThresholdReached:result.ThresholdReached}
	err := p.Insert(ctx, tx, boil.Whitelist("id", "sequence_pattern", "date_created", "threshold_reached", "service_id"))
	if err != nil{
		logger.DatabaseInsertFailed("pattern", result.PatternId, err.Error())
	}

	//add all examples if threshold has been reached as these will have already been pruned
	//otherwise limit to max three.
	if result.ThresholdReached{
		for _, e := range result.Examples{
			ex := models.Example{ExampleDetail:e.Message, PatternID:result.PatternId}
			err = ex.Insert(ctx, tx, boil.Infer())
			if err != nil{
				logger.DatabaseInsertFailed("example", result.PatternId, err.Error())
			}
		}
	}else{
		prev := ""
		count := 0
		for _, e := range result.Examples{
			if prev != e.Message{
				ex := models.Example{ExampleDetail:e.Message, PatternID:result.PatternId}
				err = ex.Insert(ctx, tx, boil.Infer())
				if err != nil{
					logger.DatabaseInsertFailed("example", result.PatternId, err.Error())
				}
				prev = e.Message
				count++
			}
			if count >= 3{
				break
			}
		}
	}
	//add initial statistics
	st := models.Statistic{PatternID:result.PatternId, CumulativeMatchCount:int64(result.ExampleCount), OriginalMatchCount:int64(result.ExampleCount), DateLastMatched:time.Now()}
	err = st.Insert(ctx, tx, boil.Infer())
	if err != nil{
		logger.DatabaseInsertFailed("statistics", result.PatternId, err.Error())
	}
}