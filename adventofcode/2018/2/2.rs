#![feature(iterator_find_map)]
use std::fs;
use std::collections::{HashSet, HashMap};

#[derive(Default, Debug)]
struct Node<'a> {
    depth: usize,
    prefix: &'a str,
    same_prefix: HashSet<&'a str>, //
    neighbors: HashMap<char, Node<'a>>
}

fn add_to_tree<'a>(root: &mut Node<'a>, id: &'a str) {
    let mut current = root;

    current.same_prefix.insert(&id);

    for (i, c) in id.chars().enumerate() {
        {
            let mut next = current
                .neighbors
                .entry(c)
                .or_insert(Node::default());

            next.depth = i + 1;
            next.same_prefix.insert(&id);
            next.prefix = &id[..=i];
        }

        current = { current }.neighbors.get_mut(&c).unwrap();
    }
}

fn check_child_matches<'a>(node: &Node<'a>) -> Option<String> {
    let edges: Vec<_> = node.neighbors.keys().collect();
    let suffix_candidates: HashMap<char, HashSet<&str>> = edges
        .iter()
        .map(|c| {
            let mut suffix_count = HashMap::<&str, HashSet<&str>>::new();
            let ref ids = node.neighbors.get(&c).unwrap().same_prefix;

            for id in ids {
                suffix_count
                    .entry(&id[node.depth + 1..])
                    .or_insert(HashSet::new())
                    .insert(id);
            }

            (
                **c,
                suffix_count
                    .iter()
                    .filter(|(_, count)| count.len() == 1)
                    .map(|(suffix, _)| *suffix)
                    .collect(),
            )
        }).collect();

    let mut suffix_counter: HashMap<&str, usize> = HashMap::new();

    for suffix_set in suffix_candidates.values() {
        for suffix in suffix_set {
            *suffix_counter.entry(suffix).or_insert(0) += 1;
        }
    }

    if let Some(suffix) = 
        suffix_counter
            .iter()
            .find_map(|(suffix, count)| if *count == 2 { Some(suffix) } else { None })
    {
        Some(format!("{}{}", node.prefix, &suffix))
    } else {
        None
    }
}

fn find_some_match(node: &Node) -> Option<String> {
    if let Some(result) = check_child_matches(&node) {
        return Some(result)
    } else {
        for child in node.neighbors.values() {
            if let Some(result) = find_some_match(&child) {
                return Some(result);
            }
        }

        None
    }
}

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();

    let mut root = Node::default();

    for id in input.lines() {
        add_to_tree(&mut root, id);
    }

    let result = find_some_match(&root);
    println!("{:?}", result);
}
