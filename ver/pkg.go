package ver

import (
	"fmt"
	"log/slog"
	"runtime/debug"
	"slices"
	"strconv"
	"strings"
)

const Delimiters = "._-"

// BuildVersion - version of this binary, taken from the build info the Go toolchain
// stamps into every binary. A build from a tagged, clean checkout reports the
// tag, an untagged checkout a pseudo-version, and `go run` reports "(devel)".
func BuildVersion() (version string) {
	version = "(unknown)"
	if bi, ok := debug.ReadBuildInfo(); ok && bi.Main.Version != "" {
		version = bi.Main.Version
	}
	return version
}

// Newer - return v1 > v2
func Newer(v1, v2 string) (newer bool) { return segmentCompare(segmentSplit(v1), segmentSplit(v2)) }

// return s1 > s2
func segmentCompare(s1, s2 []string) (newer bool) {
	prefix := "SegmentCompare"
	var (
		e     error
		equal = true
		int1  int
		int2  int
		len1  = len(s1)
		len2  = len(s2)
	)
	for i := range min(len1, len2) {
		slog.Debug(prefix, fmt.Sprintf("s1[%d]", i), s1[i])
		slog.Debug(prefix, fmt.Sprintf("s2[%d]", i), s2[i])
		int1, e = strconv.Atoi(trimNonNumeric(s1[i]))
		if e == nil {
			int2, e = strconv.Atoi(trimNonNumeric(s2[i]))
		}
		if e == nil {
			equal = equal && int1 == int2
			newer = int1 > int2
		}
		if e != nil || !equal {
			break
		}
	}

	if e == nil && equal {
		newer = len1 > len2
	}
	return newer
}

// Trim leading 'v', split s by VerDelimiters
func segmentSplit(s string) (segments []string) {
	var (
		char = ""
	)
	s = strings.TrimPrefix(s, "v")
	s = strings.TrimPrefix(s, "V")
	segments = []string{}
	for l := range len(s) {
		char = string(s[l])
		if strings.Contains(Delimiters, char) {
			segments = append(segments, "")
		} else {
			if len(segments) == 0 {
				segments = append(segments, "")
			}
			segments[len(segments)-1] += char
		}
	}
	return segments
}

func trimNonNumeric(s string) string {
	prefix := "trimNonNumeric"
	var (
		start int
		end   int
	)
	// get start
	for i, c := range s {
		start = i
		if c >= '0' && c <= '9' {
			break
		}
	}
	// get end
	for i, c := range slices.Backward([]byte(s)) {
		end = i
		if c >= '0' && c <= '9' {
			break
		}
	}
	end++

	slog.Debug(prefix, "s", s, "start", start, "end", end)
	slog.Debug(prefix, "out", s[start:end])
	return s[start:end]
}
