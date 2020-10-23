package logger

import (
"go.uber.org/zap"
)

type loggerWrapper struct {
	lw *zap.SugaredLogger
}

func (logger *loggerWrapper) Errorf(format string, args ...interface{}) {
	logger.lw.Errorf(format, args)
}
func (logger *loggerWrapper) Fatalf(format string, args ...interface{}) {
	logger.lw.Fatalf(format, args)
}
func (logger *loggerWrapper) Fatal(args ...interface{}) {
	logger.lw.Fatal(args)
}
func (logger *loggerWrapper) Infof(format string, args ...interface{}) {
	logger.lw.Infof(format, args)
}
func (logger *loggerWrapper) Warnf(format string, args ...interface{}) {
	logger.lw.Warnf(format, args)
}
func (logger *loggerWrapper) Debugf(format string, args ...interface{}) {
	logger.lw.Debugf(format, args)
}
func (logger *loggerWrapper) Printf(format string, args ...interface{}) {
	logger.lw.Infof(format, args)
}
func (logger *loggerWrapper) Println(args ...interface{}) {
	logger.lw.Info(args, "\n")
}



// initLog create logger
func InitLog()  {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	Sugarlog := logger.Sugar()
	Sugarlog.Info()
	SetLogger(Sugarlog)
}

