miqversions
---

*Because naming is hard. Really, really hard.*

### Usage

```bash
$ miqversions
+----------------+------------------------------+------------+-------+------------+---------+
|    MANAGEIQ    | CLOUDFORMS MANAGEMENT ENGINE | CLOUDFORMS | RUBY  | POSTGRESQL | ANSIBLE |
+----------------+------------------------------+------------+-------+------------+---------+
|                | 5.1.z                        | 2.0        |       |            |         |
|                | 5.2.z                        | 3.0        |       |            |         |
| Anand          | 5.3.z                        | 3.1        |       |            |         |
| Botvinnik      | 5.4.z                        | 3.2        |       |            |         |
| Capablanca     | 5.5.z                        | 4.0        | 2.2.x | 9.4.x      |         |
| Darga          | 5.6.z                        | 4.1        | 2.2.x | 9.4.x      |         |
| Euwe           | 5.7.z                        | 4.2        | 2.3.x | 9.5.x      |         |
| Fine           | 5.8.z                        | 4.5        | 2.3.x | 9.5.x      |         |
| Gaprindashvili | 5.9.z                        | 4.6        | 2.3.x | 9.5.x      | 2.4.y   |
| Hammer         | 5.10.z                       | 4.7        | 2.4.x | 9.5.x      | 2.7.y   |
| Ivanchuck      | 5.11.z                       | 4.8        | 2.4.x | 9.5.x      | 2.?     |
+----------------+------------------------------+------------+-------+------------+---------+
```

### Installation

#### Homebrew (MacOS)

```bash
$ brew tap chrisarcand/formulae
$ brew install miqversions
```

#### Other platforms

[Golang](https://golang.org/) must be installed.

```bash
$ go get github.com/chrisarcand/miqversions
```

### License

[Apache License, Version 2.0](https://github.com/chrisarcand/miqversion/blob/master/LICENSE) © 2017 [Chris Arcand](https://github.com/chrisarcand)
