from fastapi import FastAPI
from pydantic import BaseModel
from transformers import AutoModelForSequenceClassification, AutoTokenizer
import torch
import torch.nn.functional as F

app = FastAPI()

# Загрузка модели при старте
tokenizer = AutoTokenizer.from_pretrained("cointegrated/rubert-tiny-toxicity")
model = AutoModelForSequenceClassification.from_pretrained("cointegrated/rubert-tiny-toxicity")

class CommentRequest(BaseModel):
    text: str

@app.post("/moderate")
async def moderate_comment(comment: CommentRequest):
    # Ваша логика проверки
    inputs = tokenizer(comment.text, return_tensors="pt", truncation=True, padding=True)
    with torch.no_grad():
        logits = model(**inputs).logits
        probas = torch.sigmoid(logits).squeeze().tolist()
    
    labels = [model.config.id2label[i] for i in range(len(probas))]
    score_by_label = dict(zip(labels, probas))
    toxicity_score = 1.0 - score_by_label.get("non-toxic", 0.0)
    
    is_toxic = toxicity_score > 0.75
    
    return {
        "is_approved": not is_toxic,
        "toxicity_score": toxicity_score,
        "details": score_by_label
    }

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)