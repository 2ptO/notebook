# -*- coding: utf-8 -*-
"""
Created on Wed Oct 19 07:00:23 2016

@author: deepan
"""

import csv
import numpy

trainFile = open('/Users/Deepan/workspace/src/codementor/train.csv')
trainData = csv.DictReader(trainFile)

vectors = []
truths = []
predictions = []

def makeVector(data):
    #Given a line of data from the csv dict reader, build a vector
    #of age, sex, and passenger class
    age = data['Age']
    
    #there could be unknown age. convert them to an average value.
    age = 29.5 if age == '' else float(age)
    
    #we need everything in numbers to fit in the regression model.
    #so convert sex into number form.
    sex = 1 if data['Sex'] == 'male' else 0
    
    pclass = int(data['Pclass'])
    
    return [age,sex,pclass]
    
    

#build the truth vector and predictions vector. 
for data in trainData:
    survived = True if data['Survived'] == '1' else False
    truths.append(survived)
    vectors.append(makeVector(data))

#now that we got the truths and vectorized data, lets fit that into 
#a regression model and predict from that. 

from sklearn.linear_model import LogisticRegression
from sklearn.neighbors import KNeighborsClassifier

#model = LogisticRegression()
model = KNeighborsClassifier(n_neighbors=20, weights='distance')
model.fit(vectors, truths)

#have to reload the dict reader. Otherwise, iterator returns nothing.,
trainFile = open('/Users/Deepan/workspace/src/codementor/train.csv')
trainData = csv.DictReader(trainFile)
for data in trainData:
    vec = makeVector(data)
    #note: vec itself is a list.. but we are turning that into a list of list (2d array)
    #otherwise, we will get DeprectionWarning from sklearn validation.py
    predictions.append(model.predict([vec]))

correct = 0.0
for i in range(len(truths)):
    if truths[i] == predictions[i]:
        correct += 1

successRate = correct/len(truths)
print successRate, "correct answers..."
    