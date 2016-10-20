# -*- coding: utf-8 -*-
"""
Spyder Editor

This is a temporary script file.
"""

import random

#number_to_guess = random.randint(1,50)
#counter = 0
#
#user_input = -1
#
#while (user_input != number_to_guess):
#    user_input = int(raw_input("Enter your guess:"))
#    counter += 1
#    if user_input > number_to_guess:
#        print "Too high guess.."
#    elif user_input < number_to_guess:
#        print "Too less gues.."
#    else:
#        print "Got it right!"
#
#print "number of guess = ", counter
import csv
import numpy
import matplotlib.pyplot as plt

myfile = open("/Users/Deepan/workspace/src/codementor/train.csv")

lines = csv.DictReader(myfile)


predictions = []
truths = []

def predict(datapoint):
    print datapoint['Name']
    age = datapoint['Age']
    if age == '':
        return False
    else:
        age = float(age)
        if age >= 18:
            return True
        else:
            return False

yes = []
no = []
total_records = 0

for line in lines:
    #print line['Name']
    #survivors
    total_records += 1
    age = line['Age']
    if age == '':
        pass
    else:
        age = float(age)
        if line['Survived'] == '1':
            truths.append(True)
            yes.append(age)
        else:
            truths.append(False)
            no.append(age)
    
    predictions.append(predict(line))

    print "--------"
#
#correct = 0.0
#for x in range(len(predictions)):
#    if predictions[x] == truths[x]:
#        correct += 1
#        
#print "Correctness = ", correct/len(predictions)
print "Mean [yes] = ", numpy.mean(yes), len(yes)
print "Mean [no] =", numpy.mean(no), len(no)
print "Total records = ", total_records
print "YES.."
plt.hist(yes, bins=20)
print "NO"
plt.hist(no, bins=20)

