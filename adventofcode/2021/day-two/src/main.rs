use std::fs;
use std::str::FromStr;

#[derive(Debug)]
enum Direction {
    Up,
    Down,
    Forward,
    Unknown,
}

#[derive(Debug)]
struct SubmarineCommand {
    direction: Direction,
    units: usize,
}

#[derive(Debug)]
struct Position {
    x: usize,  // horizontal position
    y: usize,  // depth
    aim: usize,
}

#[derive(Debug)]
struct Submarine {
    position: Position,
}

impl Submarine {
    fn new() -> Self {
        Submarine {
            position: Position { x: 0, y: 0, aim: 0 },
        }
    }

    fn execute(&mut self, command: &SubmarineCommand) {
        match command.direction {
            Direction::Up => {
                self.position.aim -= command.units
            },
            Direction::Down => {
                self.position.aim += command.units
            },
            Direction::Forward => {
                self.position.x += command.units;
                self.position.y += self.position.aim * command.units
            }
            _ => {},
        }
    }
}

#[derive(Debug)]
struct InvalidSubmarineCommand;

impl FromStr for SubmarineCommand {
    type Err = InvalidSubmarineCommand;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let split = s.split(' ').collect::<Vec<&str>>();

        if split.len() != 2 {
            return Err(InvalidSubmarineCommand)
        }

        Ok(SubmarineCommand {
            direction: match split[0] {
                "up" => Direction::Up,
                "down" => Direction::Down,
                "forward" => Direction::Forward,
                _ => Direction::Unknown,
            },
            units: split[1].parse().unwrap(),
        })
    }
}


fn calculate_course(commands: &[SubmarineCommand]) {
    let mut submarine = Submarine::new();

    for command in commands {
        submarine.execute(command);
    }

    dbg!(submarine.position.x * submarine.position.y);
}

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();

    let commands = input
        .split("\n")
        .map(|command| command.parse::<SubmarineCommand>())
        .filter_map(Result::ok)
        .collect::<Vec<_>>();


    calculate_course(&commands);
}
