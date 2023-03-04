
class AdjNode:
    def __init__(self, data) -> None:
        self.vertex = data
        self.next = None

class Graph:
    def __init__(self, vertices) -> None:
        self.V = vertices
        self.graph = [None] * self.V
    
    def add_edge(self, src, dest):
        node = AdjNode(dest)
        node.next = self.graph[src]
        self.graph[dest] = node

    
    def print_graph(self):
        for i in range(self.V):
            print(f"Adjacency list of vertex {i}\n head")
            temp = self.graph[i]
            while temp:
                print(f" -> {temp.vertex}", end="")
                temp = temp.next
            print(" \n")

def run():
    V = 5
    graph = Graph(V)
    graph.add_edge(0, 1)
    print(graph.graph)
    graph.add_edge(0, 4)
    print(graph.graph)
    graph.print_graph()