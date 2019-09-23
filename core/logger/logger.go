package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

//Logger - global var
var Log *logrus.Logger

func Init() {
	// Create a new instance of the logger. You can have any number of instances.
	Log = logrus.New()

	// Log as JSON instead of the default ASCII formatter.
	Log.SetFormatter(&logrus.JSONFormatter{})

	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	Log.Out = os.Stdout
}
