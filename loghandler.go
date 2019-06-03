package sequence

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

//this log wrapper makes it easy to change logging library in one
//place if other golang logging libraries are preferred.

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger(fname string, level string) *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}
	if len(fname) == 0{
		fname = "sequence.log"
	}
	file, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		standardLogger.Out = file
	} else {
		standardLogger.HandleInfo("Failed to log to file, using default stderr")
	}

	//if set level for logging, default to info
	switch level{
	case "fatal":
		standardLogger.SetLevel(logrus.FatalLevel)
	case "error":
		standardLogger.SetLevel(logrus.ErrorLevel)
	case "debug":
		standardLogger.SetLevel(logrus.DebugLevel)
	case "trace":
		standardLogger.SetLevel(logrus.TraceLevel)
	default:
		standardLogger.SetLevel(logrus.InfoLevel)
	}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

// Declare variables to store log messages as new Events
var (
	eventGenericDebug   = Event{000, "%s"}
	eventGenericInfo    = Event{100, "%s"}
	eventAnalyzeInfo    = Event{101, "Analyzed %d messages, found %d unique patterns, %d are new. %d messages errored, time taken: %s"}
	eventOutputInfo     = Event{102, "Output %d patterns to file, the top 5 matched patterns are %s, time taken: %s"}
	eventGenericError   = Event{200, "%s"}
	eventAnalysisFailed = Event{201, "Unable to analyze: %s"}
	eventDbInsertFailed = Event{301, "Failed to insert record into %s table, id: %s, reason: %s"}
	eventDbUpdateFailed = Event{302, "Failed to update record in %s table, id: %s, reason: %s"}
	eventDbSelectFailed = Event{303, "Failed to select record(s) from %s table, query: %s, reason: %s"}
	eventGenericFatal   = Event{400, "Fatal error occurred, reason: %s"}
	eventGenericPanic   = Event{500, "Panic occurred, reason: %s"}
)

func (l *StandardLogger) LogAnalysisFailed(lr LogRecord){
	l.WithFields(logrus.Fields{
		"id": eventAnalysisFailed.id,
	}).Debugf(eventAnalysisFailed.message, lr.Message)
}

func (l *StandardLogger) DatabaseInsertFailed(tablename string, id string, reason string){
	l.WithFields(logrus.Fields{
		"id": eventDbInsertFailed.id,
	}).Errorf(eventDbInsertFailed.message, tablename, id, reason)
}

func (l *StandardLogger) DatabaseUpdateFailed(tablename string, query string, reason string){
	l.WithFields(logrus.Fields{
		"id": eventDbUpdateFailed.id,
	}).Errorf(eventDbUpdateFailed.message, tablename, query, reason)
}

func (l *StandardLogger) DatabaseSelectFailed(tablename string, query string, reason string){
	l.WithFields(logrus.Fields{
		"id": eventDbSelectFailed.id,
	}).Errorf(eventDbSelectFailed.message, tablename, query, reason)
}

func (l *StandardLogger) AnalyzeInfo(analyzedCount int, patternsCount int, new int, errCount int, taken time.Duration){
	l.WithFields(logrus.Fields{
		"id":           	eventAnalyzeInfo.id,
		"analyzed_msg": 	analyzedCount,
		"patterns_found":	patternsCount,
		"patterns_new":		new,
		"errored_msg":		errCount,
	}).Infof(eventAnalyzeInfo.message, analyzedCount, patternsCount, new, errCount, taken)
}

func (l *StandardLogger) OutputToFileInfo(outputCount int, top5 string, taken time.Duration){
	l.WithFields(logrus.Fields{
		"id":           	eventOutputInfo.id,
		"output_patterns": 	outputCount,
		"top_5":			top5,
	}).Infof(eventOutputInfo.message, outputCount, top5, taken)
}

func (l *StandardLogger) HandleFatal(err string){
	l.WithFields(logrus.Fields{
		"id": eventGenericFatal.id,
	}).Fatalf(eventGenericFatal.message, err)
}

func (l *StandardLogger) HandlePanic(err string){
	l.WithFields(logrus.Fields{
		"id": eventGenericPanic.id,
	}).Panicf(eventGenericPanic.message, err)
}

func (l *StandardLogger) HandleInfo(message string){
	l.WithFields(logrus.Fields{
		"id": eventGenericInfo.id,
	}).Infof(eventGenericInfo.message, message)
}

func (l *StandardLogger) HandleError(message string){
	l.WithFields(logrus.Fields{
		"id": eventGenericError.id,
	}).Errorf(eventGenericError.message, message)
}

func (l *StandardLogger) HandleDebug(message string){
	l.WithFields(logrus.Fields{
		"id": eventGenericDebug.id,
	}).Debugf(eventGenericDebug.message, message)
}

