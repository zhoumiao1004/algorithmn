# coding: utf-8

from typing import TypedDict
from langgraph.graph import StateGraph, END

# 1. 定义状态
class State(TypedDict):
    count: int
    message: str

# 2. 定义节点（就是普通函数）
def node_a(state: State) -> State:
    print(f"→ Node A: count={state['count']}")
    state["count"] += 1
    state["message"] = "经过A节点"
    return state

def node_b(state: State) -> State:
    print(f"→ Node B: count={state['count']}")
    state["count"] += 1
    state["message"] = "经过B节点"
    return state

def node_c(state: State) -> State:
    print(f"→ Node C: count={state['count']}")
    state["message"] = "到达终点C"
    return state

# 3. 条件函数：根据 count 决定走 B 还是 C
def router(state: State) -> str:
    if state["count"] < 3:
        return "node_b"     # 继续循环
    else:
        return "node_c"     # 结束

# 4. 构建图
graph = StateGraph(State)
graph.add_node("node_a", node_a)
graph.add_node("node_b", node_b)
graph.add_node("node_c", node_c)

graph.set_entry_point("node_a")
graph.add_edge("node_a", "node_b")
graph.add_conditional_edges("node_b", router, {
    "node_b": "node_b",   # 回到自己（循环）
    "node_c": "node_c"    # 结束
})
graph.add_edge("node_c", END)

# 5. 执行
app = graph.compile()
result = app.invoke({"count": 0, "message": "开始"})
print(f"最终状态: {result}")
