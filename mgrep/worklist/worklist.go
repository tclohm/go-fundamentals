package worklist

// Entry into channel
type Entry struct {
	Path string
}

// channel with entry structure
type Worklist struct {
	jobs chan Entry
}

// place data onto channel
func (w *Worklist) Add(work Entry) {
	w.jobs <- work
}

// job out of channel
func (w *Worklist) Next() Entry {
	j := <-w.jobs
	return j
}

// generate new worklist, buffer channel
func New(bufSize int) Worklist {
	return Worklist{make(chan Entry, bufSize)}
}

// generate new job
func NewJob(path string) Entry {
	return Entry{path}
}

// empty job means time to quit the channel
// self terminate
func (w *Worklist) Finalize(numWorkers int) {
	for i := 0 ; i < numWorkers ; i++ {
		w.Add(Entry{""})
	}
}