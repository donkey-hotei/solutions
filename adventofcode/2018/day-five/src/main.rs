use std::str::FromStr;
use std::string::ParseError;

#[derive(Debug)]
struct Polymer {
    units: Vec<Unit>,
}

#[derive(Debug, Copy, Clone)]
struct Unit(u8);

impl Unit {
    fn new(letter: u8) -> Unit {
        Unit(letter)
    }

    fn reacts_with(&self, other: &Unit) -> bool {
        (self.0 as i8 - other.0 as i8).abs() == 32
    }
}

impl Polymer {
    fn len(&self) -> usize {
        self.units.len()
    }

    fn react(&self) -> Polymer {
        Polymer {
            units: self.units
                .iter()
                .fold(
                    Vec::with_capacity(self.len() / 2),
                    |mut v: Vec<Unit>, &a| {

                        if let Some(b) = v.last() {
                            if a.reacts_with(b) {
                                v.pop();
                                return v
                            }
                        }

                        v.push(a);
                        v
                    })
        }
    }
}

impl FromStr for Polymer {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let units = s.as_bytes()
            .iter()
            .map(|&unit| Unit::new(unit))
            .collect();

        Ok(Polymer { units })
    }
}

fn main() {
    let input = include_str!("../input.txt").trim();

    if let Ok(polymer) = input.parse::<Polymer>() {
        // part one
        println!("{:?}", polymer.react().len());

        // part two
        let shortest_polymer = (b'A'..=b'Z')
            .map(|letter| {
                input
                    .replace(letter as char, "")
                    .replace((letter as char).to_ascii_lowercase(), "")
                    .parse::<Polymer>()
                    .unwrap()
                    .react()
            })
            .min_by_key(|polymer| polymer.len());

        println!("{:?}", shortest_polymer.unwrap().react().len());
    }
}
