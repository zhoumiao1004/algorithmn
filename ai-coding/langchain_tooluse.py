import os

from langchain_openai import ChatOpenAI
from langchain_classic.agents import create_tool_calling_agent, AgentExecutor
from langchain_core.tools import tool
from langchain_core.prompts import ChatPromptTemplate


# ---------------------------------------------------------------------------
# 配置
API_KEY = os.environ.get("API_KEY")
BASE_URL = "https://api.deepseek.com/v1"
LLM_MODEL = "deepseek-v4-flash"


# 1. 定义工具
@tool
def calculate(expression: str) -> str:
    """计算数学表达式。输入：数学表达式字符串"""
    try:
        return str(eval(expression))
    except Exception as e:
        return f"计算错误：{e}"

@tool
def get_city_weather(city: str) -> str:
    """查询城市天气（模拟）"""
    weather_data = {
        "北京": "晴，28°C，湿度40%",
        "深圳": "多云，32°C，湿度75%",
        "上海": "小雨，26°C，湿度85%",
    }
    return weather_data.get(city, f"暂无{city}的天气数据")

# 2. 创建 Agent
llm = ChatOpenAI(
    model=LLM_MODEL,
    temperature=0,
    api_key=API_KEY,
    base_url=BASE_URL,
)
prompt = ChatPromptTemplate.from_messages([
    ("system", "你是一个助手，可以调用工具来回答问题。"),
    ("human", "{input}"),
    ("placeholder", "{agent_scratchpad}"),
])

agent = create_tool_calling_agent(llm, [calculate, get_city_weather], prompt)
agent_executor = AgentExecutor(
    agent=agent,
    tools=[calculate, get_city_weather],
    verbose=True,
    max_iterations=5,  # ← 防止死循环
)

# 3. 运行
result = agent_executor.invoke({"input": "北京天气怎么样？再算一下 (15+27)*3 等于多少？"})
print(result["output"])