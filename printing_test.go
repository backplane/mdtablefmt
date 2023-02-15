package main

import (
	"bytes"
	"strings"
	"testing"
)

const input string = `
stat | before | after
---: | :-----: | :------:
Memory | 8GiB | 127GiB
CPUs | 4 | 16
Storage space | 423 G | 31 T
Storage driver | aufs | overlay2
OS | Debian GNU/Linux 8 | Ubuntu 20.04 LTS
Docker version | 17.05.0-ce | 19.03.9-ce
Memory limit support | no | yes
Swap limit support | no | yes
Kernel memory limit support | no | yes
OOM kill disable support | no | yes
CPU cfs quota support | no | yes
CPU cfs period support | no | yes
Virtualization in container support | no | yes
`

const expected_output_padded string = `                               stat |       before       |       after     
----------------------------------: | :----------------: | :--------------:
                             Memory |        8GiB        |      127GiB     
                               CPUs |          4         |        16       
                      Storage space |        423 G       |       31 T      
                     Storage driver |        aufs        |     overlay2    
                                 OS | Debian GNU/Linux 8 | Ubuntu 20.04 LTS
                     Docker version |     17.05.0-ce     |    19.03.9-ce   
               Memory limit support |         no         |        yes      
                 Swap limit support |         no         |        yes      
        Kernel memory limit support |         no         |        yes      
           OOM kill disable support |         no         |        yes      
              CPU cfs quota support |         no         |        yes      
             CPU cfs period support |         no         |        yes      
Virtualization in container support |         no         |        yes      
`

const expected_output_stripped string = `                               stat |       before       |       after
----------------------------------: | :----------------: | :--------------:
                             Memory |        8GiB        |      127GiB
                               CPUs |          4         |        16
                      Storage space |        423 G       |       31 T
                     Storage driver |        aufs        |     overlay2
                                 OS | Debian GNU/Linux 8 | Ubuntu 20.04 LTS
                     Docker version |     17.05.0-ce     |    19.03.9-ce
               Memory limit support |         no         |        yes
                 Swap limit support |         no         |        yes
        Kernel memory limit support |         no         |        yes
           OOM kill disable support |         no         |        yes
              CPU cfs quota support |         no         |        yes
             CPU cfs period support |         no         |        yes
Virtualization in container support |         no         |        yes
`

func TestPrintMarkdownTable(t *testing.T) {
	want := expected_output_stripped

	table, err := ParseMarkdownTable(strings.NewReader(input))
	if err != nil {
		t.Fatalf("failed to parse markdown table from test input string, error: %s", err)
	}
	buf := new(bytes.Buffer)
	PrintMarkdownTable(buf, table, false)
	have := buf.String()

	if have != want {
		t.Fatalf("PrintMarkdownTable failed; have:\nSTART\n%s\nEND\n, want:\nSTART\n%s\nEND\n", have, want)
	}
}

func TestPrintMarkdownTableUnstripped(t *testing.T) {
	want := expected_output_padded

	table, err := ParseMarkdownTable(strings.NewReader(input))
	if err != nil {
		t.Fatalf("failed to parse markdown table from test input string, error: %s", err)
	}
	buf := new(bytes.Buffer)
	PrintMarkdownTable(buf, table, true)
	have := buf.String()

	if have != want {
		t.Fatalf("PrintMarkdownTable failed; have:\nSTART\n%s\nEND\n, want:\nSTART\n%s\nEND\n", have, want)
	}
}
