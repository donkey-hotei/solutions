#!/usr/bin/python
from functools import cmp_to_key


class Player(object):
    def __init__(self, name, score):
        self.name, self.score = name, score

    def __repr__(self):
        return "Player({0}, {1})".format(self.name, self.score)

    def comprator(self):
        return (-self.score, name)
