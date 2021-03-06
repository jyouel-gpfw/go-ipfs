package lock

import (
	"io"
	"path"

	lock "github.com/ipfs/go-ipfs/Godeps/_workspace/src/github.com/camlistore/lock"
	"github.com/ipfs/go-ipfs/util"
	"github.com/ipfs/go-ipfs/util/debugerror"
)

// LockFile is the filename of the daemon lock, relative to config dir
// TODO rename repo lock and hide name
const LockFile = "daemon.lock"

func Lock(confdir string) (io.Closer, error) {
	c, err := lock.Lock(path.Join(confdir, LockFile))
	return c, debugerror.Wrap(err)
}

func Locked(confdir string) bool {
	if !util.FileExists(path.Join(confdir, LockFile)) {
		return false
	}
	if lk, err := Lock(confdir); err != nil {
		return true
	} else {
		lk.Close()
		return false
	}
}
