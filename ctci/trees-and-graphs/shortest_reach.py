#!/usr/bin/python
from collections import deque


class Graph:
    """
    A simple undirected Graph implementation using an adjacency list.
    """
    def __init__(self, n):
        self.n = n  # nodes
        self.adjlist = {k: set() for k in range(n)}

    def connect(self, u, v):
        """Create an edge between nodes u and v."""
        self.adjlist[u].add(v)
        self.adjlist[v].add(u)

    def traverse(self, start):
        """
        Finds the shortest path from a start node to each of it's
        descendents using BFS.
        """
        visited = set([start])
        queue = deque([[start]])
        while queue:
            path = queue.popleft()
            node = path[-1]
            yield node, path
            for neighbor in self.adjlist[node] - visited:
                visited.add(neighbor)
                queue.append(path + [neighbor])

    def find_all_distances(self, start):
        """
        Finds the distances from a start node to each of it's descendents,
        returning the list in node-order.
        """
        distances = []

        if start not in self.adjlist:
            return distances

        paths = {node: 6 * (len(path) - 1)
                 for node, path
                 in self.traverse(start)}

        for node in range(self.n):
            if node == start:
                continue
            if node not in paths:
                distances.append(-1)
            else:
                distances.append(paths[node])

        return distances


if __name__ == "__main__":
    import sys
    t = int(sys.stdin.readline())
    for i in range(t):
        n, m = [int(value) for value in sys.stdin.readline().split()]
        graph = Graph(n)

        for i in range(m):
            x, y = [int(x) for x in sys.stdin.readline().split()]
            graph.connect(x-1, y-1)
        s = int(sys.stdin.readline())
        print(" ".join(map(str, graph.find_all_distances(s-1))))
