use std::fs;
use regex::Regex;
use std::str::FromStr;


#[derive(Debug)]
struct PasswordPolicy {
    minimum: usize,
    maximum: usize,
    letter: String,
    password: String,
}

impl PasswordPolicy {
    fn is_valid(&self) -> bool {
        let count = self.password.matches(&self.letter).count();

        if self.minimum <= count && count <= self.maximum {
            return true;
        }

        false
    }

    fn is_valid_by_position(&self) -> bool {
        let position_one = self.password.as_bytes()[self.minimum - 1] == self.letter.as_bytes()[0];
        let position_two = self.password.as_bytes()[self.maximum - 1] == self.letter.as_bytes()[0];

        position_one ^ position_two
    }
}

impl FromStr for PasswordPolicy {
    type Err = regex::Error;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let re = Regex::new(
            r"(?P<minimum>\d+)-(?P<maximum>\d+) (?P<letter>[a-z]): (?P<password>[a-z]+)"
        ).unwrap();

        let captures = match re.captures(s) {
            Some(m) => m,
            None => return Err(regex::Error::Syntax(s.to_string())),
        };

        Ok(PasswordPolicy {
            minimum: captures
                .name("minimum")
                .unwrap()
                .as_str()
                .parse::<usize>()
                .unwrap(),
            maximum: captures
                .name("maximum")
                .unwrap()
                .as_str()
                .parse::<usize>()
                .unwrap(),
            letter: captures
                .name("letter")
                .unwrap()
                .as_str()
                .to_string(),
            password: captures
                .name("password")
                .unwrap()
                .as_str()
                .to_string(),
        })
    }
}

fn part_one(passwords: &[PasswordPolicy]) {
    let mut valid_password_count = 0;

    for password in passwords {
        if password.is_valid() {
            valid_password_count += 1;
        }
    }

    println!("part one: {:?}", valid_password_count);
}

fn part_two(passwords: &[PasswordPolicy]) {
    let mut valid_password_count = 0;

    for password in passwords {
        if password.is_valid_by_position() {
            valid_password_count += 1;
        }
    }

    println!("part two: {:?}", valid_password_count);
}

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();

    let passwords = input
        .split("\n")
        .map(|entry| entry.parse::<PasswordPolicy>())
        .filter_map(Result::ok)
        .collect::<Vec<_>>();

    part_one(&passwords);
    part_two(&passwords);
}
