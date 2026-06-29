# coding: utf-8

from llama_index.core import VectorStoreIndex, StorageContext, SimpleDirectoryReader, Settings
from llama_index.vector_stores.milvus import MilvusVectorStore
from llama_index.llms.openai import OpenAI
from llama_index.embeddings.huggingface import HuggingFaceEmbedding

# 使用 DeepSeek（兼容 OpenAI API）
llm = OpenAI(
    model="deepseek-chat",
    api_base="https://api.deepseek.com/v1",
    api_key="sk-538c77b76dbb47b6aff851f68ab428cb",
)

# 使用本地 HuggingFace Embedding 模型
embed_model = HuggingFaceEmbedding(
    model_name="BAAI/bge-small-zh-v1.5",  # 轻量中文嵌入模型
)

Settings.llm = llm
Settings.embed_model = embed_model

# 1. 加载文档
documents = SimpleDirectoryReader("./data/").load_data()

# 2. 配置 Milvus 向量存储（启用混合搜索）
vector_store = MilvusVectorStore(
    uri="./milvus_demo.db",  # 使用 Milvus Lite，数据存于本地文件
    dim=512,                # bge-small-zh-v1.5 向量维度为 512
    enable_sparse=True,     # 启用全文搜索（BM25）
    overwrite=True          # 覆盖已有同名集合
)

# 3. 构建索引
storage_context = StorageContext.from_defaults(vector_store=vector_store)
index = VectorStoreIndex.from_documents(documents, storage_context=storage_context)

# 4. 查询（使用混合搜索模式）
query_engine = index.as_query_engine(vector_store_query_mode="hybrid")
response = query_engine.query("你的问题")
print(response)
