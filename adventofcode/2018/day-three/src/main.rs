#[macro_use]
extern crate lazy_static;

use std::fs;
use regex::Regex;
use std::str::FromStr;
use std::error::Error;
use std::collections::HashMap;

struct Tile {
    id: usize,
    left: usize,
    top: usize,
    width: usize,
    height: usize,
}

impl Tile {
    fn iter_points(&self) -> IterPoints {
        IterPoints {
            tile: self,
            x: self.left,
            y: self.top,
        }
    }
}

struct IterPoints<'t> {
    tile: &'t Tile,
    x: usize,
    y: usize,
}

impl<'t> Iterator for IterPoints<'t> {
    type Item = (usize, usize);

    fn next(&mut self) -> Option<Self::Item> {
        if self.y >= self.tile.top + self.tile.height {
            self.y = self.tile.top;
            self.x += 1;
        }

        if self.x >= self.tile.left + self.tile.width {
            return None;
        }

        let (x, y) = (self.x, self.y);

        self.y += 1;

        Some((x, y))
    }
}

// point maps to number of overlapping tiles at that point
type Grid = HashMap<(usize, usize), usize>;

impl FromStr for Tile {
    type Err = Box<Error>;

    fn from_str(s: &str) -> Result<Tile, Box<Error>> {
        lazy_static! {
            static ref RE: Regex = Regex::new(
                r"#(?P<id>\d+) @ (?P<left>\d+),(?P<top>\d+): (?P<width>\d+)x(?P<height>\d+)")
                .unwrap();

        }

        let caps = match RE.captures(s) {
            Some(caps) => caps,
            None => return Err(Box::<Error>::from(format!("failure to parse claim"))),
        };


        Ok(Tile {
            id: caps["id"].parse()?,
            top: caps["top"].parse()?,
            left: caps["left"].parse()?,
            width: caps["width"].parse()?,
            height: caps["height"].parse()?,
        })
    }
}

fn main() {
    let input = fs::read_to_string("input.txt").expect("error reading file");

    let mut tiles: Vec<Tile> = Vec::new();

    for line in input.lines() {
        tiles.push(line.parse().unwrap());
    }

    let mut grid = Grid::new();

    for tile in &tiles {
        for (x, y) in tile.iter_points() {
            *grid.entry((x, y)).or_default() += 1;
        }
    }

    let count = grid.values().filter(|&&count| count > 1).count();
    println!("{:?}", count);


    for tile in tiles {
        if tile.iter_points().all(|p| grid[&p] == 1) {
            println!("{:?}", tile.id);
        }
    }
}
