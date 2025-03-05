import json

file=open("question.json")

data=json.load(file)
count=0
for i in data:
    print(i['question'])
    count+=1

print('number of question',count)