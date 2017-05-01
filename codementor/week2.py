# -*- coding: utf-8 -*-
"""
Created on Tue Oct 18 19:32:07 2016

My learning from week 2 class of CodeMentor Machine Learning

@author: deepan
"""

import numpy
import csv
import matplotlib.pyplot as plt


trainFile = open("/Users/Deepan/workspace/src/codementor/train.csv")
trainData = csv.DictReader(trainFile)

noAge = 0
numPassengers = 0
truths = []
predictions = []
yes = []
no = []

#we get one line of data to predict.
def predict(data):
    age = data['Age']
    if age == '':
        return False #no age. so can't tell the survival 
    else:
        age = float(age)
        if age <= 10:
            return True #children likely to be saved by the adults.
        else:
            if data['Sex'] == 'female':
                return True #female passengers are likely to be saved by males
            else:
                #classify the rest based on their passenger class.
                pclass = data['Pclass']
                if pclass == '1':
                    return True #higher class, better chance of survival
                else:
                    return False


for data in trainData:
    numPassengers += 1    
    age = data['Age']
    
    #Gather the truth value first
    if data['Survived'] == '1':
        truths.append(True)
    else:
        truths.append(False)
    
    #lets collect some data to draw the histogram based on passenger class
    if age == '':
        noAge += 1 #skip them from the yes/no calculations for now
    else:
        pclass = int(data['Pclass'])
        if data['Survived'] == '1':
            yes.append(pclass)
        else:
            no.append(pclass)
            
    #exclude passengers without age
    predictions.append(predict(data))    
    

correct = 0.0
for i in range(len(truths)):
    if truths[i] == predictions[i]:
        correct += 1


print "Total records = ", len(truths)
print "Records with noage = ", noAge
print "Survival Correctness = ", (correct/len(predictions))*100
print "Passengers (with known age) survived: {0}, dead: {1}".format(len(yes), len(no))
print "Yes: ", numpy.mean(yes)
print "No: ", numpy.mean(no)

yesHist = numpy.histogram(yes)
noHist = numpy.histogram(no)

print type(yesHist), type(yesHist[1])

for x in range(len(noHist)):
    print yesHist[x]-noHist[x]


plt.hist(yes, bins=3)
plt.show()

plt.hist(no, bins=3)
plt.show()

#learning things from week 2 class
#packages used
#   csv- to convert raw csv file into readable csv compatible dict form.
#   numpy - to do some math functions such mean, median, sd, histogram
#   matplotlib.pyplot - to plot the histogram buckets
#
#calculated truths and manual predictions based on age, sex, pclass
#correctness came in the range of 74%
#calculated the histogram of survivors based on the passenger class
#and plotted them using 3 bins.