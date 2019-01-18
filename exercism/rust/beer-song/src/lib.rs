fn bottle(n: i32) -> String {
    match n {
        1 => format!("1 bottle"),
        _ => format!("{} bottles", n)
    }
}

pub fn verse(n: i32) -> String {
    match n {
        0 => {
             format!(
             "No more bottles of beer on the wall, no more bottles of beer.\n\
              Go to the store and buy some more, {} bottles of beer on the wall.\n",
              99)
        },
        1 => {
            format!(
            "1 bottle of beer on the wall, 1 bottle of beer.\n\
            Take it down and pass it around, no more bottles of beer on the wall.\n")
        }
        _ => {
            format!("{} bottles of beer on the wall, {} bottles of beer.\n\
                    Take one down and pass it around, {} bottles of beer on the wall.\n",
                    n,
                    n,
                    bottle(n - 1))
        }
    }
}

pub fn sing(start: i32, end: i32) -> String {
    (end..=start).rev().map(|n| verse(n)).collect()
}
