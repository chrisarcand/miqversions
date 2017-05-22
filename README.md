miqversions
---

*Because naming is hard. Really, really hard.*

### Usage

```bash
$ miqversions
+------------+------------------------------+---------------+
|  MANAGEIQ  | CLOUDFORMS MANAGEMENT ENGINE |  CLOUDFORMS   |
+------------+------------------------------+---------------+
|            | 5.1.x                        |           2.0 |
|            | 5.2.x                        |           3.0 |
| Anand      | 5.3.x                        |           3.1 |
| Botvinnik  | 5.4.x                        |           3.2 |
| Capablanca | 5.5.x                        |           4.0 |
| Darga      | 5.6.x                        |           4.1 |
| Euwe       | 5.7.x                        |           4.2 |
| Fine       | 5.8.x (pending)              | 4.5 (pending) |
+------------+------------------------------+---------------+
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
