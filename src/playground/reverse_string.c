#include <stdio.h>
#include <string.h>

int reverse(char *word);

int main(void) {
    char input[128];
    printf("Enter your string to reverse:");
    scanf("%s", input);
    reverse(input);
    return 1;
}

/*
 * Given a word, reverse its characters.
 * abcd -> dcba
 */
int reverse(char *word)
{
    //find the end of the string.
    char *start = word;
    char *end = start;
    while (*end != '\0') {
        end++;
    }
    printf ("start = %p, end = %p", start, end);
    return 1;
}
