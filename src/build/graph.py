from src.build import unit


class BuildGraph:
    def __init__(self) -> None:
        self._graph: dict[str, list[str]] = {}
        self._node: dict[str, unit.Unit] = {}

    def add_dep(self, node: str, dep: str) -> None:
        self._graph[dep].append(node)

    # 拓扑排序编译
    def parser_graph(self) -> None:
        _build_queue: list[str] = []
        # 初始化入度字典
        in_node = {node: 0 for node in self._graph}

        for node in self._graph:
            for neighbor in self._graph[node]:
                in_node[neighbor] += 1

        queue = [node for node in self._graph if in_node[node] == 0]

        while queue:
            # 弹出队列中的节点
            current_node = queue.pop(0)
            # 将其添加到build_queue中
            _build_queue.append(current_node)

            # 处理所有邻居节点，降低其入度
            for neighbor in self._graph[current_node]:
                in_node[neighbor] -= 1
                # 如果入度为0，加入队列
                if in_node[neighbor] == 0:
                    queue.append(neighbor)

            # 检查图是否有环
            if len(_build_queue) != len(self._graph):
                raise ValueError("The graph has at least one cycle, topological sort is not possible.")

        for node in _build_queue:
            self._node[node].parser()

    def add_node(self,node: unit.Unit) -> None:
        self._node[node.name] = node
        self._graph[node.name] = []

    def get_node(self,name: str) -> unit.Unit:
        return self._node[name]