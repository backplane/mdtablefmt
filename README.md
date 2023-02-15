# mdtablefmt

command-line utility for reformatting markdown tables

`mdtablefmt` reads the text for a single markdown table from the standard input and prints the table to the standard output with the columns properly aligned.

## Example

These tables both render the same but the latter (which has been piped through `mdtablefmt`) is easier to read.

### Input

```markdown
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
```

### Output

```markdown
                               stat |       before       |       after
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
```

The output normally has trailing spaces stripped, if you want to retain trailing spaces you can use the `-no-strip` flag. The following shows the difference in output (**with spaces converted to `.` here for better visibility**).

Without `-no-strip`:

```
...............................stat.|.......before.......|.......after
----------------------------------:.|.:----------------:.|.:--------------:
.............................Memory.|........8GiB........|......127GiB
...............................CPUs.|..........4.........|........16
......................Storage.space.|........423.G.......|.......31.T
.....................Storage.driver.|........aufs........|.....overlay2
.................................OS.|.Debian.GNU/Linux.8.|.Ubuntu.20.04.LTS
.....................Docker.version.|.....17.05.0-ce.....|....19.03.9-ce
...............Memory.limit.support.|.........no.........|........yes
.................Swap.limit.support.|.........no.........|........yes
........Kernel.memory.limit.support.|.........no.........|........yes
...........OOM.kill.disable.support.|.........no.........|........yes
..............CPU.cfs.quota.support.|.........no.........|........yes
.............CPU.cfs.period.support.|.........no.........|........yes
Virtualization.in.container.support.|.........no.........|........yes
```

With `-no-strip`:

```
...............................stat.|.......before.......|.......after.....
----------------------------------:.|.:----------------:.|.:--------------:
.............................Memory.|........8GiB........|......127GiB.....
...............................CPUs.|..........4.........|........16.......
......................Storage.space.|........423.G.......|.......31.T......
.....................Storage.driver.|........aufs........|.....overlay2....
.................................OS.|.Debian.GNU/Linux.8.|.Ubuntu.20.04.LTS
.....................Docker.version.|.....17.05.0-ce.....|....19.03.9-ce...
...............Memory.limit.support.|.........no.........|........yes......
.................Swap.limit.support.|.........no.........|........yes......
........Kernel.memory.limit.support.|.........no.........|........yes......
...........OOM.kill.disable.support.|.........no.........|........yes......
..............CPU.cfs.quota.support.|.........no.........|........yes......
.............CPU.cfs.period.support.|.........no.........|........yes......
Virtualization.in.container.support.|.........no.........|........yes......
```


## Usage

The program prints the following text when invoked with the -h or --help flag:

```
usage: mdtablefmt

utility for formatting Markdown tables; reads valid markdown table data on the
standard input and writes a properly spaced and padded version on the standard
output

For more information, visit the source code repository:
https://github.com/backplane/mdtablefmt

  -no-strip
    	fully pad the rightmost output column instead of stripping trailing whitespace

```