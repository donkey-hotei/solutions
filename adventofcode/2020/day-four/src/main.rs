use std::fs;
use regex::Regex;
use raster::Color;
use std::str::FromStr;
use itertools::Itertools;
use std::collections::HashMap;

#[derive(Debug)]
struct Passport {
    byr: String,  // birth year
    iyr: String,  // issue year
    eyr: String,  // expiration year
    hgt: String,  // height
    hcl: String,  // hair color
    ecl: String,  // eye color
    pid: String,  // passport id
    cid: String,  // country id
}

impl Passport {
    fn is_valid(&self) -> bool {
        let byr = self.byr.parse::<usize>().unwrap_or_default();
        let iyr = self.iyr.parse::<usize>().unwrap_or_default();
        let eyr = self.eyr.parse::<usize>().unwrap_or_default();

        let ecl_regex = Regex::new("[amb|blu|brn|gry|grn|hzl|oth]").unwrap();

        1920 <= byr && byr <= 2002
            && 2010 <= iyr && iyr <= 2020
            && 2020 <= eyr && eyr <= 2030
            && self.hgt != ""
            && {
                let hgt = self.hgt[..self.hgt.len()-2]
                    .parse::<usize>()
                    .unwrap_or_default();

                match &self.hgt[self.hgt.len()-2..] {
                    "in" => {
                        59 <= hgt && hgt <= 76
                    },
                    "cm" => {
                        150 <= hgt && hgt <= 193
                    },
                    _ => false,
                }
            }
            && Color::hex(&self.hcl).is_ok()
            && ecl_regex.is_match(&self.ecl)
            && self.pid.len() == 9
    }
}

#[derive(Debug)]
struct PassportParseError;

impl FromStr for Passport {
    type Err = PassportParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {        
        let fields = s
            .split(|c| c == ' ' || c == '\n')
            .map(|field| field.split(":").collect_tuple().unwrap_or(("", "")))
            .collect::<HashMap<&str, &str>>();

        Ok(Passport {
            byr: fields.get("byr").unwrap_or(&"").to_string(),
            iyr: fields.get("iyr").unwrap_or(&"").to_string(),
            eyr: fields.get("eyr").unwrap_or(&"").to_string(),
            hgt: fields.get("hgt").unwrap_or(&"").to_string(),
            hcl: fields.get("hcl").unwrap_or(&"").to_string(),
            ecl: fields.get("ecl").unwrap_or(&"").to_string(),
            pid: fields.get("pid").unwrap_or(&"").to_string(),
            cid: fields.get("cid").unwrap_or(&"").to_string(),
        })
    }
}

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();

    let passports = input
        .split("\n\n")
        .map(|passport| passport.parse::<Passport>())
        .filter_map(Result::ok)
        .collect::<Vec<_>>();

    let mut count = 0;

    for passport in passports {
        if passport.is_valid() {
            count += 1;
        }
    }

    println!("{:?}", count);
}
