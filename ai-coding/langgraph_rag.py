import os
import numpy as np
from typing import TypedDict, List

from langgraph.graph import StateGraph, END
from langchain_openai import ChatOpenAI
from langchain_core.prompts import ChatPromptTemplate
from langchain_core.embeddings import Embeddings
from langchain_core.vectorstores import InMemoryVectorStore
from langchain_core.documents import Document
from sklearn.feature_extraction.text import TfidfVectorizer


# ---------------------------------------------------------------------------
# 配置
API_KEY = os.environ.get("API_KEY")
BASE_URL = "https://api.deepseek.com/v1"
LLM_MODEL = "deepseek-v4-flash"


# ---------------------------------------------------------------------------
# 本地 TF-IDF 嵌入
class TfidfEmbeddings(Embeddings):
    def __init__(self, max_features: int = 2048):
        self.max_features = max_features
        # ngram_range=(1,2) 捕捉双词搭配, sublinear_tf 压制高频词影响
        self.vectorizer = TfidfVectorizer(
            max_features=max_features,
            ngram_range=(1, 2),
            sublinear_tf=True,
        )

    def embed_documents(self, texts: List[str]) -> List[List[float]]:
        vectors = self.vectorizer.fit_transform(texts).toarray()
        for i in range(len(vectors)):
            norm = np.linalg.norm(vectors[i])
            vectors[i] = vectors[i] / norm if norm > 0 else np.full(self.max_features, 1.0 / self.max_features)
        return vectors.tolist()

    def embed_query(self, text: str) -> List[float]:
        vec = self.vectorizer.transform([text]).toarray()[0]
        norm = np.linalg.norm(vec)
        if norm > 0:
            vec = vec / norm
        else:
            vec = np.full(self.max_features, 1.0 / self.max_features)
        return vec.tolist()


# ---------------------------------------------------------------------------
# 初始化向量库（构建一次，全局复用）
def _init_vectorstore() -> InMemoryVectorStore:
    loader = __import__("langchain_community.document_loaders", fromlist=["TextLoader"]).TextLoader
    splitter = __import__("langchain_text_splitters", fromlist=["RecursiveCharacterTextSplitter"]).RecursiveCharacterTextSplitter

    docs = loader("README.md", encoding="utf-8").load()
    chunks = splitter(chunk_size=500, chunk_overlap=50).split_documents(docs)

    # 加一个项目摘要文档，弥合自然语言问法和关键词之间的差距
    summary = Document(
        page_content=(
            "这个项目是 labuladong 的算法笔记，主要讲解算法和数据结构知识，包括：\n"
            "- 核心刷题框架：双指针、滑动窗口、二叉树、动态规划、回溯、BFS、贪心等\n"
            "- 数据结构：数组、链表、队列、栈、哈希表、二叉树、图、排序算法\n"
            "- 经典算法：动态规划、回溯搜索、图论、最短路径、并查集\n"
            "- 配套工具：Chrome 插件、VS Code 插件、JetBrains 插件\n"
            "项目基于 LeetCode 题目，强调举一反三和算法思维培养。"
        )
    )
    chunks.insert(0, summary)

    embedding = TfidfEmbeddings(max_features=2048)
    vs = InMemoryVectorStore(embedding)
    vs.add_documents(chunks)
    return vs


vectorstore = _init_vectorstore()


# ─── 定义状态 ───
class GraphState(TypedDict):
    question: str
    documents: List[str]
    needs_web: bool
    answer: str
    retries: int


# ─── 1. 检索节点 ───
def retrieve(state: GraphState) -> GraphState:
    retriever = vectorstore.as_retriever(search_kwargs={"k": 3})
    docs = retriever.invoke(state["question"])
    state["documents"] = [d.page_content for d in docs]
    return state


# ─── 2. 判断节点（是否需联网） ───
def grade_documents(state: GraphState) -> GraphState:
    llm = ChatOpenAI(model=LLM_MODEL, temperature=0, api_key=API_KEY, base_url=BASE_URL)

    prompt = ChatPromptTemplate.from_messages([
        ("human", """问题：{question}
检索到的文档：{docs}

这些文档是否能回答问题？如果可以，只回答 yes；如果不可以或信息不足，只回答 no。""")
    ])
    result = llm.invoke(prompt.invoke({
        "question": state["question"],
        "docs": " ".join(state["documents"]),
    }))
    state["needs_web"] = "no" in result.content.lower()
    return state


# ─── 3. 生成答案节点 ───
def generate(state: GraphState) -> GraphState:
    llm = ChatOpenAI(model=LLM_MODEL, temperature=0, api_key=API_KEY, base_url=BASE_URL)
    context = "\n\n".join(state["documents"])

    prompt = ChatPromptTemplate.from_messages([
        ("human", """基于以下信息回答问题。如果信息不足，直接说不知道。

上下文：
{context}

问题：{question}""")
    ])
    result = llm.invoke(prompt.invoke({
        "context": context,
        "question": state["question"],
    }))
    state["answer"] = result.content
    return state


# ─── 4. 联网搜索节点（模拟） ───
def web_search(state: GraphState) -> GraphState:
    state["documents"].append(f"[联网补充] 关于'{state['question']}'的最新信息...")
    state["retries"] += 1
    return state


# ─── 条件边 ───
def decide_next(state: GraphState) -> str:
    if state["needs_web"] and state["retries"] < 2:
        return "web_search"
    else:
        return "generate"


# ─── 构建图 ───
def build_rag_agent():
    workflow = StateGraph(GraphState)

    workflow.add_node("retrieve", retrieve)
    workflow.add_node("grade", grade_documents)
    workflow.add_node("web_search", web_search)
    workflow.add_node("generate", generate)

    workflow.set_entry_point("retrieve")
    workflow.add_edge("retrieve", "grade")
    workflow.add_conditional_edges(
        "grade",
        decide_next,
        {"web_search": "web_search", "generate": "generate"}
    )
    workflow.add_edge("web_search", "grade")
    workflow.add_edge("generate", END)

    return workflow.compile()


# ─── 运行 ───
agent = build_rag_agent()
result = agent.invoke({
    "question": "这个项目主要讲解什么内容？",
    "documents": [],
    "needs_web": False,
    "answer": "",
    "retries": 0,
})
print(result["answer"])