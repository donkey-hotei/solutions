use std::fs;
use std::iter;
use itermore::IterMore;

fn pairwise<I>(right: I) -> impl Iterator<Item = (Option<I::Item>, I::Item)>
where
        I: IntoIterator + Clone,
{
        let left = iter::once(None).chain(right.clone().into_iter().map(Some));

        left.zip(right)
}


fn part_one(depths: &[i64]) {
    let mut increase_count: i64 = 0;

    for pair in pairwise(depths) {
        match pair {
            (None, _) => continue,
            (Some(previous), current) => {
                if previous < current {
                    increase_count += 1
                }
            }
        }
    }

    println!("part one: increase count = {:?}", increase_count);
}


fn part_two(depths: &[i64]) {
    let mut increase_count = 0;

    let mut previous_sum: Option<i64> = None;

    for [a, b, c] in depths.iter().windows() {
        let current_sum = a + b + c;

        if previous_sum.is_none() {
            previous_sum = Some(current_sum);

            continue;
        }

        if previous_sum.unwrap() < current_sum {
            increase_count += 1;
        }

        previous_sum = Some(current_sum);
    }

    println!("part two: increase count = {:?}", increase_count);
}


fn main() {
    let input = fs::read_to_string("input.txt").unwrap();

    let depths = input
        .split("\n")
        .map(|entry| entry.parse::<i64>())
        .filter_map(Result::ok)
        .collect::<Vec<_>>();

    part_one(&depths);
    part_two(&depths);
}
