package observe

import (
	"log"
	"net/http"
	"net/http/pprof"
	"runtime"
)

func ApplyRuntimeProfiling(blockProfileRate int, mutexProfileFraction int) {
	if blockProfileRate > 0 {
		runtime.SetBlockProfileRate(blockProfileRate)
		log.Printf("[Pprof] block profile enabled, rate=%d", blockProfileRate)
	}
	if mutexProfileFraction > 0 {
		runtime.SetMutexProfileFraction(mutexProfileFraction)
		log.Printf("[Pprof] mutex profile enabled, fraction=%d", mutexProfileFraction)
	}
}
func NewPprofServer(addr string) *http.Server {
	mux := http.NewServeMux()
	// 标准 pprof
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	mux.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
	mux.Handle("/debug/pprof/block", pprof.Handler("block"))
	mux.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))
	mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
