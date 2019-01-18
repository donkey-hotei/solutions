use std::iter::once;

pub fn build_proverb(list: &[&str]) -> String {
    if list.len() == 0 { return String::new() }

    list
        .windows(2)
        .map(|w| format!("For want of a {} the {} was lost.\n", w[0], w[1]))
        .chain(once(format!("And all for the want of a {}.", list[0])))
        .collect()
}
