# coding: utf-8

import os
from llama_index.core import VectorStoreIndex, StorageContext, SimpleDirectoryReader, Settings
from llama_index.vector_stores.milvus import MilvusVectorStore
from llama_index.llms.openai_like import OpenAILike
from llama_index.embeddings.huggingface import HuggingFaceEmbedding
from llama_index.readers.file import PDFReader

# 使用 DeepSeek（兼容 OpenAI API）
# 注意：DeepSeek 不在 OpenAI 官方模型列表里，必须用 OpenAILike 而非 OpenAI
# API key 通过环境变量 DEEPSEEK_API_KEY 传入，避免硬编码
llm = OpenAILike(
    model="deepseek-chat",
    api_base="https://api.deepseek.com/v1",
    api_key=os.environ["DEEPSEEK_API_KEY"],
    is_chat_model=True,        # 走 /chat/completions 端点
    context_window=64000,      # deepseek-chat 上下文窗口
)

# 使用本地 HuggingFace Embedding 模型
embed_model = HuggingFaceEmbedding(
    model_name="BAAI/bge-small-zh-v1.5",  # 轻量中文嵌入模型
)

Settings.llm = llm
Settings.embed_model = embed_model

# 1. 加载文档（显式指定 PDFReader，否则 PDF 会被当作原始字节读入）
documents = SimpleDirectoryReader(
    "./data/",
    file_extractor={".pdf": PDFReader()},
).load_data()
print(f"已加载 {len(documents)} 个文档，总字符数 {sum(len(d.text) for d in documents)}")

# 2. 配置 Milvus 向量存储
# 说明：
#   1) 当前 milvus-lite 3.0 缺少 AllocTimestamp 实现，BM25 稀疏向量无法插入。
#      暂时只用稠密向量；如需混合检索，请降级 pymilvus 或改用独立 Milvus 服务。
#   2) pymilvus 2.6.16 + milvus-lite 3.0 的 search 接口用 output_fields=["*"] 时
#      只返回 id，不返回 text。必须显式列出字段才能让 LlamaIndex 读到节点文本。
vector_store = MilvusVectorStore(
    uri="./milvus_demo.db",  # 使用 Milvus Lite，数据存于本地文件
    dim=512,                # bge-small-zh-v1.5 向量维度为 512
    enable_sparse=False,    # 关闭稀疏检索（BM25）以兼容 milvus-lite
    overwrite=True,         # 覆盖已有同名集合
    output_fields=["id", "doc_id", "text", "embedding"],  # 关键：必须显式列字段
)

# 3. 构建索引
storage_context = StorageContext.from_defaults(vector_store=vector_store)
index = VectorStoreIndex.from_documents(documents, storage_context=storage_context)

# 4. 查询（稠密向量模式，循环一问一答）
query_engine = index.as_query_engine()
print("天翼云息壤知识库已就绪，输入问题开始对话（输入 q / quit / exit 退出）。\n")
while True:
    try:
        question = input("问：").strip()
    except (EOFError, KeyboardInterrupt):
        print("\n再见！")
        break
    if not question:
        continue
    if question.lower() in {"q", "quit", "exit", "退出"}:
        print("再见！")
        break
    response = query_engine.query(question)
    print(f"答：{response}\n")
