package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make(Out)
	pipeline := func(done In, incStream In, stage Stage) Out {
		stageStream := make(Bi)
		go func() {
			defer close(stageStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-incStream:
					if !ok {
						return
					}
					stageStream <- doStage(v, stage)
				}

			}
		}()
		return stageStream
	}

	for s := range stages {
		if s == 0 {
			out = pipeline(done, in, stages[s])
			continue
		}
		out = pipeline(done, out, stages[s])
	}
	return out
}

func doStage(value interface{}, stage Stage) interface{} {
	out := make(Bi)
	go func() {
		defer close(out)
		out <- value
	}()
	return <-stage(out)
}
