import json
from pprint import pprint

from time import sleep

file=open("question.json")
data=json.load(file)
double_answer=[]

sleep_time=1
for i in data:
    if i['answer'] in i['choices']:
        try:
           key=i['choices'].keys()
           if "A" in key:
                question=i["question"]
                choice_a=i['choices']["A"]
                choice_b=i['choices']["B"]
                choice_c=i['choices']["C"]
                choice_d=i['choices']["D"]
                answer=i['choices'][i['answer']]
                explanation=i['explanation']
                
                sleep_time+=1
           else:
                question=i["question"]
                choice_a=i['choices']["a"]
                choice_b=i['choices']["b"]
                choice_c=i['choices']["c"]
                choice_d=i['choices']["d"]
                answer=i['choices'][i['answer']]
                explanation=i['explanation']
                
                sleep_time+=1
           if sleep_time%5==0:
               sleep(5)
               print("At index",sleep_time)
        except Exception as e:
            pprint(e)
    else:
        double_answer.append(i)

