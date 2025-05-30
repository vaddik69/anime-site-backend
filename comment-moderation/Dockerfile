FROM python:3.9-slim

WORKDIR /app

# Устанавливаем зависимости для torch
RUN apt-get update && apt-get install -y libgomp1

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 8000
CMD ["uvicorn", "moderation:app", "--host", "0.0.0.0", "--port", "8000"]