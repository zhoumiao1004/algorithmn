# coding: utf-8
"""
RAG 演示：本地 TF-IDF 向量化 + DeepSeek-V4 大模型
环境受限（无法访问 HuggingFace / Azure），故采用纯本地方案。
"""
import os
import numpy as np
from typing import List

from sklearn.feature_extraction.text import TfidfVectorizer

from langchain_community.document_loaders import TextLoader
from langchain_text_splitters import RecursiveCharacterTextSplitter
from langchain_openai import ChatOpenAI
from langchain_core.prompts import ChatPromptTemplate
from langchain_core.embeddings import Embeddings
from langchain_core.vectorstores import InMemoryVectorStore
from langchain_core.documents import Document
from langchain_core.runnables import RunnablePassthrough
from langchain_core.output_parsers import StrOutputParser


# ---------------------------------------------------------------------------
# 配置
API_KEY = os.environ.get("API_KEY")
BASE_URL = "https://api.deepseek.com/v1"
LLM_MODEL = "deepseek-v4-flash"

# ---------------------------------------------------------------------------
# 本地 TF-IDF 嵌入（无需联网下载模型）
class TfidfEmbeddings(Embeddings):
    """用 sklearn 的 TfidfVectorizer 实现本地文本嵌入，维度固定为 512。"""

    def __init__(self, max_features: int = 512):
        self.max_features = max_features
        self.vectorizer = TfidfVectorizer(max_features=max_features)
        self._fitted = False

    def embed_documents(self, texts: List[str]) -> List[List[float]]:
        vectors = self.vectorizer.fit_transform(texts).toarray()
        # L2 归一化，全零向量 → 均匀小值避免 NaN
        for i in range(len(vectors)):
            norm = np.linalg.norm(vectors[i])
            if norm > 0:
                vectors[i] = vectors[i] / norm
            else:
                vectors[i] = np.full(self.max_features, 1.0 / self.max_features)
        self._fitted = True
        return vectors.tolist()

    def embed_query(self, text: str) -> List[float]:
        if not self._fitted:
            return [0.0] * self.max_features
        vec = self.vectorizer.transform([text]).toarray()[0]
        norm = np.linalg.norm(vec)
        if norm > 0:
            vec = vec / norm
        else:
            # 全零向量会导致余弦相似度 0/0 = NaN，返回均匀小值
            vec = np.full(self.max_features, 1.0 / self.max_features)
        return vec.tolist()


# ---------------------------------------------------------------------------
# 1. 加载文档
loader = TextLoader("README.md", encoding="utf-8")
docs = loader.load()
print(f"📄 已加载文档，共 {len(docs)} 段")

# 2. 分块
splitter = RecursiveCharacterTextSplitter(chunk_size=500, chunk_overlap=50)
chunks = splitter.split_documents(docs)
print(f"✂️  分块完成，共 {len(chunks)} 个片段")

# 3. 向量化 + 内存向量存储
embedding = TfidfEmbeddings(max_features=512)
# 先提取文本
chunk_texts = [chunk.page_content for chunk in chunks]
# 批量嵌入
vectors = embedding.embed_documents(chunk_texts)
print(f"🔢 向量化完成，维度: {len(vectors[0]) if vectors else 0}")

# 构建 InMemoryVectorStore（用 Document + 嵌入对的方式）
vectorstore = InMemoryVectorStore(embedding)
vectorstore.add_documents(chunks)
print(f"💾 向量存储完成")

# 4. 构建检索链
llm = ChatOpenAI(
    model=LLM_MODEL,
    temperature=0,
    api_key=API_KEY,
    base_url=BASE_URL,
)

prompt = ChatPromptTemplate.from_template("""
基于以下上下文回答问题。如果找不到答案，就说不知道。

上下文：
{context}

问题：{input}
""")

retriever = vectorstore.as_retriever(search_kwargs={"k": 3})

def format_docs(docs: List[Document]) -> str:
    return "\n\n---\n\n".join(doc.page_content for doc in docs)

rag_chain = (
    {"context": retriever | format_docs, "input": RunnablePassthrough()}
    | prompt
    | llm
    | StrOutputParser()
)

# 5. 查询
print("\n🔍 查询：这个项目是做什么的？")
print("=" * 50)
result = rag_chain.invoke("这个项目是做什么的？")
print(f"\n📝 回答:\n{result}")
