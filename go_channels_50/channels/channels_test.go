package channels

import (
	"context"
	"reflect"
	"sort"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func collectWithTimeout[T any](t *testing.T, ch <-chan T) []T {
	t.Helper()
	if ch == nil {
		t.Fatalf("канал равен nil")
	}
	var result []T
	deadline := time.After(300 * time.Millisecond)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return result
			}
			result = append(result, v)
		case <-deadline:
			t.Fatalf("timeout при чтении канала; возможно канал не закрыт или горутина зависла")
		}
	}
}

func feedInts(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, n := range nums {
			ch <- n
		}
	}()
	return ch
}

func sameMultiset(a, b []int) bool {
	if len(a) != len(b) { return false }
	sort.Ints(a); sort.Ints(b)
	return reflect.DeepEqual(a, b)
}

func TestSendOne(t *testing.T) {
	ch := make(chan int, 1)
	SendOne(ch, 42)
	select {
	case got := <-ch:
		if got != 42 { t.Fatalf("SendOne sent %d, want 42", got) }
	default:
		t.Fatalf("SendOne не отправил значение")
	}
}

func TestReceiveOne(t *testing.T) {
	ch := make(chan int, 1); ch <- 7
	if got := ReceiveOne(ch); got != 7 { t.Fatalf("ReceiveOne() = %d, want 7", got) }
}

func TestCloseChannel(t *testing.T) {
	ch := make(chan int)
	CloseChannel(ch)
	select {
	case _, ok := <-ch:
		if ok { t.Fatalf("канал должен быть закрыт") }
	case <-time.After(300 * time.Millisecond):
		t.Fatalf("канал не закрыт")
	}
}

func TestSendSlice(t *testing.T) {
	got := collectWithTimeout(t, SendSlice([]int{1,2,3}))
	want := []int{1,2,3}
	if !reflect.DeepEqual(got, want) { t.Fatalf("SendSlice() = %v, want %v", got, want) }
}

func TestCollectChannel(t *testing.T) {
	got := CollectChannel(feedInts(1,2,3))
	want := []int{1,2,3}
	if !reflect.DeepEqual(got, want) { t.Fatalf("CollectChannel() = %v, want %v", got, want) }
}

func TestSumChannel(t *testing.T) {
	if got := SumChannel(feedInts(1,2,3,4)); got != 10 { t.Fatalf("SumChannel() = %d, want 10", got) }
}

func TestCountChannel(t *testing.T) {
	if got := CountChannel(feedInts(1,2,3)); got != 3 { t.Fatalf("CountChannel() = %d, want 3", got) }
}

func TestDrainChannel(t *testing.T) {
	ch := feedInts(1,2,3)
	done := make(chan struct{})
	go func(){ DrainChannel(ch); close(done) }()
	select { case <-done: case <-time.After(300*time.Millisecond): t.Fatalf("DrainChannel завис") }
}

func TestGenerateRange(t *testing.T) {
	got := collectWithTimeout(t, GenerateRange(3,6))
	want := []int{3,4,5,6}
	if !reflect.DeepEqual(got, want) { t.Fatalf("GenerateRange() = %v, want %v", got, want) }
}

func TestRepeatValue(t *testing.T) {
	got := collectWithTimeout(t, RepeatValue(5,3))
	want := []int{5,5,5}
	if !reflect.DeepEqual(got, want) { t.Fatalf("RepeatValue() = %v, want %v", got, want) }
}

func TestReceiveOrDefault(t *testing.T) {
	ch := make(chan int, 1); ch <- 9
	if got := ReceiveOrDefault(ch, 100); got != 9 { t.Fatalf("got %d, want 9", got) }
	if got := ReceiveOrDefault(ch, 100); got != 100 { t.Fatalf("got %d, want default 100", got) }
}

func TestTrySend(t *testing.T) {
	ch := make(chan int, 1)
	if !TrySend(ch, 10) { t.Fatalf("первый TrySend должен вернуть true") }
	if got := <-ch; got != 10 { t.Fatalf("sent %d, want 10", got) }
	ch <- 1
	if TrySend(ch, 2) { t.Fatalf("TrySend в полный канал должен вернуть false") }
}

func TestReceiveWithTimeout(t *testing.T) {
	ch := make(chan int, 1); ch <- 12
	got, ok := ReceiveWithTimeout(ch, time.Second)
	if !ok || got != 12 { t.Fatalf("got (%d,%v), want (12,true)", got, ok) }
	_, ok = ReceiveWithTimeout(ch, 20*time.Millisecond)
	if ok { t.Fatalf("ожидали timeout") }
}

func TestSendWithTimeout(t *testing.T) {
	ch := make(chan int, 1)
	if !SendWithTimeout(ch, 33, time.Second) { t.Fatalf("ожидали успешную отправку") }
	if got := <-ch; got != 33 { t.Fatalf("sent %d, want 33", got) }
	ch <- 1
	if SendWithTimeout(ch, 2, 20*time.Millisecond) { t.Fatalf("ожидали timeout на полном канале") }
}

func TestWaitDone(t *testing.T) {
	done := make(chan struct{})
	finished := make(chan struct{})
	go func(){ WaitDone(done); close(finished) }()
	select { case <-finished: t.Fatalf("WaitDone завершился до закрытия done") ; case <-time.After(20*time.Millisecond): }
	close(done)
	select { case <-finished: case <-time.After(300*time.Millisecond): t.Fatalf("WaitDone не завершился") }
}

func TestStopOnDone(t *testing.T) {
	done := make(chan struct{})
	in := make(chan int)
	go func(){ in <- 1; in <- 2; close(done); in <- 3 }()
	got := StopOnDone(done, in)
	if len(got) > 2 { t.Fatalf("StopOnDone прочитал слишком много: %v", got) }
}

func TestFirstReceived(t *testing.T) {
	a := make(chan int); b := make(chan int)
	go func(){ time.Sleep(30*time.Millisecond); a <- 1 }()
	go func(){ time.Sleep(5*time.Millisecond); b <- 2 }()
	if got := FirstReceived(a,b); got != 2 { t.Fatalf("FirstReceived() = %d, want 2", got) }
}

func TestFirstReceivedWithSource(t *testing.T) {
	a := make(chan int); b := make(chan int)
	go func(){ time.Sleep(5*time.Millisecond); a <- 10 }()
	go func(){ time.Sleep(30*time.Millisecond); b <- 20 }()
	got, src := FirstReceivedWithSource(a,b)
	if got != 10 || src != "a" { t.Fatalf("got (%d,%q), want (10,\"a\")", got, src) }
}

func TestTimeoutResult(t *testing.T) {
	got, ok := TimeoutResult(func() int { return 55 }, time.Second)
	if !ok || got != 55 { t.Fatalf("got (%d,%v), want (55,true)", got, ok) }
	_, ok = TimeoutResult(func() int { time.Sleep(80*time.Millisecond); return 1 }, 10*time.Millisecond)
	if ok { t.Fatalf("ожидали timeout") }
}

func TestRunWithTimeout(t *testing.T) {
	if !RunWithTimeout(func(){}, time.Second) { t.Fatalf("быстрая функция должна успеть") }
	if RunWithTimeout(func(){ time.Sleep(80*time.Millisecond) }, 10*time.Millisecond) { t.Fatalf("медленная функция не должна успеть") }
}

func TestMergeTwoChannels(t *testing.T) {
	got := collectWithTimeout(t, MergeTwoChannels(feedInts(1,2), feedInts(3,4)))
	if !sameMultiset(got, []int{1,2,3,4}) { t.Fatalf("MergeTwoChannels() = %v", got) }
}

func TestMergeManyChannels(t *testing.T) {
	chs := []<-chan int{feedInts(1,2), feedInts(3), feedInts(4,5)}
	got := collectWithTimeout(t, MergeManyChannels(chs))
	if !sameMultiset(got, []int{1,2,3,4,5}) { t.Fatalf("MergeManyChannels() = %v", got) }
}

func TestSplitEvenOdd(t *testing.T) {
	even, odd := SplitEvenOdd(feedInts(1,2,3,4))
	if !reflect.DeepEqual(collectWithTimeout(t, even), []int{2,4}) { t.Fatalf("even wrong") }
	if !reflect.DeepEqual(collectWithTimeout(t, odd), []int{1,3}) { t.Fatalf("odd wrong") }
}

func TestFanOut(t *testing.T) {
	outs := FanOut(feedInts(1,2,3,4,5), 3)
	if len(outs) != 3 { t.Fatalf("len outs = %d, want 3", len(outs)) }
	var all []int
	for _, out := range outs { all = append(all, collectWithTimeout(t, out)...)}
	if !sameMultiset(all, []int{1,2,3,4,5}) { t.Fatalf("FanOut values = %v", all) }
}

func TestBroadcast(t *testing.T) {
	outs := Broadcast(feedInts(1,2), 3)
	if len(outs) != 3 { t.Fatalf("len outs = %d, want 3", len(outs)) }
	for i,out := range outs {
		got := collectWithTimeout(t, out)
		if !reflect.DeepEqual(got, []int{1,2}) { t.Fatalf("out %d = %v", i, got) }
	}
}

func TestWorkerPoolSquare(t *testing.T) {
	got := collectWithTimeout(t, WorkerPoolSquare(feedInts(1,2,3), 2))
	if !sameMultiset(got, []int{1,4,9}) { t.Fatalf("WorkerPoolSquare() = %v", got) }
}

func TestWorkerPoolProcess(t *testing.T) {
	got := collectWithTimeout(t, WorkerPoolProcess(feedInts(1,2,3), 2, func(x int) int { return x+10 }))
	if !sameMultiset(got, []int{11,12,13}) { t.Fatalf("WorkerPoolProcess() = %v", got) }
}

func TestProcessConcurrently(t *testing.T) {
	got := ProcessConcurrently([]int{1,2,3}, func(x int) int { return x*x })
	want := []int{1,4,9}
	if !reflect.DeepEqual(got, want) { t.Fatalf("ProcessConcurrently() = %v, want %v", got, want) }
}

func TestLimitConcurrency(t *testing.T) {
	var current int32
	var maxSeen int32
	tasks := make([]func(), 20)
	for i := range tasks {
		tasks[i] = func(){
			v := atomic.AddInt32(&current, 1)
			for {
				m := atomic.LoadInt32(&maxSeen)
				if v <= m || atomic.CompareAndSwapInt32(&maxSeen, m, v) { break }
			}
			time.Sleep(10*time.Millisecond)
			atomic.AddInt32(&current, -1)
		}
	}
	LimitConcurrency(tasks, 3)
	if maxSeen > 3 { t.Fatalf("одновременно было %d задач, want <= 3", maxSeen) }
	if maxSeen == 0 { t.Fatalf("задачи не выполнились") }
}

func TestCountWorkersUsed(t *testing.T) {
	if got := CountWorkersUsed(feedInts(1,2,3,4), 2); got != 4 { t.Fatalf("CountWorkersUsed() = %d, want 4", got) }
}

func TestMapChannel(t *testing.T) {
	got := collectWithTimeout(t, MapChannel(feedInts(1,2,3), func(x int) int { return x*2 }))
	if !reflect.DeepEqual(got, []int{2,4,6}) { t.Fatalf("MapChannel() = %v", got) }
}

func TestFilterChannel(t *testing.T) {
	got := collectWithTimeout(t, FilterChannel(feedInts(1,2,3,4), func(x int) bool { return x%2 == 0 }))
	if !reflect.DeepEqual(got, []int{2,4}) { t.Fatalf("FilterChannel() = %v", got) }
}

func TestTakeN(t *testing.T) {
	got := collectWithTimeout(t, TakeN(feedInts(1,2,3,4), 2))
	if !reflect.DeepEqual(got, []int{1,2}) { t.Fatalf("TakeN() = %v", got) }
}

func TestSkipN(t *testing.T) {
	got := collectWithTimeout(t, SkipN(feedInts(1,2,3,4), 2))
	if !reflect.DeepEqual(got, []int{3,4}) { t.Fatalf("SkipN() = %v", got) }
}

func TestPipelineSquareEven(t *testing.T) {
	got := collectWithTimeout(t, PipelineSquareEven(feedInts(1,2,3,4)))
	if !reflect.DeepEqual(got, []int{4,16}) { t.Fatalf("PipelineSquareEven() = %v", got) }
}

func TestPipelineSum(t *testing.T) {
	got := collectWithTimeout(t, PipelineSum(feedInts(1,2,3)))
	if !reflect.DeepEqual(got, []int{6}) { t.Fatalf("PipelineSum() = %v", got) }
}

func TestBatchChannel(t *testing.T) {
	got := collectWithTimeout(t, BatchChannel(feedInts(1,2,3,4,5), 2))
	want := [][]int{{1,2},{3,4},{5}}
	if !reflect.DeepEqual(got, want) { t.Fatalf("BatchChannel() = %v, want %v", got, want) }
}

func TestFlattenChannels(t *testing.T) {
	in := make(chan []int)
	go func(){ defer close(in); in <- []int{1,2}; in <- []int{3}; in <- []int{4,5} }()
	got := collectWithTimeout(t, FlattenChannels(in))
	if !reflect.DeepEqual(got, []int{1,2,3,4,5}) { t.Fatalf("FlattenChannels() = %v", got) }
}

func TestDistinctChannel(t *testing.T) {
	got := collectWithTimeout(t, DistinctChannel(feedInts(1,2,1,3,2,4)))
	if !reflect.DeepEqual(got, []int{1,2,3,4}) { t.Fatalf("DistinctChannel() = %v", got) }
}

func TestRunningSum(t *testing.T) {
	got := collectWithTimeout(t, RunningSum(feedInts(1,2,3,4)))
	if !reflect.DeepEqual(got, []int{1,3,6,10}) { t.Fatalf("RunningSum() = %v", got) }
}

func TestContextDone(t *testing.T) {
	ch := make(chan int,1); ch <- 77
	got, ok := ContextDone(context.Background(), ch)
	if !ok || got != 77 { t.Fatalf("got (%d,%v), want (77,true)", got, ok) }
	ctx, cancel := context.WithCancel(context.Background()); cancel()
	_, ok = ContextDone(ctx, make(chan int))
	if ok { t.Fatalf("ожидали false при отмененном context") }
}

func TestGenerateWithContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := GenerateWithContext(ctx)
	if ch == nil { t.Fatalf("nil channel") }
	for i:=0; i<3; i++ {
		select { case got := <-ch: if got != i { t.Fatalf("got %d, want %d", got, i) }; case <-time.After(300*time.Millisecond): t.Fatalf("timeout") }
	}
	cancel()
	select { case _, ok := <-ch: if ok { /* ok if one buffered value sneaks through */ }; case <-time.After(300*time.Millisecond): t.Fatalf("generator не остановился") }
}

func TestMapWithContext(t *testing.T) {
	ctx := context.Background()
	got := collectWithTimeout(t, MapWithContext(ctx, feedInts(1,2), func(x int) int { return x+1 }))
	if !reflect.DeepEqual(got, []int{2,3}) { t.Fatalf("MapWithContext() = %v", got) }
}

func TestMergeWithContext(t *testing.T) {
	ctx := context.Background()
	got := collectWithTimeout(t, MergeWithContext(ctx, []<-chan int{feedInts(1), feedInts(2,3)}))
	if !sameMultiset(got, []int{1,2,3}) { t.Fatalf("MergeWithContext() = %v", got) }
}

func TestWorkerPoolWithContext(t *testing.T) {
	ctx := context.Background()
	got := collectWithTimeout(t, WorkerPoolWithContext(ctx, feedInts(1,2,3), 2, func(x int) int { return x*x }))
	if !sameMultiset(got, []int{1,4,9}) { t.Fatalf("WorkerPoolWithContext() = %v", got) }
}

func TestStopGenerator(t *testing.T) {
	ch, stop := StopGenerator()
	if ch == nil || stop == nil { t.Fatalf("channel/stop не должны быть nil") }
	select { case <-ch: case <-time.After(300*time.Millisecond): t.Fatalf("generator не генерирует") }
	stop()
	select { case _, ok := <-ch: if ok { /* допускаем одно значение */ }; case <-time.After(300*time.Millisecond): t.Fatalf("generator не остановился") }
}

func TestOrDone(t *testing.T) {
	done := make(chan struct{})
	got := collectWithTimeout(t, OrDone(done, feedInts(1,2,3)))
	if !reflect.DeepEqual(got, []int{1,2,3}) { t.Fatalf("OrDone() = %v", got) }
}

func TestTee(t *testing.T) {
	a,b := Tee(feedInts(1,2,3))
	var ag,bg []int
	var wg sync.WaitGroup
	wg.Add(2)
	go func(){ defer wg.Done(); ag = collectWithTimeout(t,a) }()
	go func(){ defer wg.Done(); bg = collectWithTimeout(t,b) }()
	wg.Wait()
	want := []int{1,2,3}
	if !reflect.DeepEqual(ag, want) || !reflect.DeepEqual(bg, want) { t.Fatalf("Tee() = %v and %v", ag, bg) }
}

func TestBridge(t *testing.T) {
	outer := make(chan (<-chan int))
	go func(){ defer close(outer); outer <- feedInts(1,2); outer <- feedInts(3); outer <- feedInts(4,5) }()
	got := collectWithTimeout(t, Bridge(outer))
	if !reflect.DeepEqual(got, []int{1,2,3,4,5}) { t.Fatalf("Bridge() = %v", got) }
}

func TestGracefulShutdownWorkers(t *testing.T) {
	got := collectWithTimeout(t, GracefulShutdownWorkers(feedInts(1,2,3), 2, func(x int) int { return x+100 }))
	if !sameMultiset(got, []int{101,102,103}) { t.Fatalf("GracefulShutdownWorkers() = %v", got) }
}
