#[macro_use]
extern crate prettytable;
use prettytable::format;

fn main() {
    let mut table = table!(
      [l -> "",               l -> "5.1.z",  r -> "2.0"],
      [l -> "",               l -> "5.2.z",  r -> "3.0"],
      [l -> "Anand",          l -> "5.3.z",  r -> "3.1"],
      [l -> "Botvinnik",      l -> "5.4.z",  r -> "3.2"],
      [l -> "Capablanca",     l -> "5.5.z",  r -> "4.0"],
      [l -> "Darga",          l -> "5.6.z",  r -> "4.1"],
      [l -> "Euwe",           l -> "5.7.z",  r -> "4.2"],
      [l -> "Fine",           l -> "5.8.z",  r -> "4.5"],
      [l -> "Gaprindashvili", l -> "5.9.z",  r -> "4.6"],
      [l -> "Hammer",         l -> "5.10.z", r -> "4.7"]
    );

    table.set_titles(row![c => "MANAGEIQ", "CLOUDFORMS MANAGEMENT ENGINE", "CLOUDFORMS"]);

    table.set_format(*format::consts::FORMAT_NO_LINESEP_WITH_TITLE);
    table.printstd();
}
