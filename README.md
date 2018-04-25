miqversions
---

*Because naming is hard. Really, really hard.*

### Usage

```bash
$ miqversions
+----------------+------------------------------+------------+
|    MANAGEIQ    | CLOUDFORMS MANAGEMENT ENGINE | CLOUDFORMS |
+----------------+------------------------------+------------+
|                | 5.1.z                        |        2.0 |
|                | 5.2.z                        |        3.0 |
| Anand          | 5.3.z                        |        3.1 |
| Botvinnik      | 5.4.z                        |        3.2 |
| Capablanca     | 5.5.z                        |        4.0 |
| Darga          | 5.6.z                        |        4.1 |
| Euwe           | 5.7.z                        |        4.2 |
| Fine           | 5.8.z                        |        4.5 |
| Gaprindashvili | 5.9.z                        |        4.6 |
| Hammer         | 5.10.z                       |        4.7 |
+----------------+------------------------------+------------+
```

### Installation

#### Homebrew (MacOS)

```bash
$ brew tap chrisarcand/formulae
$ brew install miqversions
```

#### Other platforms

[Rust](https://www.rust-lang.org) must be installed.

```bash
$ cargo build --release
$ cp target/release/miqversions /some/where/on/your/path/
```

### License

[Apache License, Version 2.0](https://github.com/chrisarcand/miqversion/blob/master/LICENSE) © 2017 [Chris Arcand](https://github.com/chrisarcand)
