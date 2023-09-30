from pathlib import Path
import os
import environ
import requests
import re
from datetime import datetime
import telebot
from telebot import custom_filters
import json
import html


# Base settings
BASE_DIR = Path(__file__).resolve(strict=True).parent
ENV_FILE = BASE_DIR / ".env"

# Environments
env = environ.Env()
if Path(ENV_FILE).exists():
    env.read_env(str(ENV_FILE))
    
# Const environment variable
BOT_TOKEN = env("BOT_TOKEN")
SEARCH_ENDPOINT = env("SEARCH_ENDPOINT")

PHONE_REGEX = r"^84[0-9]{9}$"
FACEBOOK_REGEX = r"^[0-9]{15}$"
EMAIL_REGEX = r"^[\w\.-]+@[a-zA-Z\d\.-]+\.[a-zA-Z]{2,}$"

# Start bot
bot = telebot.TeleBot(BOT_TOKEN)

# Bot functions
@bot.message_handler(commands=['start'], chat_id=[1002935752])
def bot_send_welcome(message):
    msg = "Hello, I'm Kyo! Give me a fulcrum, I will 'scale' the world!\n...\nCommands:\n/phone\n/facebook"
    print(f"{datetime.now().strftime('%d/%m/%Y %H:%M:%S')} - [INFO]: Hello")
    bot.reply_to(message, msg)


@bot.message_handler(chat_id=[1002935752])
async def handler(message):
    content = message.text
    if re.match(PHONE_REGEX, message.text):
        payload = {'input': message.text}
        res = requests.request("POST", SEARCH_ENDPOINT + "/phone", data=payload)
        if res.status_code == 200:
            content = res.text
    elif re.match(FACEBOOK_REGEX, message.text):
        payload = {'input': message.text}
        res = requests.request("POST", SEARCH_ENDPOINT + "/facebook", data=payload)
        if res.status_code == 200:
            content = res.text
    elif re.match(EMAIL_REGEX, message.text):
        payload = {'input': message.text}
        res = requests.request("POST", SEARCH_ENDPOINT + "/email", data=payload)
        if res.status_code == 200:
            content = res.text

    # print("{} - [INFO - FBTOPHONE]: {}".format(datetime.now().strftime('%d/%m/%Y %H:%M:%S'), msg.replace('\n', ' ')))
    await bot.reply_to(message, content, parse_mode="Markdown")


def json_to_markdown(json_data):
    # Phân tích JSON thành một đối tượng Python
    data = json.loads(json_data)

    # Hàm đệ quy để tạo markdown từ dữ liệu JSON
    def recursive_markdown(data, indent=0):
        markdown = ''
        for key, value in data.items():
            if isinstance(value, dict):
                markdown += f'{"  " * indent}- **{key}**:\n{recursive_markdown(value, indent + 1)}'
            else:
                # Mã hóa các ký tự đặc biệt bằng thư viện html
                encoded_value = html.escape(str(value))
                markdown += f'{"  " * indent}- **{key}**: {encoded_value}\n'
        return markdown

    # Gọi hàm đệ quy và trả về kết quả
    return recursive_markdown(data)



bot.add_custom_filter(custom_filters.ChatFilter())
bot.infinity_polling()