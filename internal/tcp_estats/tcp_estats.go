package tcp_estats

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/ringbuf"
	"github.com/cilium/ebpf/rlimit"
)

var (
	estats_db *DB

	removeMemlock = func() error {
		return rlimit.RemoveMemlock()
	}
)

func init() {
}

type TcpEstats struct {
	objs                                                       tcp_estatsObjects
	createActive, createInactive                               link.Link
	updateSegrecv, updateSegsend, updateFinishSegrecv          link.Link
	dataRecv, dataQueueOfo                                     link.Link
	rttEstimator, enterLoss, enterCwr                          link.Link
	fastretransAlert                                           link.Link
	globalRd, connRd, perfRd, pathRd, stackRd, appRd, extrasRd *ringbuf.Reader
}

func (t *TcpEstats) createProgramLinks() error {
	var err error
	t.createActive, err = link.AttachTracing(link.TracingOptions{
		Program: t.objs.tcp_estatsPrograms.TcpEstatsCreateActive,
	})
	if err != nil {
		return fmt.Errorf("attaching tracing [TcpEstatsCreateActive]: %v", err)
	}

	t.createInactive, err = link.AttachTracing(link.TracingOptions{
		Program: t.objs.tcp_estatsPrograms.TcpEstatsCreateInactive,
	})
	if err != nil {
		return fmt.Errorf("attaching tracing [TcpEstatsCreateInactive]: %v", err)
	}

	t.updateSegrecv, err = link.AttachTracing(link.TracingOptions{
		Program: t.objs.tcp_estatsPrograms.TcpEstatsUpdateSegrecv,
	})
	if err != nil {
		return fmt.Errorf("attaching tracing [TcpEstatsUpdateSegrecv]: %v", err)
	}

	t.updateSegsend, err = link.AttachTracing(link.TracingOptions{
		Program: t.objs.tcp_estatsPrograms.TcpEstatsUpdateSegsend,
	})
	if err != nil {
		return fmt.Errorf("attaching tracing [TcpEstatsUpdateSegsend]: %v", err)
	}

	t.updateFinishSegrecv, err = link.AttachTracing(link.TracingOptions{
		Program: t.objs.tcp_estatsPrograms.TcpEstatsUpdateFinishSegrecv,
	})
	if err != nil {
		return fmt.Errorf("attaching tracing [TcpEstatsUpdateFinishSegrecv]: %v", err)
	}

	t.dataRecv, err = link.AttachTracing(link.TracingOptions{
		Program: t.objs.tcp_estatsPrograms.TcpEventDataRecv,
	})
	if err != nil {
		return fmt.Errorf("attaching tracing [TcpEventDataRecv]: %v", err)
	}

	/*
	t.dataQueueOfo, err = link.AttachTracing(link.TracingOptions{
		Program: t.objs.tcp_estatsPrograms.TcpDataQueueOfo,
	})
	if err != nil {
		return fmt.Errorf("attaching tracing [TcpDataQueueOfo]: %v", err)
	}
	*/

	t.rttEstimator, err = link.AttachTracing(link.TracingOptions{
		Program: t.objs.tcp_estatsPrograms.TcpRttEstimator,
	})
	if err != nil {
		return fmt.Errorf("attaching tracing [TcpRttEstimator]: %v", err)
	}

	t.enterLoss, err = link.AttachTracing(link.TracingOptions{
		Program: t.objs.tcp_estatsPrograms.TcpEnterLoss,
	})
	if err != nil {
		return fmt.Errorf("attaching tracing [TcpEnterLoss]: %v", err)
	}

	t.enterCwr, err = link.AttachTracing(link.TracingOptions{
		Program: t.objs.tcp_estatsPrograms.TcpEnterCwr,
	})
	if err != nil {
		return fmt.Errorf("attaching tracing [TcpEnterCwr]: %v", err)
	}

	t.fastretransAlert, err = link.AttachTracing(link.TracingOptions{
		Program: t.objs.tcp_estatsPrograms.TcpFastretransAlert,
	})
	if err != nil {
		return fmt.Errorf("attaching tracing [TcpFastretransAlert]: %v", err)
	}

	return nil
}

func (t *TcpEstats) createRingBuffers() error {
	var err error

	t.globalRd, err = ringbuf.NewReader(t.objs.tcp_estatsMaps.GlobalTable)
	if err != nil {
		return fmt.Errorf("opening global table reader: %v", err)
	}

	t.connRd, err = ringbuf.NewReader(t.objs.tcp_estatsMaps.ConnectionTable)
	if err != nil {
		return fmt.Errorf("opening connection table reader: %v", err)
	}

	t.perfRd, err = ringbuf.NewReader(t.objs.tcp_estatsMaps.PerfTable)
	if err != nil {
		return fmt.Errorf("opening perf table reader: %v", err)
	}

	t.pathRd, err = ringbuf.NewReader(t.objs.tcp_estatsMaps.PathTable)
	if err != nil {
		return fmt.Errorf("opening path table reader: %v", err)
	}

	t.stackRd, err = ringbuf.NewReader(t.objs.tcp_estatsMaps.StackTable)
	if err != nil {
		return fmt.Errorf("opening stack table reader: %v", err)
	}

	t.appRd, err = ringbuf.NewReader(t.objs.tcp_estatsMaps.AppTable)
	if err != nil {
		return fmt.Errorf("opening app table reader: %v", err)
	}

	t.extrasRd, err = ringbuf.NewReader(t.objs.tcp_estatsMaps.ExtrasTable)
	if err != nil {
		return fmt.Errorf("opening extras table reader: %v", err)
	}

	return nil
}

func New() (*TcpEstats, error) {
	if err := removeMemlock(); err != nil {
		return nil, err
	}

	estats_db = newDB()
	t := TcpEstats{}

	// load pre-compiled programs into the kernel
	t.objs = tcp_estatsObjects{}
	if err := loadTcp_estatsObjects(&t.objs, nil); err != nil {
		return nil, fmt.Errorf("loading objects: %v", err)
	}

	if err := t.createProgramLinks(); err != nil {
		return nil, err
	}

	if err := t.createRingBuffers(); err != nil {
		return nil, err
	}

	return &t, nil
}

func (t *TcpEstats) Close() error {
	// TODO: consider multiple errors instead of stopping if one errors.
	if err := t.globalRd.Close(); err != nil {
		return err
	}
	if err := t.connRd.Close(); err != nil {
		return err
	}
	if err := t.pathRd.Close(); err != nil {
		return err
	}
	if err := t.perfRd.Close(); err != nil {
		return err
	}
	if err := t.stackRd.Close(); err != nil {
		return err
	}
	if err := t.appRd.Close(); err != nil {
		return err
	}
	if err := t.extrasRd.Close(); err != nil {
		return err
	}

	if err := t.createActive.Close(); err != nil {
		return err
	}
	if err := t.createInactive.Close(); err != nil {
		return err
	}
	if err := t.updateSegrecv.Close(); err != nil {
		return err
	}
	if err := t.updateSegsend.Close(); err != nil {
		return err
	}
	if err := t.updateFinishSegrecv.Close(); err != nil {
		return err
	}
	if err := t.dataRecv.Close(); err != nil {
		return err
	}
	if err := t.dataQueueOfo.Close(); err != nil {
		return err
	}
	if err := t.rttEstimator.Close(); err != nil {
		return err
	}
	if err := t.enterLoss.Close(); err != nil {
		return err
	}
	if err := t.enterCwr.Close(); err != nil {
		return err
	}

	if err := t.objs.Close(); err != nil {
		return err
	}

	return nil
}

func (t *TcpEstats) Run() {
	go readLoop[GlobalVar](t.globalRd)
	go readLoop[ConnectionVar](t.connRd)
	go readLoop[PerfVar](t.perfRd)
	go readLoop[PathVar](t.pathRd)
	go readLoop[StackVar](t.stackRd)
	go readLoop[AppVar](t.appRd)
	go readLoop[ExtrasVar](t.extrasRd)
}

func (t *TcpEstats) Dump() ([]byte, error) {
	return json.MarshalIndent(estats_db, "", "  ")
}

func readLoop[V Vars](rd *ringbuf.Reader) {
	var record Record
	for {
		item, err := rd.Read()
		if err != nil {
			if errors.Is(err, ringbuf.ErrClosed) {
				// log.Println("received signal, exiting loop..")
				return
			}
			continue
		}

		// parse to structure
		if err := binary.Read(bytes.NewBuffer(item.RawSample), native, &record); err != nil {
			//log.Printf("parsing entry: %v", err)
			continue
		}

		// There might be a way to get away with a RLock here followed
		// by a Lock in the unlikely case we need to insert, but just taking
		// the more expensive lock is easier.
		estats_db.Lock()

		k := key{
			PidTgid: record.PidTgid,
			Saddr:   record.Saddr,
			Daddr:   record.Daddr,
			Sport:   record.Sport,
			Dport:   record.Dport,
		}

		e, ok := estats_db.m[k]
		if !ok {
			e = newEstats()
			estats_db.m[k] = e
		}
		estats_db.Unlock()

		doOp[V](e, record.Op, V(record.Var), record.Val)
	}
}
