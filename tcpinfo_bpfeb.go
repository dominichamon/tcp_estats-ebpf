// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64be || armbe || mips || mips64 || mips64p32 || ppc64 || s390 || s390x || sparc || sparc64
// +build arm64be armbe mips mips64 mips64p32 ppc64 s390 s390x sparc sparc64

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// loadTcpinfo returns the embedded CollectionSpec for tcpinfo.
func loadTcpinfo() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_TcpinfoBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load tcpinfo: %w", err)
	}

	return spec, err
}

// loadTcpinfoObjects loads tcpinfo and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//     *tcpinfoObjects
//     *tcpinfoPrograms
//     *tcpinfoMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadTcpinfoObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadTcpinfo()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// tcpinfoSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tcpinfoSpecs struct {
	tcpinfoProgramSpecs
	tcpinfoMapSpecs
}

// tcpinfoSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tcpinfoProgramSpecs struct {
	TcpClose *ebpf.ProgramSpec `ebpf:"tcp_close"`
}

// tcpinfoMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tcpinfoMapSpecs struct {
	Samples *ebpf.MapSpec `ebpf:"samples"`
}

// tcpinfoObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadTcpinfoObjects or ebpf.CollectionSpec.LoadAndAssign.
type tcpinfoObjects struct {
	tcpinfoPrograms
	tcpinfoMaps
}

func (o *tcpinfoObjects) Close() error {
	return _TcpinfoClose(
		&o.tcpinfoPrograms,
		&o.tcpinfoMaps,
	)
}

// tcpinfoMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadTcpinfoObjects or ebpf.CollectionSpec.LoadAndAssign.
type tcpinfoMaps struct {
	Samples *ebpf.Map `ebpf:"samples"`
}

func (m *tcpinfoMaps) Close() error {
	return _TcpinfoClose(
		m.Samples,
	)
}

// tcpinfoPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadTcpinfoObjects or ebpf.CollectionSpec.LoadAndAssign.
type tcpinfoPrograms struct {
	TcpClose *ebpf.Program `ebpf:"tcp_close"`
}

func (p *tcpinfoPrograms) Close() error {
	return _TcpinfoClose(
		p.TcpClose,
	)
}

func _TcpinfoClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed tcpinfo_bpfeb.o
var _TcpinfoBytes []byte