package main

/*
Golang Fragments Example
http://www.gofragments.net/client/blog/concurrency/2016/01/28/findDuplicatedFiles/index.html
*/

import (
    "crypto/sha1"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "runtime"
    "sort"
    "sync"
    "time"
)

const maxSizeOfSmallFile = 1024 * 32
const maxGoroutines = 100

type pathsInfo struct {
    size  int64
    paths []string
}
type fileInfo struct {
    sha1 []byte
    size int64
    path string
}

func main() {
    // Use all the machine's cores
    // runtime.GOMAXPROCS(runtime.NumCPU()) // no longer needed since Go1.5
    /* alternative: using the command line if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
       fmt.Printf("usage: %s <path>\n", filepath.Base(os.Args[0]))
       os.Exit(1)
   }
   */
    defer timeTrack(time.Now(), "task duration:")
    // buffered channel, faster.
    infoChan := make(chan fileInfo, maxGoroutines*2)
    // alternative: go findDuplicates(infoChan, os.Args[1])
    findDuplicates(infoChan, "./")
    pathData := mergeResults(infoChan)
    outputResults(pathData)
}
func findDuplicates(infoChan chan fileInfo, dirname string) {
    // to make sure that we wait until they have all finished
    // before the findDuplicates() function returns let's use sync.WaitGroup().
    waiter := &sync.WaitGroup{}
    // makeWalkFunc is launched for each file in dirname
    filepath.Walk(dirname, makeWalkFunc(infoChan, waiter))
    waiter.Wait()
    close(infoChan)
}

// a factory creating funcs of type:
// "WalkFunc func(path string, info os.FileInfo, err error) error"
// to schedule/dispatch the files processing
func makeWalkFunc(infoChan chan fileInfo,
    waiter *sync.WaitGroup) func(string, os.FileInfo, error) error {
    // the closure will processFile if file is not large,
    // else will start a new goroutine
    return func(path string, info os.FileInfo, err error) error {
        if err == nil && info.Size() > 0 && (info.Mode()&os.ModeType == 0) {
            if info.Size() < maxSizeOfSmallFile ||
                runtime.NumGoroutine() > maxGoroutines {
                processFile(path, info, infoChan, nil)
            } else {
                // for each large files a goroutine is started,
                // waiter is increased
                waiter.Add(1)
                go processFile(path, info, infoChan, func() { waiter.Done() })
            }
        }
        return nil
    }
}

// collecting info about each file
func processFile(filename string, info os.FileInfo,
    infoChan chan fileInfo, done func()) {
    if done != nil {
        defer done()
    }
    file, err := os.Open(filename)
    if err != nil {
        log.Println("error:", err)
        return
    }
    defer file.Close()
    hash := sha1.New()
    if size, err := io.Copy(hash, file); size != info.Size() || err != nil {
        if err != nil {
            log.Println("error:", err)
        } else {
            log.Println("error: failed to read the whole file:", filename)
        }
        return
    }
    infoChan <- fileInfo{hash.Sum(nil), info.Size(), filename}
}

// gathering data from each processed file into
func mergeResults(infoChan <-chan fileInfo) map[string]*pathsInfo {
    pathData := make(map[string]*pathsInfo)
    /*     To produce the keys we have created a format string which has 16
           zero-padded hexadecimal digits to represent the file’s size and
           enough hexadecimal digits to represent the file’s SHA-1.
           We have used leading zeros for the file size part of the
           key so that we can sort the keys by size later on. The sha1.Size
           constant holds the number of bytes occupied by an SHA-1 (i.e., 20).
           Since one byte represented in hexadecimal has two digits,
           we must use twice as many characters as there are bytes for the SHA-1
           in the format string.
           (Incidentally,we could have created the format string with
           format = "%016X:%" + fmt.Sprintf("%dX", sha1.Size*2).)
   */
    format := fmt.Sprintf("%%016X:%%%dX", sha1.Size*2) // == "%016X:%40X"
    // range over the channel 'infoChan'
    for info := range infoChan {
        // signature: func Sprintf(format string, a ...interface{}) string
        key := fmt.Sprintf(format, info.size, info.sha1)
        value, found := pathData[key]
        if !found {
            value = &pathsInfo{size: info.size}
            pathData[key] = value
        }
        value.paths = append(value.paths, info.path)
    }
    return pathData
}
func outputResults(pathData map[string]*pathsInfo) {
    // a slice of strings, length, capacity
    keys := make([]string, 0, len(pathData))
    for key := range pathData {
        keys = append(keys, key)
    }
    sort.Strings(keys)
    for k, key := range keys {
        groupOfDuplicates := pathData[key]
        fmt.Printf("\n----\nGroup[%d] is = %#v\n", k, groupOfDuplicates.paths)
        if len(groupOfDuplicates.paths) > 1 {
            fmt.Printf("%d duplicated files (%s bytes):\n",
                len(groupOfDuplicates.paths), commas(groupOfDuplicates.size))
            sort.Strings(groupOfDuplicates.paths)
            for _, name := range groupOfDuplicates.paths {
                fmt.Printf("\t%s\n", name)
            }
        }
    }
}

// commas() returns a string representing the whole number with comma
// grouping.
func commas(x int64) string {
    value := fmt.Sprint(x)
    for i := len(value) - 3; i > 0; i -= 3 {
        value = value[:i] + "," + value[i:]
    }
    return value
}
func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    // fmt.Sprintf("%s%03d", t.Format("20060102150405"), t.Nanosecond()/1e6)
    fmt.Printf("\n----\nfunction %s took %s\n", name, elapsed)
}

/* Expected Output
----
Group[0] is = []string{"text-1.txt", "text-2.txt", "text-5.txt"}
3 duplicated files (12 bytes):
   text-1.txt
   text-2.txt
   text-5.txt
----
Group[1] is = []string{"text-3.txt", "text-4.txt"}
2 duplicated files (20 bytes):
   text-3.txt
   text-4.txt
----
Group[2] is = []string{"findduplicates.go"}
----
Group[3] is = []string{"findduplicates.exe"}
----
function task duration: took 9.0005ms
*/
