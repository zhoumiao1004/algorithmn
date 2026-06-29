# coding: utf-8

from anthropic import Anthropic
import re
import math

client = Anthropic()

TOOL_SCHEMA = {
    "name": "calculator",
    "description": "执行数学表达式计算",
    "input_schema": {
        "type": "object",
        "properties": {
            "expression": {
                "type": "string",
                "description": "数学表达式，如 2 + 2 或 sqrt(16)"
            }
        },
        "required": ["expression"]
    }
}

def calc(expression: str) -> str:
    """安全计算器：只支持基础数学运算"""
    # 移除危险字符，只允许数字和运算符
    safe_expr = re.sub(r'[^0-9+\-*/.()sqrtlogcosinasin ]', '', expression)
    try:
        result = eval(safe_expr, {"__builtins__": {}, "sqrt": math.sqrt, "log": math.log, "cos": math.cos, "sin": math.sin}, {})
        return str(result)
    except:
        return "计算错误"

messages = []
while True:
    user_input = input("你: ")
    messages.append({"role": "user", "content": user_input})

    response = client.messages.create(
        model="claude-sonnet-4-20250514",
        max_tokens=1024,
        tools=[TOOL_SCHEMA],
        messages=messages
    )

    # 处理响应
    if response.stop_reason == "tool_use":
        for block in response.content:
            if block.type == "tool_use":
                result = calc(block.input["expression"])
                messages.append({"role": "user", "content": f"计算结果: {result}"})
                messages.append({"role": "assistant", "content": f"好的，结果是 {result}。"})
                print(f"Agent: 结果是 {result}")
    else:
        print(f"Agent: {response.content[0].text}")
