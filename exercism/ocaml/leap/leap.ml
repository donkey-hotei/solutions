let leap_year year =
  if year mod 4 = 0
  then
    year mod 400 = 0 || not (year mod 100 = 0)
  else false
