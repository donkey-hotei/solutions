pub fn is_armstrong_number(num: u32) -> bool {
    let mut sum = 0;
    let mut exp = 0;
    let mut tmp = num;

    while tmp != 0 {
        tmp /= 10;
        exp += 1;
    }

    tmp = num;

    while tmp != 0 {
        sum += (tmp % 10).pow(exp);
        tmp /= 10;
    }

    sum == num
}
