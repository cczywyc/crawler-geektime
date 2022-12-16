package log

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestNewFilePlugin(t *testing.T) {
	const filePrefix = "test"
	const fileSuffix = ".log"
	const gzipSuffix = ".gz"

	p, c := NewFilePlugin(filePrefix+fileSuffix, zapcore.DebugLevel)
	logger := NewLogger(p)
	b := make([]byte, 10000)
	var count = 10000
	for count > 0 {
		count--
		logger.Info(string(b))
	}

	err := c.Close()
	require.NoError(t, err)
	time.Sleep(1 * time.Second)

	fs, err := ioutil.ReadDir(".")
	require.NoError(t, err)

	var gzCount, logCount int
	for _, f := range fs {
		name := f.Name()
		if strings.HasPrefix(name, filePrefix) {
			if strings.HasSuffix(name, fileSuffix) {
				logCount++
				assert.NoError(t, os.Remove(f.Name()))
				continue
			}
			if strings.HasSuffix(name, fileSuffix+gzipSuffix) {
				gzCount++
				logCount++
				assert.NoError(t, os.Remove(f.Name()))
				continue
			}
		}
	}
	require.Equal(t, 4, logCount)
	require.Equal(t, 2, gzCount)
}
